package recordService

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"netinfo/internal/cache"
)

// DeleteRecordByID 删除记录(ID)
func DeleteRecordByID(c *gin.Context) {
	var reqQuery RecordQuery
	// get请求的query表单数据绑定到结构体
	err := c.ShouldBindQuery(&reqQuery)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// 删除内存数据库中记录内的元素
	var newDatabase = []cache.NetInfoInMemoryData{}
	database, ok := cache.Get(cache.Database)
	if ok {
		found := false
		for _, d := range database.([]cache.NetInfoInMemoryData) {
			if d.ID == reqQuery.ID {
				found = true
			} else {
				newDatabase = append(newDatabase, d)
			}
		}
		if found {
			cache.Set(cache.Database, newDatabase)
			return
		}
	}
	c.Status(http.StatusNotFound)
}

// DeleteRecordAll 删除所有记录
func DeleteRecordAll(c *gin.Context) {
	cache.Set(cache.Database, []cache.NetInfoInMemoryData{})
}
