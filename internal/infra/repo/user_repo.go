package repo

import (
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct{}

func (ur UserRepo) SaveUser(dbExec sqlx.ExtContext, l entities.User) (entities.User, error) {
	return entities.User{}, nil
}
