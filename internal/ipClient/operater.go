package ipClient

import (
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"github.com/gek64/gek/gNet"
	"github.com/imroc/req/v3"
	"log"
	"net/netip"
	"netinfo/ent"
	"netinfo/ent/schema"
	"netinfo/internal/controllers/recordService"
	"time"
)

func GetRequestBody(description string) (updateRecordBody recordService.RecordBody, err error) {
	netInterfaces, err := GetNetInterfaces()
	if err != nil {
		return recordService.RecordBody{}, err
	}

	deviceID, err := machineid.ID()
	if err != nil {
		return recordService.RecordBody{}, err
	}

	updateRecordBody.ID = deviceID
	updateRecordBody.Description = description
	updateRecordBody.NetInterfaces = netInterfaces

	return updateRecordBody, err
}

func SendRequest(url string, description string, username string, password string, skipCertVerify bool) (result ent.Record, err error) {
	client := req.C()

	// 默认不启用跳过TLS证书检测
	if skipCertVerify {
		client.EnableInsecureSkipVerify()
	}

	// 组装好发送POST请求的Body
	body, err := GetRequestBody(description)
	if err != nil {
		return ent.Record{}, err
	}

	// 发送POST请求
	resp, err := client.R().
		SetBody(body).
		SetSuccessResult(&result).
		SetRetryCount(3).
		SetRetryBackoffInterval(1*time.Second, 5*time.Second).
		SetBasicAuth(username, password).
		Put(url)
	if err != nil {
		return ent.Record{}, err
	}

	// 返回值
	if resp.IsSuccessState() {
		return result, nil
	} else {
		return ent.Record{}, fmt.Errorf(resp.ToString())
	}
}

func SendRequestLoop(url string, interval time.Duration, description string, username string, password string, skipCertVerify bool) {
	for {
		_, err := SendRequest(url, description, username, password, skipCertVerify)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("update completed")
		}

		time.Sleep(interval)
	}
}

func GetNetInterfaces() (netInterfaces []schema.NetInterface, err error) {
	nis, err := gNet.GetNetInterfaces()
	if err != nil {
		return nil, err
	}

	for _, ni := range nis {
		// 只取网路接口标记为UP的
		if !ni.Flag.Up {
			continue
		}

		// 拼接转换网络接口中的IPV4 与IPV6地址
		var ips []netip.Addr
		for _, ipString := range append(ni.Ipv4, ni.Ipv6...) {
			// 跳过回环地址
			if isLoopback, _ := gNet.IsLoopback(ipString); isLoopback {
				continue
			}
			// 跳过地址转换出错
			ipAddr, err := netip.ParseAddr(ipString)
			if err != nil {
				continue
			}

			ips = append(ips, ipAddr)
		}

		// 跳过回环网络接口
		if len(ips) > 0 {
			netInterfaces = append(netInterfaces, schema.NetInterface{
				Name: ni.Name,
				IPs:  ips,
				Mac:  ni.Mac,
			})
		}
	}

	return netInterfaces, nil
}

func PrintNetInterfaces() (err error) {
	netInterfaces, err := GetNetInterfaces()
	if err != nil {
		return err
	}
	for i, netInterface := range netInterfaces {
		fmt.Printf("interface: %d\nname: %s\nmac: %s\nips: %v\n", i, netInterface.Name, netInterface.Mac, netInterface.IPs)
	}
	return nil
}
