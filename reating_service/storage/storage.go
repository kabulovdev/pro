package storage

import (
	"gitlab.com/pro/reating_service/storage/postgres"
	"gitlab.com/pro/reating_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Reating() repo.ReatingInfoI
}

type storagePg struct {
	db          *sqlx.DB
	reatingRepo repo.ReatingInfoI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		reatingRepo: postgres.NewPostRepo(db),
	}
}
func (s storagePg) Reating() repo.ReatingInfoI {
	return s.reatingRepo
}
