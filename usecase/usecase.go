package usecase

type IPRepo interface {
	GetNetwork(country string) (string, error)
}
