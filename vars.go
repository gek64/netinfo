package main

type IPInfoMax struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
}

type IPInfoMin struct {
	IP string `json:"ip"`
}

type RespMax struct {
	IPInfo    IPInfoMax
	UserAgent string
}

type RespMin struct {
	IPInfo    IPInfoMin
	UserAgent string
}
