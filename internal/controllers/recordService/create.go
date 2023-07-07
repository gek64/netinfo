package recordService

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"netinfo/ent"
	"netinfo/ent/schema"
	"netinfo/internal/middleware"
)

type CreateRecordBody struct {
	ID            string                `json:"id" form:"id" binding:"required"`
	Description   string                `json:"description" form:"description"`
	NetInterfaces []schema.NetInterface `json:"netInterfaces" form:"netInterfaces" binding:"required"`
}

// CreateRecord 创建内容记录
func CreateRecord(c *gin.Context) {
	var reqBody CreateRecordBody
	client := c.MustGet(middleware.Client).(*ent.Client)
	ctx := c.MustGet(middleware.Context).(context.Context)

	// post请求的json数据绑定到结构体
	err := c.ShouldBindBodyWith(&reqBody, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// 新增到数据库
	recordCreated, err := client.Record.Create().
		SetID(reqBody.ID).
		SetDescription(reqBody.Description).
		SetNetInterfaces(reqBody.NetInterfaces).
		Save(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recordCreated)
}
