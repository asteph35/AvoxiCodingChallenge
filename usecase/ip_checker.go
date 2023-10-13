package usecase

import (
	"AvoxiCodingChallenge/models"
	"net"
)

type (
	IPManager struct {
		repo IPRepo
	}
)

func NewIPManager(r IPRepo) IPManager {
	return IPManager{repo: r}
}

// IPChecker takes an ip address and list of countries and checks whether that ip is in each countries CIDR range or not
func (ipm IPManager) IPChecker(ipAddress string, countries ...string) (models.CountryMap, error) {
	address := net.ParseIP(ipAddress)
	if address == nil {
		return nil, InvalidIP
	}

	countryMap := make(map[string]bool)
	for _, country := range countries {
		networks, err := ipm.repo.GetNetworks(country)
		if err != nil {
			return nil, err
		}
		countryMap[country] = false
		for _, network := range networks {
			_, cidrRange, err := net.ParseCIDR(network)
			if err != nil {
				return nil, err
			}
			if cidrRange.Contains(address) {
				countryMap[country] = true
			}
		}
	}
	return countryMap, nil
}
