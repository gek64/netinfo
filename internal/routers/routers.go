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
		recordRouterGroup.POST("/", recordService.UpsertRecord)
		// 修改记录
		recordRouterGroup.PUT("/", recordService.UpsertRecord)
		// 查询记录
		recordRouterGroup.GET("/", recordService.ReadRecordByID)
		recordRouterGroup.GET("/all", recordService.ReadRecordAll)
	}

	queryRouterGroup := router.Group("/")
	{
		queryRouterGroup.GET("/", queryService.Query)
	}
}
