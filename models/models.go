package models

type CountryMap map[string]bool

type IPCheckerReq struct {
	IP        string   `json:"ipAddress,omitempty"`
	Countries []string `json:"countries,omitempty"`
}

type IPCheckerRes struct {
	Countries CountryMap `json:"countries,omitempty"`
}

func NewIPCheckerReq(ip string, countries ...string) IPCheckerReq {
	return IPCheckerReq{
		IP:        ip,
		Countries: countries,
	}
}

func NewIPCheckerRes(countryMap CountryMap) IPCheckerRes {
	return IPCheckerRes{
		countryMap,
	}
}
