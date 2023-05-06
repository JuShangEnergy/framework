package rest

import (
	"github.com/JuShangEnergy/framework/cloud"
	"github.com/JuShangEnergy/framework/errs"
	"github.com/JuShangEnergy/framework/livequery"
	"github.com/JuShangEnergy/framework/types"
	"github.com/JuShangEnergy/framework/utils"
)

// Find 根据条件查找数据
// 返回格式如下：
//
//	{
//		"results":[
//			{...},
//		],
//		"count":10
//	}
func Find(auth *Auth, className string, where, options types.M, clientSDK map[string]string) (types.M, error) {

	err := enforceRoleSecurity("find", className, auth)
	if err != nil {
		return nil, err
	}
	w, o, err := maybeRunQueryTrigger(cloud.TypeBeforeFind, className, where, options, auth, false, isRequestAggregate(options), isRequestDistinct(options))
	if err != nil {
		return nil, err
	}
	if w != nil {
		where = w
	}
	if o != nil {
		options = o
	}
	query, err := NewQuery(auth, className, where, options, clientSDK)
	if err != nil {
		return nil, err
	}

	return query.Execute()
}

// Get ...
func Get(auth *Auth, className, objectID string, options types.M, clientSDK map[string]string) (types.M, error) {
	restWhere := types.M{"objectId": objectID}
	err := enforceRoleSecurity("get", className, auth)
	if err != nil {
		return nil, err
	}
	w, o, err := maybeRunQueryTrigger(cloud.TypeBeforeFind, className, restWhere, options, auth, true, false, false)
	if err != nil {
		return nil, err
	}
	if w != nil {
		restWhere = w
	}
	if o != nil {
		options = o
	}
	query, err := NewQuery(auth, className, restWhere, options, clientSDK)
	if err != nil {
		return nil, err
	}

	return query.Execute()
}

// Delete 删除指定对象
func Delete(auth *Auth, className, objectID string) error {

	if className == "_User" && auth.CouldUpdateUserID(objectID) == false {
		return errs.E(errs.SessionMissing, "insufficient auth to delete user")
	}

	err := enforceRoleSecurity("delete", className, auth)
	if err != nil {
		return err
	}

	var inflatedObject types.M
	// 如果存在删前回调、或者删后回调、或者要删除的属于 _Session 类，则需要获取到要删除的对象数据
	hasTriggers := checkTriggers(className, []string{cloud.TypeBeforeDelete, cloud.TypeAfterDelete})
	hasLiveQuery := checkLiveQuery(className)
	if hasTriggers || hasLiveQuery || className == "_Session" {
		response, err := Find(auth, className, types.M{"objectId": objectID}, types.M{}, nil)
		if err != nil || utils.HasResults(response) == false {
			return errs.E(errs.ObjectNotFound, "Object not found for delete.")
		}

		result := utils.A(response["results"])
		inflatedObject = utils.M(result[0])
		if inflatedObject == nil {
			return errs.E(errs.ObjectNotFound, "Object not found for delete.")
		}
		inflatedObject["className"] = className
		if className == "_Session" && !auth.IsMaster {
			objectId, ok := inflatedObject["user"].(types.M)["objectId"].(string)
			if auth.User == nil || !ok || objectId != auth.User["objectId"].(string) {
				return errs.E(errs.InvalidSessionToken, "invalid session token")
			}
		}
	}

	destroy := NewDestroy(auth, className, types.M{"objectId": objectID}, inflatedObject)

	return destroy.Execute()
}

// Create 创建对象
// 返回数据格式如下：
//
//	{
//		"status":201,
//		"response":{...},
//		"location":"http://..."
//	}
func Create(auth *Auth, className string, object types.M, clientSDK map[string]string) (types.M, error) {

	err := enforceRoleSecurity("create", className, auth)
	if err != nil {
		return nil, err
	}
	write, err := NewWrite(auth, className, nil, object, nil, clientSDK)
	if err != nil {
		return nil, err
	}

	return write.Execute()
}

// Update 更新对象
// 返回更新后的字段，一般只有 updatedAt
func Update(auth *Auth, className, objectID string, object types.M, clientSDK map[string]string) (types.M, error) {

	err := enforceRoleSecurity("update", className, auth)
	if err != nil {
		return nil, err
	}

	var originalRestObject types.M

	// 如果存在删前回调、或者删后回调，则需要获取到要删除的对象数据
	var response types.M
	hasTriggers := checkTriggers(className, []string{cloud.TypeBeforeSave, cloud.TypeAfterSave})
	hasLiveQuery := checkLiveQuery(className)
	if hasTriggers || hasLiveQuery {
		response, err = Find(auth, className, types.M{"objectId": objectID}, types.M{}, clientSDK)
		if err != nil || utils.HasResults(response) == false {
			return nil, errs.E(errs.ObjectNotFound, "Object not found for update.")
		}

		result := utils.A(response["results"])
		originalRestObject = utils.M(result[0])
		if originalRestObject == nil {
			return nil, errs.E(errs.ObjectNotFound, "Object not found for update.")
		}
	}

	write, err := NewWrite(auth, className, types.M{"objectId": objectID}, object, originalRestObject, clientSDK)
	if err != nil {
		return nil, err
	}

	return write.Execute()
}

// enforceRoleSecurity 对指定的类与操作进行安全校验
func enforceRoleSecurity(method string, className string, auth *Auth) error {
	classesWithMasterOnlyAccess := []string{"_JobStatus", "_PushStatus", "_Hooks", "_GlobalConfig", "_JobSchedule"}

	// 非 Master 不得对 _Installation 进行删除与查找操作操作
	if className == "_Installation" && auth.IsMaster == false {
		if method == "delete" || method == "find" {
			msg := "Clients aren't allowed to perform the " + method + " operation on the installation collection."
			return errs.E(errs.OperationForbidden, msg)
		}
	}

	//all volatileClasses are masterKey only
	if utils.StringInSlice(className, classesWithMasterOnlyAccess) && !auth.IsMaster {
		msg := "Clients aren't allowed to perform the " + method + " operation on the " + className + " collection."
		return errs.E(errs.OperationForbidden, msg)
	}

	// readOnly masterKey is not allowed
	if auth.IsReadOnly {
		if method == "delete" || method == "create" || method == "update" {
			msg := "read-only masterKey isn't allowed to perform the " + className + " operation."
			return errs.E(errs.OperationForbidden, msg)
		}
	}
	return nil
}

func checkTriggers(className string, triggerTypes []string) bool {
	result := false
	for _, triggerType := range triggerTypes {
		result = result || cloud.TriggerExists(triggerType, className)
	}
	return result
}

func checkLiveQuery(className string) bool {
	return livequery.TLiveQuery != nil && livequery.TLiveQuery.HasLiveQuery(className)
}

func isRequestAggregate(options types.M) bool {
	if _, has := options["aggregate"]; has {
		return has
	}
	return false
}

func isRequestDistinct(options types.M) bool {
	if _, has := options["distinct"]; has {
		return has
	}
	return false
}
