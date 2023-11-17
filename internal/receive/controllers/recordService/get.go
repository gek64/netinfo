package recordService

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"netinfo/internal/cache"
)

// ReadRecordByID 搜索记录(ID)
func ReadRecordByID(c *gin.Context) {
	var reqQuery RecordQuery
	// get请求的query表单数据绑定到结构体
	err := c.ShouldBindQuery(&reqQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	database, ok := cache.Get(Database)
	if ok {
		for _, d := range database.([]NetInfoInMemoryData) {
			if d.ID == reqQuery.ID {
				c.JSON(http.StatusOK, d)
				return
			}
		}
	}
	c.Status(http.StatusNotFound)
}

// ReadRecordAll 搜索所有记录
func ReadRecordAll(c *gin.Context) {
	database, ok := cache.Get(Database)
	if ok {
		c.JSON(http.StatusOK, database)
	} else {
		c.JSON(http.StatusOK, []NetInfoInMemoryData{})
	}
}
