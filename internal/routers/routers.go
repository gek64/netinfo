package routers

import (
	"github.com/gin-gonic/gin"
	"netinfo/internal/controllers/queryService"
	"netinfo/internal/controllers/recordService"
)

func LoadRouters(router *gin.Engine) {
	recordRouterGroup := router.Group("/record")
	{
		// 增加记录
		recordRouterGroup.POST("/", recordService.CreateRecord)
		// 删除记录
		recordRouterGroup.DELETE("/", recordService.DeleteRecordByID)
		// 修改记录
		recordRouterGroup.PUT("/", recordService.UpdateRecordByID)
		// 查询记录
		recordRouterGroup.GET("/", recordService.ReadRecordByID)
		recordRouterGroup.GET("/all", recordService.ReadRecordAll)
	}

	queryRouterGroup := router.Group("/query")
	{
		queryRouterGroup.GET("/", queryService.Query)
	}
}
