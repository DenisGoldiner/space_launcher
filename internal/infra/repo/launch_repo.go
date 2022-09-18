package repo

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/DenisGoldiner/space_launcher/internal/entities"
)

// LaunchEntity represents the launch DB entity.
type LaunchEntity struct {
	ID          string    `db:"id"`
	LaunchpadID string    `db:"launchpad_id"`
	Destination string    `db:"destination"`
	LaunchDate  time.Time `db:"launch_date"`
	UserID      string    `db:"user_id"`
}

func (le LaunchEntity) toEntitiesLaunch() entities.Launch {
	return entities.Launch{
		ID:          le.ID,
		LaunchpadID: le.LaunchpadID,
		Destination: entities.Destination(le.Destination),
		LaunchDate:  le.LaunchDate,
		UserID:      le.UserID,
	}
}

// LaunchRepo is a repo layer implementation for launch actions.
type LaunchRepo struct{}

// GetAllLaunches returns all launches from DB.
func (lr LaunchRepo) GetAllLaunches(ctx context.Context, dbExec sqlx.ExtContext) ([]entities.Launch, error) {
	getAllLaunchesQuery := `SELECT id, launchpad_id, destination, launch_date, user_id FROM "launch"`

	rows, err := dbExec.QueryxContext(ctx, getAllLaunchesQuery)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	return scanLaunches(rows)
}

// GetPadLaunches returns all launches per launchpad in specific date range from DB.
func (lr LaunchRepo) GetPadLaunches(
	ctx context.Context,
	dbExec sqlx.ExtContext,
	launchpadID string,
	launchDateRange entities.TimeRange,
) ([]entities.Launch, error) {
	getAllLaunchesQuery := `
		SELECT id, launchpad_id, destination, launch_date, user_id FROM "launch"
		WHERE launchpad_id = $1 AND launch_date >= $2 AND launch_date < $3`

	rows, err := dbExec.QueryxContext(ctx, getAllLaunchesQuery, launchpadID, launchDateRange.From, launchDateRange.To)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	return scanLaunches(rows)
}

func scanLaunches(rows *sqlx.Rows) ([]entities.Launch, error) {
	var allLaunches []entities.Launch

	for rows.Next() {
		var launch LaunchEntity
		if err := rows.StructScan(&launch); err != nil {
			return nil, err
		}

		allLaunches = append(allLaunches, launch.toEntitiesLaunch())
	}

	return allLaunches, nil
}

// SaveLaunch persists a launch to DB.
func (lr LaunchRepo) SaveLaunch(ctx context.Context, dbExec sqlx.ExtContext, u entities.User, l entities.Launch) error {
	saveLaunchQuery := `INSERT INTO "launch" (launchpad_id, destination, launch_date, user_id) VALUES ($1, $2, $3, $4)`
	if _, err := dbExec.ExecContext(
		ctx,
		saveLaunchQuery,
		l.LaunchpadID,
		l.Destination,
		l.LaunchDate,
		u.ID,
	); err == nil {
		return err
	}

	return nil
}

// DeleteLaunch deletes the launch from the DB by launchpad and concrete date.
func (lr LaunchRepo) DeleteLaunch(ctx context.Context, dbExec sqlx.ExtContext, l entities.Launch) error {
	saveLaunchQuery := `DELETE FROM "launch" WHERE launchpad_id = $1 AND launch_date = $2`
	if _, err := dbExec.ExecContext(ctx, saveLaunchQuery, l.LaunchpadID, l.LaunchDate); err == nil {
		return err
	}

	return nil
}
