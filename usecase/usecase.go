package usecase

type IPRepo interface {
	GetNetworks(country string) ([]string, error)
}
