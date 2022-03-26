package main

import (
	"fmt"
	"os"
)

var (
	ASNDB = []string{
		"GeoLite2-ASN-Blocks-IPv4",
		"GeoLite2-ASN-Blocks-IPv6",
	}
	CITYDB = []string{
		"GeoLite2-City-Blocks-IPv4",
		"GeoLite2-City-Blocks-IPv6",
		"GeoLite2-City-Locations",
	}
	COUNTRYDB = []string{
		"GeoLite2-Country-Blocks-IPv4",
		"GeoLite2-Country-Blocks-IPv6",
		"GeoLite2-Country-Locations",
	}
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

func checkDatabase(databaseLocation string) (ready bool, err error) {
	dirEntries, err := os.ReadDir(databaseLocation)
	if err != nil {
		return false, err
	}

	for _, entry := range dirEntries {
		fmt.Println(entry.Name())
	}

	return
}
