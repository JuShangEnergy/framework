package controllers

import (
	"github.com/JuShangEnergy/framework/cloud"
	"github.com/JuShangEnergy/framework/errs"
	"github.com/JuShangEnergy/framework/types"
)

// FunctionsController 处理 /functions 接口的请求
type FunctionsController struct {
	ClassesController
}

// HandleCloudFunction 执行指定的云函数
// 返回数据格式如下：
//
//	{
//		"result":"func res"
//	}
//
// @router /:functionName [post]
func (f *FunctionsController) HandleCloudFunction() {
	functionName := f.Ctx.Input.Param(":functionName")
	theFunction := cloud.GetFunction(functionName)
	theValidator := cloud.GetValidator(functionName)
	if theFunction == nil {
		f.HandleError(errs.E(errs.ScriptFailed, "Invalid function: "+functionName), 0)
		return
	}

	if f.JSONBody == nil {
		f.JSONBody = types.M{}
	}

	params := types.M{}
	for k, v := range f.JSONBody {
		params[k] = v
	}
	for k, v := range f.Query {
		params[k] = v
	}

	headers := map[string]string{}
	for k := range f.Ctx.Request.Header {
		headers[k] = f.Ctx.Request.Header.Get(k)
	}

	request := cloud.FunctionRequest{
		Params:         params,
		Master:         false,
		InstallationID: f.Info.InstallationID,
		FunctionName:   functionName,
		Headers:        headers,
	}
	if f.Auth != nil {
		request.Master = f.Auth.IsMaster
		request.User = f.Auth.User
	}

	if theValidator != nil {
		result := theValidator(request)
		if result == false {
			f.HandleError(errs.E(errs.ValidationError, "Validation failed."), 0)
			return
		}
	}

	response := &cloud.FunctionResponse{}
	theFunction(request, response)
	if response.Err != nil {
		f.HandleError(response.Err, 0)
		return
	}

	if response.Response != nil {
		v, ok := response.Response["result"].(types.M)
		if ok && v["Content-Type"] != nil {
			f.Ctx.Output.Header("Content-Type", v["Content-Type"].(string))
			f.Ctx.Output.Header("Content-Length", v["Content-Length"].(string))
			f.Ctx.Output.Header("Content-Disposition", "attachment; filename="+v["filename"].(string))
			f.Ctx.Output.Body(v["data"].([]byte))
			return
		}
	}

	f.Data["json"] = response.Response
	f.ServeJSON()
}

// @router /:functionName [get]
func (f *FunctionsController) HandleCloudFunctionGet() {
	functionName := f.Ctx.Input.Param(":functionName")
	theFunction := cloud.GetFunction(functionName)
	theValidator := cloud.GetValidator(functionName)
	if theFunction == nil {
		f.HandleError(errs.E(errs.ScriptFailed, "Invalid function: "+functionName), 0)
		return
	}

	if f.JSONBody == nil {
		f.JSONBody = types.M{}
	}

	params := types.M{}
	for k, v := range f.JSONBody {
		params[k] = v
	}
	for k, v := range f.Query {
		params[k] = v
	}

	headers := map[string]string{}
	for k := range f.Ctx.Request.Header {
		headers[k] = f.Ctx.Request.Header.Get(k)
	}

	request := cloud.FunctionRequest{
		Params:         params,
		Master:         false,
		InstallationID: f.Info.InstallationID,
		FunctionName:   functionName,
		Headers:        headers,
	}
	if f.Auth != nil {
		request.Master = f.Auth.IsMaster
		request.User = f.Auth.User
	}

	if theValidator != nil {
		result := theValidator(request)
		if result == false {
			f.HandleError(errs.E(errs.ValidationError, "Validation failed."), 0)
			return
		}
	}

	response := &cloud.FunctionResponse{}
	theFunction(request, response)
	if response.Err != nil {
		f.HandleError(response.Err, 0)
		return
	}

	if response.Response != nil {
		v, ok := response.Response["result"].(types.M)
		if ok && v["Content-Type"] != nil {
			f.Ctx.Output.Header("Content-Type", v["Content-Type"].(string))
			f.Ctx.Output.Header("Content-Length", v["Content-Length"].(string))
			f.Ctx.Output.Header("Content-Disposition", "attachment; filename="+v["filename"].(string))
			f.Ctx.Output.Body(v["data"].([]byte))
			return
		}
	}
	f.Data["json"] = response.Response
	f.ServeJSON()

}

// Get ...
// @router / [get]
func (f *FunctionsController) Get() {
	f.ClassesController.Get()
}

// Post ...
// @router / [post]
func (f *FunctionsController) Post() {
	f.ClassesController.Post()
}

// Delete ...
// @router / [delete]
func (f *FunctionsController) Delete() {
	f.ClassesController.Delete()
}

// Put ...
// @router / [put]
func (f *FunctionsController) Put() {
	f.ClassesController.Put()
}
