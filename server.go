package main

import (
	"encoding/json"
	"fmt"
	"gek_net"
	"geolite2"
	"net/http"
	"net/netip"
	"strconv"
)

type Response struct {
	IP IP     `json:"ip"`
	Ua string `json:"ua"`
}
type IP struct {
	Address string `json:"address"`
	As      string `json:"as"`
	Asn     string `json:"asn"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func httpReturnClientNetworkInfo(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		resp.Header().Set("Content-Type", "application/json")

		response, err := getResp(req)
		if err != nil {
			fmt.Println(err)
		}

		// 转换为json []byte
		respBody, err := json.Marshal(response)
		if err != nil {
			http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
		}

		// 写入resp中
		_, err = resp.Write(respBody)
		if err != nil {
			http.Error(resp, "Internal Server Error", http.StatusInternalServerError)
		}
	default:
		http.Error(resp, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func getResp(request *http.Request) (resp Response, err error) {
	var (
		asnResult geolite2.AsnResult
		cbResult  geolite2.CityBlockResult
		clResult  geolite2.CityLocationResult
	)

	ip, err := gek_net.GetIPFromRequest(request)
	if err != nil {
		return Response{}, err
	}
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return Response{}, err
	}

	if addr.Is4() {
		asnResult, err = asnSearch(addr.String(), "database/GeoLite2-ASN-Blocks-IPv4.csv")
		if err != nil {
			return Response{}, err
		}
		cbResult, err = cityBlockSearch(addr.String(), "database/GeoLite2-City-Blocks-IPv4.csv")
		if err != nil {
			return Response{}, err
		}
	} else {
		asnResult, err = asnSearch(addr.String(), "database/GeoLite2-ASN-Blocks-IPv6.csv")
		if err != nil {
			return Response{}, err
		}
		cbResult, err = cityBlockSearch(addr.String(), "database/GeoLite2-City-Blocks-IPv6.csv")
		if err != nil {
			return Response{}, err
		}
	}

	clResult, err = cityLocationSearch(cbResult.GeoNameID, "database/GeoLite2-City-Locations-en.csv")
	if err != nil {
		return Response{}, err
	}

	resp.IP.Address = addr.String()
	resp.IP.As = strconv.Itoa(asnResult.Asn)
	resp.IP.Asn = asnResult.AsnOrg
	resp.IP.City = clResult.CityName
	resp.IP.Country = clResult.CountryName
	resp.Ua = request.Header.Values("User-Agent")[0]

	return resp, nil
}
