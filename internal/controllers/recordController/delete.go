package recordController

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"netinfo/ent"
	"netinfo/ent/record"
	"netinfo/internal/middleware"
)

type deleteRecordQuery struct {
	Id uint `form:"id" binding:"required"`
}

// DeleteRecordByID 搜索内容记录
func DeleteRecordByID(c *gin.Context) {
	var reqQuery deleteRecordQuery
	client := c.MustGet(middleware.Client).(*ent.Client)
	ctx := c.MustGet(middleware.Context).(context.Context)

	// delete请求的query表单数据绑定到结构体
	err := c.ShouldBindQuery(&reqQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// 按主键id=reqQuery.Id,删除记录
	operationCount, err := client.Record.Delete().
		Where(record.IDEQ(reqQuery.Id)).
		Exec(ctx)
	if err != nil || operationCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "record not found"})
		return
	}

	// 向resp中返回删除成功
	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
