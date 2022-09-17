package adapter

import (
	"context"
	"errors"
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"time"
)

type SpaceXClient struct{}

func (sxc SpaceXClient) GetLaunchpad(ctx context.Context, launchpadID string) ([]entities.Launchpad, error) {
	return nil, errors.New("not implemented")
}

func (sxc SpaceXClient) GetConflictingLaunches(ctx context.Context, launchpadID string, date time.Time) (entities.Launch, error) {
	return entities.Launch{}, errors.New("not implemented")
}
