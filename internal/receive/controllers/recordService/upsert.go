package recordService

import (
	"github.com/gek64/gek/gNet"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"net/netip"
	"netinfo/internal/cache"
	"time"
)

// UpsertRecord 内容记录
func UpsertRecord(c *gin.Context) {
	var reqBody RecordBody
	var data cache.NetInfoInMemoryData

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
		return
	}
	requestIP, err := netip.ParseAddr(requestIPString)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// 组装更新记录结构体
	data.ID = reqBody.ID
	data.NetInterfaces = reqBody.NetInterfaces
	data.RequestIP = requestIP
	data.UpdatedAt = time.Now()

	// 更新内存数据库中记录内的元素
	var newDatabase []cache.NetInfoInMemoryData
	database, ok := cache.Get(cache.Database)
	if ok {
		for _, d := range database.([]cache.NetInfoInMemoryData) {
			if d.ID == data.ID {
				continue
			} else {
				newDatabase = append(newDatabase, d)
			}
		}
	}
	newDatabase = append(newDatabase, data)
	cache.Set(cache.Database, newDatabase)
}
