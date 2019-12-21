package repository

import (
	"database/sql"
	"github.com/vskut/twigo/pkg/common/entity"
	"github.com/vskut/twigo/pkg/common/repository/postgresql"
)

// Repository ...
type Repository struct {
	Users  entity.UsersRepository
	Tweets entity.TweetsRepository
}

// NewPostgreSQLRepository ...
func NewPostgreSQLRepository(db *sql.DB) *Repository {
	return &Repository{
		Users:  postgresql.NewUsersRepositoryPostgreSQL(db),
		Tweets: postgresql.NewTweetsRepositoryPostgreSQL(db),
	}
}
