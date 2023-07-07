package queryService

import (
	"github.com/gek64/gek/gNet"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/netip"
)

type Response struct {
	RequestIP netip.Addr `json:"requestIP"`
	UserAgent string     `json:"userAgent"`
}

// Query 查询
func Query(c *gin.Context) {
	// 获取访问IP
	requestIPString, err := gNet.GetIPFromRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	//访问IP字符串验证并转换为IP
	requestIP, err := netip.ParseAddr(requestIPString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// 获取用户代理
	ua := c.GetHeader("User-Agent")

	c.JSON(http.StatusOK, Response{RequestIP: requestIP, UserAgent: ua})
}
