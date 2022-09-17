package service

import (
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/jmoiron/sqlx"
)

type LaunchDBRequester interface {
	SaveLaunch(dbExec sqlx.ExtContext, l entities.Launch) error
}

type UserDBRequester interface {
	SaveUser(dbExec sqlx.ExtContext, l entities.User) (entities.User, error)
}

type SpaceLauncherService struct {
	DBCon      *sqlx.DB
	LaunchRepo LaunchDBRequester
	UserRepo   UserDBRequester
}

// TODO: call SpaceX API to validate

func (sls SpaceLauncherService) CreateBooking(u entities.User, l entities.Launch) error {

	return nil
}
