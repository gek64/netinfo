package recordService

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"netinfo/ent"
	"netinfo/ent/record"
	"netinfo/ent/schema"
	"netinfo/internal/middleware"
)

type UpdateRecordBody struct {
	Id            uint                  `json:"id" binding:"required"`
	Description   string                `json:"description"`
	NetInterfaces []schema.NetInterface `json:"netInterfaces"`
}

// UpdateRecordByID 修改记录
func UpdateRecordByID(c *gin.Context) {
	var reqBody UpdateRecordBody
	client := c.MustGet(middleware.Client).(*ent.Client)
	ctx := c.MustGet(middleware.Context).(context.Context)

	// put请求的json数据绑定到结构体
	err := c.ShouldBindWith(&reqBody, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// 更新到数据库
	operationCount, err := client.Record.Update().
		Where(record.IDEQ(reqBody.Id)).
		SetDescription(reqBody.Description).
		SetNetInterfaces(reqBody.NetInterfaces).
		Save(ctx)
	if err != nil || operationCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
