package ipClient

import (
	"fmt"
	"github.com/imroc/req/v3"
	"netinfo/ent"
	"strconv"
)

type Resp struct {
	mes string
}

func Create(url string, description string) (result ent.Record, err error) {
	client := req.C()
	// 组装好发送POST请求的Body
	body, err := CreateRecordBody(description)
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
		return ent.Record{}, fmt.Errorf(resp.Status)
	}
}

func Update(url string, id uint, description string) (result Resp, err error) {
	client := req.C()
	// 组装好发送PUT请求的Body
	body, err := UpdateRecordBody(id, description)
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
		return Resp{}, fmt.Errorf(resp.Status)
	}
}

func Delete(url string, id uint) (result Resp, err error) {
	client := req.C()

	// 发送DELETE请求
	resp, err := client.R().
		SetQueryParam("id", strconv.Itoa(int(id))).
		SetSuccessResult(&result).
		Put(url)
	if err != nil {
		return Resp{}, err
	}

	// 返回值
	if resp.IsSuccessState() {
		return result, nil
	} else {
		return Resp{}, fmt.Errorf(resp.Status)
	}
}
