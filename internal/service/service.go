package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"

	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/DenisGoldiner/space_launcher/pkg"
)

type LaunchDBRequester interface {
	SaveLaunch(ctx context.Context, dbExec sqlx.ExtContext, u entities.User, l entities.Launch) error
	GetAllLaunches(ctx context.Context, dbExec sqlx.ExtContext) ([]entities.Launch, error)
	GetLaunch(ctx context.Context, dbExec sqlx.ExtContext, lID string, lDate time.Time) (entities.Launch, error)
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
			return nil, errors.New("this could not ever happen but it happened, DB data inconsistency")
		}

		allBookings[user] = append(allBookings[user], launch)
	}

	return allBookings, nil
}

func (sls SpaceLauncherService) CreateBooking(ctx context.Context, u entities.User, l entities.Launch) error {
	err := sls.validateBooking(ctx, l)
	if errors.Is(err, RetiredLaunchpadErr) || errors.Is(err, TakenDateErr) {
		return pkg.WrapErr(err.Error(), BusinessValidationErr)
	}

	if err != nil {
		return err
	}

	return sls.createBooking(ctx, u, l)
}

func (sls SpaceLauncherService) validateBooking(ctx context.Context, l entities.Launch) error {
	if err := sls.validateLaunchpadReadiness(ctx, l); err != nil {
		return err
	}

	if err := sls.validateInternalBookings(ctx, l); err != nil {
		return err
	}

	if err := sls.validateExternalBookings(ctx, l); err != nil {
		return err
	}

	return nil
}

func (sls SpaceLauncherService) validateInternalBookings(ctx context.Context, l entities.Launch) error {
	_, err := sls.LaunchRepo.GetLaunch(ctx, sls.DBCon, l.LaunchpadID, l.LaunchDate)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}

	if err != nil {
		return err
	}

	return pkg.WrapErr("internal booking", TakenDateErr)
}

func (sls SpaceLauncherService) validateLaunchpadReadiness(ctx context.Context, l entities.Launch) error {
	foundLaunchpad, err := sls.SpaceXClient.GetLaunchpad(ctx, l.LaunchpadID)
	if err != nil {
		return pkg.WrapErr(fmt.Sprintf("for getting the launchpad, %v", err), ExternalVendorAPIErr)
	}

	if foundLaunchpad.Status == entities.LaunchpadStatusRetired {
		return RetiredLaunchpadErr
	}

	return nil
}

func (sls SpaceLauncherService) validateExternalBookings(ctx context.Context, l entities.Launch) error {
	timeRange := entities.ToDayRange(l.LaunchDate)
	plannedExternalLaunches, err := sls.SpaceXClient.GetPlannedLaunches(ctx, l.LaunchpadID, timeRange)
	if err != nil {
		return pkg.WrapErr(fmt.Sprintf("for getting the launches, %v", err), ExternalVendorAPIErr)
	}

	if len(plannedExternalLaunches) > 0 {
		return pkg.WrapErr("external booking", TakenDateErr)
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
		return pkg.WrapErr("failed to commit transaction", err)
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
