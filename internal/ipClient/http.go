package ipClient

import (
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"github.com/imroc/req/v3"
	"netinfo/ent"
)

type Resp struct {
	Mes string `json:"mes" form:"mes"`
}

func Create(url string, description string) (result ent.Record, err error) {
	client := req.C()
	// 组装好发送POST请求的Body
	body, err := GetCreateRecordBody(description)
	if err != nil {
		return ent.Record{}, err
	}
	// 发送POST请求
	resp, err := client.R().
		SetBody(body).
		SetSuccessResult(&result).
		Post(url)
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

func Update(url string, description string) (result Resp, err error) {
	client := req.C()
	// 组装好发送PUT请求的Body
	body, err := GetUpdateRecordBody(description)
	if err != nil {
		return Resp{}, err
	}
	// 发送POST请求
	resp, err := client.R().
		SetBody(body).
		SetSuccessResult(&result).
		Put(url)
	if err != nil {
		return Resp{}, err
	}

	// 返回值
	if resp.IsSuccessState() {
		return result, nil
	} else {
		return Resp{}, fmt.Errorf(resp.ToString())
	}
}

func Delete(url string) (result Resp, err error) {
	client := req.C()

	deviceID, err := machineid.ID()
	if err != nil {
		return Resp{}, err
	}

	// 发送DELETE请求
	resp, err := client.R().
		SetQueryParam("id", deviceID).
		SetSuccessResult(&result).
		Delete(url)
	if err != nil {
		return Resp{}, err
	}

	// 返回值
	if resp.IsSuccessState() {
		return result, nil
	} else {
		return Resp{}, fmt.Errorf(resp.ToString())
	}
}
