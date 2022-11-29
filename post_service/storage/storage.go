package storage

import (
	"gitlab.com/pro/post_service/storage/postgres"
	"gitlab.com/pro/post_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Post() repo.PostInfoI
}

type storagePg struct {
	db       *sqlx.DB
	postRepo repo.PostInfoI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		postRepo: postgres.NewPostRepo(db),
	}
}
func (s storagePg) Post() repo.PostInfoI {
	return s.postRepo
}
