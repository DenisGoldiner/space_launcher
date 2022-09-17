package repo

import (
	"context"
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/jmoiron/sqlx"
	"time"
)

// LaunchEntity represents the launch DB entity.
type LaunchEntity struct {
	ID          string    `db:"id"`
	LaunchpadID string    `db:"launchpad_id"`
	Destination string    `db:"destination"`
	LaunchDate  time.Time `db:"launch_date"`
	UserID      string    `db:"user_id"`
}

type LaunchRepo struct{}

func (lr LaunchRepo) GetAllLaunches(ctx context.Context, dbExec sqlx.ExtContext) ([]entities.Launch, error) {
	getAllLaunchesQuery := `SELECT id, launchpad_id, destination, launch_date, user_id FROM "launch"`

	rows, err := dbExec.QueryxContext(ctx, getAllLaunchesQuery)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var allLaunches []entities.Launch

	if rows.Next() {
		var launch entities.Launch
		if err := rows.StructScan(&launch); err != nil {
			return nil, err
		}

		allLaunches = append(allLaunches, launch)
	}

	return allLaunches, nil
}

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
