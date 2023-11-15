package debugService

import (
	"fmt"
	"github.com/gek64/gek/gNet"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/netip"
)

// Reflector 反射器
func Reflector(c *gin.Context) {
	// request ip
	requestIPString, err := gNet.GetIPFromRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	requestIP, err := netip.ParseAddr(requestIPString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// body
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK,
		Response{
			Method:    c.Request.Method,
			RequestIP: requestIP,
			Body:      bodyBytes,
			Query:     c.Request.URL.Query(),
			Header:    c.Request.Header,
		})
}
