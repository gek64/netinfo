package queryController

import (
	"encoding/json"
	"github.com/gek64/gek/gNet"
	"github.com/gin-gonic/gin"
	"io"
	"net"
	"net/http"
)

type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
}

type Resp struct {
	IP        net.Addr `json:"ip,omitempty"`
	IPInfo    IPInfo   `json:"ipInfo,omitempty"`
	UserAgent string   `json:"userAgent"`
}

func getIPInfoFromIpInfoIo(ip string) (ipInfo IPInfo, err error) {
	response, err := http.Get("https://ipinfo.io/" + ip + "/json")
	if err != nil {
		return IPInfo{}, err
	}

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return IPInfo{}, err
	}

	err = json.Unmarshal(respBytes, &ipInfo)
	if err != nil {
		return IPInfo{}, err
	}

	return ipInfo, err
}

//func respBodyBuilder(req *http.Request, useragent string, useNetDB bool) (respBody []byte, err error) {
//	// 获取req中真实ip
//	ip, err := gNet.GetIPFromRequest(req)
//	if err != nil {
//		return nil, err
//	}
//
//	// 如果要使用网络ip数据库
//	if useNetDB {
//		ipInfo, err := getIPInfoFromIpInfoIo(ip)
//		if err == nil {
//			respData := Resp{IPInfo: ipInfo}
//			respData.UserAgent = useragent
//			// 转换为json []byte
//			respBody, err = json.Marshal(respData)
//			if err != nil {
//				return nil, err
//			}
//			return respBody, nil
//		}
//	}
//
//	// 不使用网络ip数据库,或者连接网络ip数据库失败
//	respData := RespMin{IP: ip}
//	respData.UserAgent = useragent
//	// 转换为json []byte
//	respBody, err = json.Marshal(respData)
//	if err != nil {
//		return nil, err
//	}
//	return respBody, nil
//}

// Query 查询
func Query(c *gin.Context) {
	requestIP, err := gNet.GetIPFromRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
	}
	
	c.JSON(http.StatusOK, gin.H{"ip": requestIP, "user-agent": c.GetHeader("User-Agent")})
}
