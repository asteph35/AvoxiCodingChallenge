package models

type CountryMap map[string]bool

type IPCheckerReq struct {
	IP        string   `json:"ipAddress,omitempty"`
	Countries []string `json:"countries,omitempty"`
}

type IPCheckerRes struct {
	IP        string          `json:"ipAddress,omitempty"`
	Countries []CountryStatus `json:"countries,omitempty"`
}
type CountryStatus struct {
	Country string `json:"country,omitempty"`
	Status  bool   `json:"ipInCountry"`
}

func NewIPCheckerReq(ip string, countries ...string) IPCheckerReq {
	return IPCheckerReq{
		IP:        ip,
		Countries: countries,
	}
}

func NewIPCheckerRes(ip string, countryMap CountryMap) IPCheckerRes {
	var statuses []CountryStatus
	for c, s := range countryMap {
		statuses = append(statuses, CountryStatus{
			Country: c,
			Status:  s,
		})
	}
	return IPCheckerRes{
		IP:        ip,
		Countries: statuses,
	}
}
