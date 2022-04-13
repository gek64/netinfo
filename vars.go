package main

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
	IPInfo    IPInfo
	UserAgent string
}

type RespMin struct {
	IP        string
	UserAgent string
}
