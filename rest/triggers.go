package rest

import (
	"github.com/JuShangEnergy/framework/cloud"
	"github.com/JuShangEnergy/framework/types"
	"github.com/JuShangEnergy/framework/utils"
)

func getRequest(triggerType string, auth *Auth, parseObject, originalParseObject types.M) cloud.TriggerRequest {
	request := cloud.TriggerRequest{
		TriggerName: triggerType,
		Object:      parseObject,
		Master:      false,
		Headers:     auth.Info.Headers,
	}

	if originalParseObject != nil {
		request.Original = originalParseObject
	}

	if auth == nil {
		return request
	}
	request.Master = auth.IsMaster
	if auth.User != nil {
		request.User = auth.User
	}
	if auth.InstallationID != "" {
		request.InstallationID = auth.InstallationID
	}

	return request
}

func getResponse(request cloud.TriggerRequest) *cloud.TriggerResponse {
	response := &cloud.TriggerResponse{
		Request: request,
	}
	return response
}

func getRequestQuery(triggerType string, auth *Auth, query types.M, count bool, isGet bool) cloud.TriggerRequest {
	request := cloud.TriggerRequest{
		TriggerName: triggerType,
		Query:       query,
		Count:       count,
		Master:      false,
		IsGet:       isGet,
		Headers:     auth.Info.Headers,
	}

	if auth == nil {
		return request
	}
	request.Master = auth.IsMaster
	if auth.User != nil {
		request.User = auth.User
	}
	if auth.InstallationID != "" {
		request.InstallationID = auth.InstallationID
	}

	return request
}

func maybeRunTrigger(triggerType string, auth *Auth, parseObject, originalParseObject types.M) (types.M, error) {
	if parseObject == nil {
		return types.M{}, nil
	}

	trigger := cloud.GetTrigger(triggerType, utils.S(parseObject["className"]))
	if trigger == nil {
		return types.M{}, nil
	}
	request := getRequest(triggerType, auth, parseObject, originalParseObject)
	response := getResponse(request)
	trigger(request, response)
	return response.Response, response.Err
}

func maybeRunQueryTrigger(triggerType, className string, restWhere, restOptions types.M, auth *Auth, isGet bool, isAggregate, isDistinct bool) (types.M, types.M, error) {
	trigger := cloud.GetTrigger(triggerType, className)
	if trigger == nil {
		return restWhere, restOptions, nil
	}

	if restOptions == nil {
		restOptions = types.M{}
	}
	query := types.M{}
	if restWhere != nil {
		query["where"] = restWhere
	}
	count := false
	if restOptions["include"] != nil {
		query["include"] = restOptions["include"]
	}
	if restOptions["skip"] != nil {
		query["skip"] = restOptions["skip"]
	}
	if restOptions["limit"] != nil {
		query["limit"] = restOptions["limit"]
	}
	if restOptions["count"] != nil {
		count = true
	}
	if restOptions["order"] != nil {
		query["order"] = restOptions["order"]
	}

	request := getRequestQuery(triggerType, auth, query, count, isGet)
	request.IsAggregate = isAggregate
	request.IsDistinct = isDistinct
	response := getResponse(request)
	trigger(request, response)

	if response.Err != nil {
		return nil, nil, response.Err
	}
	if response.Response == nil {
		return restWhere, restOptions, nil
	}
	if where := utils.M(response.Response["where"]); where != nil {
		restWhere = where
	}
	if limit := response.Response["limit"]; limit != nil {
		restOptions["limit"] = limit
	}
	if skip := response.Response["skip"]; skip != nil {
		restOptions["skip"] = skip
	}
	if include := response.Response["include"]; include != nil {
		restOptions["include"] = include
	}
	if keys := response.Response["keys"]; keys != nil {
		restOptions["keys"] = keys
	}
	if order := response.Response["order"]; order != nil {
		restOptions["order"] = order
	}

	return restWhere, restOptions, nil
}

func maybeRunAfterFindTrigger(triggerType, className string, objects types.S, auth *Auth, restOptions, restWhere types.M) (types.S, error) {
	trigger := cloud.GetTrigger(triggerType, className)
	if trigger == nil {
		return objects, nil
	}

	if restOptions == nil {
		restOptions = types.M{}
	}
	query := types.M{}
	if restWhere != nil {
		query["where"] = restWhere
	}
	count := false
	if restOptions["include"] != nil {
		query["include"] = restOptions["include"]
	}
	if restOptions["skip"] != nil {
		query["skip"] = restOptions["skip"]
	}
	if restOptions["limit"] != nil {
		query["limit"] = restOptions["limit"]
	}
	if restOptions["count"] != nil {
		count = true
	}
	if restOptions["order"] != nil {
		query["order"] = restOptions["order"]
	}

	request := getRequestQuery(triggerType, auth, query, count, false)
	request.IsAggregate = isRequestAggregate(restOptions)
	request.IsDistinct = isRequestDistinct(restOptions)
	//request := getRequest(triggerType, auth, nil, nil)
	response := getResponse(request)
	request.Objects = objects
	trigger(request, response)

	if response.Err != nil {
		return nil, response.Err
	}
	return response.ResponseObjects, nil
}

func inflate(data, restObject types.M) types.M {
	result := types.M{}
	if data != nil {
		for k, v := range data {
			result[k] = v
		}
	}
	if restObject != nil {
		for k, v := range restObject {
			result[k] = v
		}
	}

	return result
}
