package repo

import (
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/jmoiron/sqlx"
)

type LauncherRepo struct {
}

func (lr LauncherRepo) SaveLaunch(dbExec sqlx.ExtContext, l entities.Launch) error {
	return nil
}
