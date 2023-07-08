package recordService

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"netinfo/ent"
	"netinfo/ent/record"
	"netinfo/internal/middleware"
)

// ReadRecordByID 搜索记录(ID)
func ReadRecordByID(c *gin.Context) {
	var reqQuery RecordQuery
	client := c.MustGet(middleware.Client).(*ent.Client)
	ctx := c.MustGet(middleware.Context).(context.Context)

	// get请求的query表单数据绑定到结构体
	err := c.ShouldBindQuery(&reqQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// 查询主键id为reqQuery.Id的第一个记录到result
	result, err := client.Record.Query().
		Where(record.IDEQ(reqQuery.ID)).
		First(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	// 向resp中返回查询到的记录
	c.JSON(http.StatusOK, result)
}

// ReadRecordAll 搜索所有记录
func ReadRecordAll(c *gin.Context) {
	client := c.MustGet(middleware.Client).(*ent.Client)
	ctx := c.MustGet(middleware.Context).(context.Context)

	// 查询主键id为reqQuery.Id的第一个记录到result
	records, err := client.Record.Query().
		All(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	// 向resp中返回查询到的记录
	c.JSON(http.StatusOK, records)
}
