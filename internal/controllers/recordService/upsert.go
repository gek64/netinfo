package recordService

import (
	"context"
	"github.com/gek64/gek/gNet"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"netinfo/ent"
	"netinfo/ent/record"
	"netinfo/internal/middleware"
)

// UpsertRecord 内容记录
func UpsertRecord(c *gin.Context) {
	var reqBody RecordBody
	client := c.MustGet(middleware.Client).(*ent.Client)
	ctx := c.MustGet(middleware.Context).(context.Context)

	// 请求的body数据绑定到结构体
	err := c.ShouldBindWith(&reqBody, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// 获取请求IP
	requestIPString, err := gNet.GetIPFromRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	// 数据库操作
	updateCount, err := client.Record.Update().
		Where(record.IDEQ(reqBody.ID)).
		SetDescription(reqBody.Description).
		SetNetInterfaces(reqBody.NetInterfaces).
		SetRequestIP(requestIPString).
		Save(ctx)
	if err != nil || updateCount == 0 {
		_, err := client.Record.Create().
			SetID(reqBody.ID).
			SetDescription(reqBody.Description).
			SetNetInterfaces(reqBody.NetInterfaces).
			SetRequestIP(requestIPString).
			Save(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, gin.H{})
		}
		return
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}
