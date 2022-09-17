package repo

import (
	"context"
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct{}

func (ur UserRepo) SaveUser(ctx context.Context, dbExec sqlx.ExtContext, u entities.User) (entities.User, error) {
	return entities.User{}, nil
}
