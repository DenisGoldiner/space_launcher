package service

import (
	"context"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/DenisGoldiner/space_launcher/pkg"
)

type LaunchDBRequester interface {
	SaveLaunch(ctx context.Context, dbExec sqlx.ExtContext, u entities.User, l entities.Launch) error
	GetAllLaunches(ctx context.Context, dbExec sqlx.ExtContext) ([]entities.Launch, error)
}

type UserDBRequester interface {
	SaveUser(ctx context.Context, dbExec sqlx.ExtContext, u entities.User) (entities.User, error)
	GetAllUsers(ctx context.Context, dbExec sqlx.ExtContext) ([]entities.User, error)
}

// TODO: rename it to something abstract

type SpaceXAdapter interface {
	GetLaunchpad(ctx context.Context, launchpadID string) (entities.Launchpad, error)
	GetPlannedLaunches(ctx context.Context, launchpadID string, timeRange entities.TimeRange) ([]entities.Launch, error)
}

type SpaceLauncherService struct {
	DBCon        *sqlx.DB
	LaunchRepo   LaunchDBRequester
	UserRepo     UserDBRequester
	SpaceXClient SpaceXAdapter
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

// TODO: validate unique destinations per week

func (sls SpaceLauncherService) CreateBooking(ctx context.Context, u entities.User, l entities.Launch) error {
	if err := sls.validateBooking(ctx, l); err != nil {
		return err
	}

	return sls.createBooking(ctx, u, l)
}

func (sls SpaceLauncherService) validateBooking(ctx context.Context, l entities.Launch) error {
	foundLaunchpad, err := sls.SpaceXClient.GetLaunchpad(ctx, l.LaunchpadID)
	if err != nil {
		return err
	}

	if foundLaunchpad.Status == entities.LaunchpadStatusRetired {
		return errors.New("can not plan booking for retired launchpad")
	}

	timeRange := entities.ToDayRange(l.LaunchDate)
	plannedExternalLaunches, err := sls.SpaceXClient.GetPlannedLaunches(ctx, l.LaunchpadID, timeRange)
	if err != nil {
		return err
	}

	log.Printf("%#v", plannedExternalLaunches)

	if len(plannedExternalLaunches) > 0 {
		return errors.New("the date is planned by an external vendor")
	}

	return nil
}

func (sls SpaceLauncherService) createBooking(ctx context.Context, u entities.User, l entities.Launch) error {
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
