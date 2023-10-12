package usecase

import "AvoxiCodingChallenge/repo"

type (
	IPManager struct {
		repo repo.Repo
	}
)

func (ipm IPManager) IPChecker(ip string, countries []string) bool {
	for _, country := range countries {
		_, _ = ipm.repo.GetNetwork(country)
	}
	return false

}
