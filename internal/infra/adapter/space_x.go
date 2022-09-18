package adapter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/DenisGoldiner/space_launcher/pkg"
	"io"
	"net/http"
	"time"
)

const (
	spaceXURL = "https://api.spacexdata.com"
)

type SpaceXClient struct {
	Client pkg.Client
}

func (sxc SpaceXClient) GetLaunchpad(ctx context.Context, launchpadID string) (entities.Launchpad, error) {
	endpoint := fmt.Sprintf(`%s/v4/launchpads/%s`, spaceXURL, launchpadID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, http.NoBody)
	if err != nil {
		return entities.Launchpad{}, pkg.WrapErr("create request: %w", err)
	}

	responseBody, err := sxc.Client.SendRequest(req)
	if err != nil {
		return entities.Launchpad{}, pkg.WrapErr("failed execute the request", err)
	}

	b, err := io.ReadAll(responseBody)
	if err != nil {
		return entities.Launchpad{}, pkg.WrapErr("failed to read the response body", err)
	}

	var launchpad Launchpad
	if err := json.Unmarshal(b, &launchpad); err != nil {
		return entities.Launchpad{}, pkg.WrapErr("failed to deserialize the response", err)
	}

	return entities.Launchpad(launchpad), nil
}

func (sxc SpaceXClient) GetConflictingLaunches(ctx context.Context, launchpadID string, date time.Time) (entities.Launch, error) {
	return entities.Launch{}, errors.New("not implemented")
}
