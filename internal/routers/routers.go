package routers

import (
	"github.com/gin-gonic/gin"
	"netinfo/internal/controllers/queryController"
	"netinfo/internal/controllers/recordController"
)

func LoadRouters(router *gin.Engine) {
	recordRouterGroup := router.Group("/record")
	{
		// 增加记录
		recordRouterGroup.POST("/", recordController.CreateRecord)
		// 删除记录
		recordRouterGroup.DELETE("/", recordController.DeleteRecordByID)
		// 修改记录
		recordRouterGroup.PUT("/", recordController.UpdateRecordByID)
		// 查询记录
		recordRouterGroup.GET("/", recordController.ReadRecordByID)
		recordRouterGroup.GET("/all", recordController.ReadRecordAll)
	}

	queryRouterGroup := router.Group("/query")
	{
		queryRouterGroup.GET("/", queryController.Query)
	}
}
