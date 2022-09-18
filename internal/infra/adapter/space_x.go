package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DenisGoldiner/space_launcher/internal/entities"
	"github.com/DenisGoldiner/space_launcher/pkg"
)

const (
	spaceXURL = "https://api.spacexdata.com"
)

// SpaceXClient to interact with external service.
type SpaceXClient struct {
	Client pkg.Client
}

// GetLaunchpad retrieves Launchpad by ID from the SpaceX planer.
func (sxc SpaceXClient) GetLaunchpad(ctx context.Context, lID entities.LaunchpadID) (entities.Launchpad, error) {
	endpoint := fmt.Sprintf(`%s/v4/launchpads/%s`, spaceXURL, lID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, http.NoBody)
	if err != nil {
		return entities.Launchpad{}, pkg.WrapErr("failed to create the request", err)
	}

	statusCode, responseBody, err := sxc.Client.SendRequest(req)
	if statusCode == http.StatusNotFound {
		return entities.Launchpad{}, nil
	}

	if err != nil {
		return entities.Launchpad{}, pkg.WrapErr("failed execute the request", err)
	}

	defer func() { _ = responseBody.Close() }()

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

// GetPlannedLaunches retrieves Launches list by Launchpad ID and date range from the SpaceX planer.
func (sxc SpaceXClient) GetPlannedLaunches(
	ctx context.Context,
	lID entities.LaunchpadID,
	timeRange entities.TimeRange,
) ([]entities.Launch, error) {
	endpoint := fmt.Sprintf(`%s/v5/launches/query`, spaceXURL)

	queryOptions := QueryOptions{
		Query{
			LaunchpadID: lID,
			DateUTC:     TimeRange(timeRange),
		},
	}

	body, err := json.Marshal(queryOptions)
	if err != nil {
		return nil, pkg.WrapErr("failed to encode the request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, pkg.WrapErr("failed to create the request", err)
	}

	_, responseBody, err := sxc.Client.SendRequest(req)
	if err != nil {
		return nil, pkg.WrapErr("failed execute the request", err)
	}

	defer func() { _ = responseBody.Close() }()

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
