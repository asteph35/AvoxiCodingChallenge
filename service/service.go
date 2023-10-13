package service

import (
	"AvoxiCodingChallenge/repository"
	"AvoxiCodingChallenge/usecase"
)

type Service struct {
	Repo repository.Repo
	usecase.IPManager
}

func New(repo repository.Repo) Service {
	return Service{
		Repo:      repo,
		IPManager: usecase.NewIPManager(repo),
	}
}
