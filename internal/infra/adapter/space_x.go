package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/DenisGoldiner/space_launcher/pkg"
	"io"
	"net/http"
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

func (sxc SpaceXClient) GetPlannedLaunches(ctx context.Context, launchpadID string, timeRange entities.TimeRange) ([]entities.Launch, error) {
	endpoint := fmt.Sprintf(`%s/v5/launches/query`, spaceXURL)

	queryOptions := QueryOptions{
		Query{
			LaunchpadID: launchpadID,
			DateUTC:     TimeRange(timeRange),
		},
	}

	body, err := json.Marshal(queryOptions)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, pkg.WrapErr("create request: %w", err)
	}

	responseBody, err := sxc.Client.SendRequest(req)
	if err != nil {
		return nil, pkg.WrapErr("failed execute the request", err)
	}

	b, err := io.ReadAll(responseBody)
	if err != nil {
		return nil, pkg.WrapErr("failed to read the response body", err)
	}

	var foundLaunches LaunchQueryResponse

	if err := json.Unmarshal(b, &foundLaunches); err != nil {
		return nil, pkg.WrapErr("failed to deserialize the response", err)
	}

	launches := make([]entities.Launch, len(foundLaunches.Docs))
	for i, foundLaunch := range foundLaunches.Docs {
		launches[i] = entities.Launch{
			ID:          foundLaunch.ID,
			LaunchpadID: foundLaunch.LaunchpadID,
			LaunchDate:  foundLaunch.DateUTC,
		}
	}

	return launches, nil
}
