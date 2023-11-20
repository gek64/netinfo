package routers

import (
	"github.com/gin-gonic/gin"
	"netinfo/internal/receive/controllers/debugService"
	"netinfo/internal/receive/controllers/recordService"
)

func LoadRecordRouters(router *gin.Engine) {
	recordRouterGroup := router.Group("/")
	{
		// 增加记录
		recordRouterGroup.POST("/", recordService.UpsertRecord)
		// 修改记录
		recordRouterGroup.PUT("/", recordService.UpsertRecord)
		// 查询记录
		recordRouterGroup.GET("/", recordService.ReadRecordByID)
		recordRouterGroup.GET("/all", recordService.ReadRecordAll)
		// 删除记录
		recordRouterGroup.DELETE("/", recordService.DeleteRecordByID)
		recordRouterGroup.DELETE("/all", recordService.DeleteRecordAll)
	}
}

func LoadDebugRouters(router *gin.Engine) {
	debugRouterGroup := router.Group("/debug/")
	{
		debugRouterGroup.GET("/", debugService.Reflector)
		debugRouterGroup.HEAD("/", debugService.Reflector)
		debugRouterGroup.POST("/", debugService.Reflector)
		debugRouterGroup.PUT("/", debugService.Reflector)
		debugRouterGroup.DELETE("/", debugService.Reflector)
		debugRouterGroup.OPTIONS("/", debugService.Reflector)
		debugRouterGroup.PATCH("/", debugService.Reflector)
	}
}
