package service

import (
	"context"
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/DenisGoldiner/space_launcher/pkg"
	"github.com/jmoiron/sqlx"
)

type LaunchDBRequester interface {
	SaveLaunch(ctx context.Context, dbExec sqlx.ExtContext, u entities.User, l entities.Launch) error
}

type UserDBRequester interface {
	SaveUser(ctx context.Context, dbExec sqlx.ExtContext, u entities.User) (entities.User, error)
}

type SpaceLauncherService struct {
	DBCon      *sqlx.DB
	LaunchRepo LaunchDBRequester
	UserRepo   UserDBRequester
}

// TODO: call SpaceX API to validate

// TODO: validate unique destinations per week

func (sls SpaceLauncherService) CreateBooking(ctx context.Context, u entities.User, l entities.Launch) error {
	tx, err := sls.DBCon.BeginTxx(ctx, nil)
	if err != nil {
		return pkg.WrapErr("failed to start the transaction", err)
	}
	defer func() { _ = tx.Rollback() }()

	if err := sls.createBookingTx(ctx, tx, u, l); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return pkg.WrapErr("failed to commit transaction: %w", err)
	}

	return nil
}

func (sls SpaceLauncherService) createBookingTx(ctx context.Context, tx *sqlx.Tx, u entities.User, l entities.Launch) error {
	savedUser, err := sls.UserRepo.SaveUser(ctx, tx, u)
	if err != nil {
		return err
	}

	if err := sls.LaunchRepo.SaveLaunch(ctx, tx, savedUser, l); err != nil {
		return err
	}

	return nil
}
