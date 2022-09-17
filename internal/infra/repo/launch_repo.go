package repo

import (
	"context"
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/jmoiron/sqlx"
)

type LauncherRepo struct {
}

func (lr LauncherRepo) SaveLaunch(ctx context.Context, dbExec sqlx.ExtContext, u entities.User, l entities.Launch) error {
	return nil
}
