package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/DenisGoldiner/space_launcher/pkg"
	"github.com/jmoiron/sqlx"
	"time"
)

type LaunchDBRequester interface {
	SaveLaunch(context.Context, sqlx.ExtContext, entities.User, entities.Launch) error
	GetAllLaunches(context.Context, sqlx.ExtContext) ([]entities.Launch, error)
	GetPadLaunches(context.Context, sqlx.ExtContext, entities.LaunchpadID, entities.TimeRange) ([]entities.Launch, error)
}

type UserDBRequester interface {
	SaveUser(context.Context, sqlx.ExtContext, entities.User) (entities.User, error)
	GetAllUsers(context.Context, sqlx.ExtContext) ([]entities.User, error)
}

// TODO: rename it to something abstract

type SpaceXAdapter interface {
	GetLaunchpad(context.Context, entities.LaunchpadID) (entities.Launchpad, error)
	GetPlannedLaunches(context.Context, entities.LaunchpadID, entities.TimeRange) ([]entities.Launch, error)
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
	if sls.isValidationErrors(err) {
		return pkg.WrapErr(err.Error(), BusinessValidationErr)
	}

	if err != nil {
		return err
	}

	return sls.createBooking(ctx, u, l)
}

func (sls SpaceLauncherService) isValidationErrors(err error) bool {
	if err == nil {
		return false
	}

	validationErrors := [...]error{RetiredLaunchpadErr, TakenDateErr, TakenDestinationErr}

	for _, validationError := range validationErrors {
		if errors.Is(err, validationError) {
			return true
		}
	}

	return false
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
	timeRange := entities.ToMiddleWeekRange(l.LaunchDate)
	foundLaunches, err := sls.LaunchRepo.GetPadLaunches(ctx, sls.DBCon, l.LaunchpadID, timeRange)
	if err != nil {
		return err
	}

	if len(foundLaunches) == 0 {
		return nil
	}

	for _, fl := range foundLaunches {
		if l.LaunchDate.Equal(fl.LaunchDate) {
			return pkg.WrapErr("internal booking", TakenDateErr)
		}

		if l.Destination == fl.Destination {
			return pkg.WrapErr(fmt.Sprintf("exists for %s", fl.LaunchDate.Format(time.RFC3339)), TakenDestinationErr)
		}
	}

	return nil
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
