package service

import (
	"context"
	"errors"
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/DenisGoldiner/space_launcher/pkg"
	"github.com/jmoiron/sqlx"
)

type LaunchDBRequester interface {
	SaveLaunch(ctx context.Context, dbExec sqlx.ExtContext, u entities.User, l entities.Launch) error
	GetAllLaunches(ctx context.Context, dbExec sqlx.ExtContext) ([]entities.Launch, error)
}

type UserDBRequester interface {
	SaveUser(ctx context.Context, dbExec sqlx.ExtContext, u entities.User) (entities.User, error)
	GetAllUsers(ctx context.Context, dbExec sqlx.ExtContext) ([]entities.User, error)
}

type SpaceLauncherService struct {
	DBCon      *sqlx.DB
	LaunchRepo LaunchDBRequester
	UserRepo   UserDBRequester
}

func (sls SpaceLauncherService) GetAllBookings(ctx context.Context) (map[entities.User][]entities.Launch, error) {
	allUsers, err := sls.UserRepo.GetAllUsers(ctx, sls.DBCon)
	if err != nil {
		return nil, err
	}

	allLaunches, err := sls.LaunchRepo.GetAllLaunches(ctx, sls.DBCon)
	if err != nil {
		return nil, err
	}

	userMapping := make(map[string]entities.User, len(allUsers))
	for _, user := range allUsers {
		userMapping[user.ID] = user
	}

	allBookings := make(map[entities.User][]entities.Launch, len(allUsers))

	for _, launch := range allLaunches {
		user, ok := userMapping[launch.UserID]
		if !ok {
			return nil, errors.New("this could not ever happen but it happened, data inconsistency")
		}

		allBookings[user] = append(allBookings[user], launch)
	}

	return allBookings, nil
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
