package routers

import (
	"github.com/beego/beego"
	"github.com/beego/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AggregateController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AggregateController"],
		beego.ControllerComments{
			Method:           "HandleFind",
			Router:           `/:className`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AnalyticsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AnalyticsController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AnalyticsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AnalyticsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AnalyticsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AnalyticsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AnalyticsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AnalyticsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AnalyticsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AnalyticsController"],
		beego.ControllerComments{
			Method:           "HandleEvent",
			Router:           `/:eventName`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AnalyticsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AnalyticsController"],
		beego.ControllerComments{
			Method:           "AppOpened",
			Router:           `/AppOpened`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AudiencesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AudiencesController"],
		beego.ControllerComments{
			Method:           "HandleFind",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AudiencesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AudiencesController"],
		beego.ControllerComments{
			Method:           "HandleCreate",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AudiencesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AudiencesController"],
		beego.ControllerComments{
			Method:           "HandleGet",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AudiencesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AudiencesController"],
		beego.ControllerComments{
			Method:           "HandleUpdate",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AudiencesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:AudiencesController"],
		beego.ControllerComments{
			Method:           "HandleDelete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:BatchController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:BatchController"],
		beego.ControllerComments{
			Method:           "HandleBatch",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:BatchController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:BatchController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:BatchController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:BatchController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:BatchController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:BatchController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"],
		beego.ControllerComments{
			Method:           "HandleFind",
			Router:           `/:className`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"],
		beego.ControllerComments{
			Method:           "HandleCreate",
			Router:           `/:className`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"],
		beego.ControllerComments{
			Method:           "HandleDelete",
			Router:           `/:className/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"],
		beego.ControllerComments{
			Method:           "HandleGet",
			Router:           `/:className/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ClassesController"],
		beego.ControllerComments{
			Method:           "HandleUpdate",
			Router:           `/:className/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:CloudCodeController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:CloudCodeController"],
		beego.ControllerComments{
			Method:           "DeleteJob",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:CloudCodeController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:CloudCodeController"],
		beego.ControllerComments{
			Method:           "GetJobs",
			Router:           `/jobs`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:CloudCodeController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:CloudCodeController"],
		beego.ControllerComments{
			Method:           "CreateJob",
			Router:           `/jobs`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:CloudCodeController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:CloudCodeController"],
		beego.ControllerComments{
			Method:           "EditJob",
			Router:           `/jobs/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:CloudCodeController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:CloudCodeController"],
		beego.ControllerComments{
			Method:           "GetJobsData",
			Router:           `/jobs/data`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FeaturesController"],
		beego.ControllerComments{
			Method:           "HandleGet",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FeaturesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FeaturesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FeaturesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"],
		beego.ControllerComments{
			Method:           "HandleGet",
			Router:           `/:appId/:filename`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"],
		beego.ControllerComments{
			Method:           "HandleCreate",
			Router:           `/:filename`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FilesController"],
		beego.ControllerComments{
			Method:           "HandleDelete",
			Router:           `/:filename`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FunctionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FunctionsController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FunctionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FunctionsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FunctionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FunctionsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FunctionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FunctionsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FunctionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FunctionsController"],
		beego.ControllerComments{
			Method:           "HandleCloudFunction",
			Router:           `/:functionName`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FunctionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:FunctionsController"],
		beego.ControllerComments{
			Method:           "HandleCloudFunctionGet",
			Router:           `/:functionName`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:GlobalConfigController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:GlobalConfigController"],
		beego.ControllerComments{
			Method:           "HandleGet",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:GlobalConfigController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:GlobalConfigController"],
		beego.ControllerComments{
			Method:           "HandlePut",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:GlobalConfigController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:GlobalConfigController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:GlobalConfigController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:GlobalConfigController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HealthController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HealthController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"],
		beego.ControllerComments{
			Method:           "HandleGetAllFunctions",
			Router:           `/functions`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"],
		beego.ControllerComments{
			Method:           "HandleCreateFunction",
			Router:           `/functions`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"],
		beego.ControllerComments{
			Method:           "HandleGetFunction",
			Router:           `/functions/:functionName`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"],
		beego.ControllerComments{
			Method:           "HandleUpdateFunction",
			Router:           `/functions/:functionName`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"],
		beego.ControllerComments{
			Method:           "HandleCreateTrigger",
			Router:           `/triggers`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"],
		beego.ControllerComments{
			Method:           "HandleGetAllTriggers",
			Router:           `/triggers`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"],
		beego.ControllerComments{
			Method:           "HandleUpdateTrigger",
			Router:           `/triggers/:className/:triggerName`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:HooksController"],
		beego.ControllerComments{
			Method:           "HandleGetTrigger",
			Router:           `/triggers/:className/:triggerName`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:IAPValidationController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:IAPValidationController"],
		beego.ControllerComments{
			Method:           "HandlePost",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:IAPValidationController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:IAPValidationController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:IAPValidationController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:IAPValidationController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:IAPValidationController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:IAPValidationController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"],
		beego.ControllerComments{
			Method:           "HandleFind",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"],
		beego.ControllerComments{
			Method:           "HandleCreate",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"],
		beego.ControllerComments{
			Method:           "HandleGet",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"],
		beego.ControllerComments{
			Method:           "HandleUpdate",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:InstallationsController"],
		beego.ControllerComments{
			Method:           "HandleDelete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:JobsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:JobsController"],
		beego.ControllerComments{
			Method:           "HandlePost",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:JobsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:JobsController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:JobsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:JobsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:JobsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:JobsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:JobsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:JobsController"],
		beego.ControllerComments{
			Method:           "HandleCloudJob",
			Router:           `/:jobName`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "HandleLogIn",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogoutController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogoutController"],
		beego.ControllerComments{
			Method:           "HandleLogOut",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogoutController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogoutController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogoutController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogoutController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogoutController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogoutController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogsController"],
		beego.ControllerComments{
			Method:           "HandleGet",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogsController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:LogsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "ChangePassword",
			Router:           `/choose_password`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "InvalidLink",
			Router:           `/invalid_link`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "InvalidVerificationLink",
			Router:           `/invalid_verification_link`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "LinkSendFail",
			Router:           `/link_send_fail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "LinkSendSuccess",
			Router:           `/link_send_success`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "PasswordResetSuccess",
			Router:           `/password_reset_success`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "RequestResetPassword",
			Router:           `/request_password_reset`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "ResetPassword",
			Router:           `/request_password_reset`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "ResendVerificationEmail",
			Router:           `/resend_verification_email`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "VerifyEmail",
			Router:           `/verify_email`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PublicController"],
		beego.ControllerComments{
			Method:           "VerifyEmailSuccess",
			Router:           `/verify_email_success`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PurgeController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PurgeController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PurgeController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PurgeController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PurgeController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PurgeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PurgeController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PurgeController"],
		beego.ControllerComments{
			Method:           "HandleDelete",
			Router:           `/:className`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PushController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PushController"],
		beego.ControllerComments{
			Method:           "HandlePost",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PushController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PushController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PushController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PushController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PushController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:PushController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ResetController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ResetController"],
		beego.ControllerComments{
			Method:           "HandleResetRequest",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ResetController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ResetController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ResetController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ResetController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ResetController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:ResetController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"],
		beego.ControllerComments{
			Method:           "HandleFind",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"],
		beego.ControllerComments{
			Method:           "HandleCreate",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"],
		beego.ControllerComments{
			Method:           "HandleGet",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"],
		beego.ControllerComments{
			Method:           "HandleUpdate",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:RolesController"],
		beego.ControllerComments{
			Method:           "HandleDelete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"],
		beego.ControllerComments{
			Method:           "HandleFind",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"],
		beego.ControllerComments{
			Method:           "HandleGet",
			Router:           `/:className`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"],
		beego.ControllerComments{
			Method:           "HandleCreate",
			Router:           `/:className`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"],
		beego.ControllerComments{
			Method:           "HandleUpdate",
			Router:           `/:className`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SchemasController"],
		beego.ControllerComments{
			Method:           "HandleDelete",
			Router:           `/:className`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"],
		beego.ControllerComments{
			Method:           "HandleFind",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"],
		beego.ControllerComments{
			Method:           "HandleCreate",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"],
		beego.ControllerComments{
			Method:           "HandleUpdate",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"],
		beego.ControllerComments{
			Method:           "HandleDelete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"],
		beego.ControllerComments{
			Method:           "HandleGet",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"],
		beego.ControllerComments{
			Method:           "HandleGetMe",
			Router:           `/me`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:SessionsController"],
		beego.ControllerComments{
			Method:           "HandleUpdateMe",
			Router:           `/me`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UpgradeSessionController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UpgradeSessionController"],
		beego.ControllerComments{
			Method:           "HandleUpdateToRevocableSession",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UpgradeSessionController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UpgradeSessionController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UpgradeSessionController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UpgradeSessionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UpgradeSessionController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UpgradeSessionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "HandleFind",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "HandleCreate",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "HandleUpdate",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "HandleDelete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "HandleGet",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "HandleMe",
			Router:           `/me`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:VerificationController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:VerificationController"],
		beego.ControllerComments{
			Method:           "HandleVerificationEmailRequest",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:VerificationController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:VerificationController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:VerificationController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:VerificationController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:VerificationController"] = append(beego.GlobalControllerRouter["github.com/JuShangEnergy/framework/controllers:VerificationController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
