package usecase

import (
	"AvoxiCodingChallenge/models"
	"strings"
)

type (
	IPManager struct {
		repo IPRepo
	}
)

func NewIPManager(r IPRepo) IPManager {
	return IPManager{repo: r}
}

func (ipm IPManager) IPChecker(ip string, countries ...string) (models.CountryMap, error) {
	countryMap := make(map[string]bool)
	for _, country := range countries {
		networks, err := ipm.repo.GetNetworks(country)
		if err != nil {
			return nil, err
		}
		countryMap[country] = false
		for _, network := range networks {
			lastIndex := strings.LastIndex(ip, ".")
			if lastIndex == -1 {
				return nil, InvalidIP
			}
			if strings.HasPrefix(network, ip[:lastIndex]) {
				countryMap[country] = true
			}
		}
	}
	return countryMap, nil
}
