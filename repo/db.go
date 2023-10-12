package repo

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repo struct {
	write *pgxpool.Pool
	read  *pgxpool.Pool
}

func New(write, read *pgxpool.Pool) *Repo {
	if read == nil {
		read = write
	}
	return &Repo{
		write: write,
		read:  read,
	}
}
