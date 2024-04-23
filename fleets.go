package railigentx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// FleetCollectionResponse represents the response structure for listing fleets
type FleetCollectionResponse struct {
	Data []Fleet `json:"data"`
}

type FleetResponse struct {
	Data Fleet `json:"data"`
}

// Fleet represents a fleet in the Railigent X API response
type Fleet struct {
	ID       string   `json:"id"`
	AssetIds []string `json:"assetIds"`
}

// ListFleets retrieves information about all eligible fleets.
func (c *Client) ListFleets() (*FleetCollectionResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/fleets", c.BaseURL), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Authorization", c.AuthHeader)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to list fleets: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var fleetsResponse FleetCollectionResponse
	if err := json.Unmarshal(body, &fleetsResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return &fleetsResponse, nil
}

// GetFleet retrieves information about a specific eligible fleet.
func (c *Client) GetFleet(fleetId string) (*FleetResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/fleets/%s", c.BaseURL, fleetId), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request for GetFleet: %w", err)
	}

	req.Header.Add("Authorization", c.AuthHeader)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to get fleet: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status for GetFleet: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body for GetFleet: %w", err)
	}

	var fleetResponse FleetResponse
	if err := json.Unmarshal(body, &fleetResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response for GetFleet: %w", err)
	}

	return &fleetResponse, nil
}
