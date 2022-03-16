package main

type NetInfo struct {
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
