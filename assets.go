package railigentx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AssetCollectionResponse represents the response structure for listing assets in a fleet
type AssetCollectionResponse struct {
	Data []Asset `json:"data"`
}

// AssetResponse represents the response structure for a single asset
type AssetResponse struct {
	Data Asset `json:"data"`
}

// Asset represents an asset in the Railigent X API response
type Asset struct {
	ID       string   `json:"id"`
	Features Features `json:"features"` // Can be further defined based on the API's asset feature structure
}

type Features struct {
	UIC     *UIC     `json:"uic"`
	GPS     *GPS     `json:"gps"`
	Mileage *Mileage `json:"Mileage"`
	Speed   *Speed   `json:"speed"`
	Trip    *Trip    `json:"trip"`
}

type UIC struct {
	Value string `json:"value"`
}

type GPS struct {
	Timestamp int64    `json:"timestamp"`
	Position  Position `json:"value"`
}

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Mileage struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type Speed struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type Trip struct {
	Timestamp int64 `json:"timestamp"`
	Value     struct {
		TripID string `json:"tripId"`
	} `json:"value"`
}

func (c *Client) ListAssets(fleetId string) (*AssetCollectionResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/fleets/%s/assets?features.gps,features.mileage,features.speed,features.trip,features.uic", c.BaseURL, fleetId), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request for ListAssets: %w", err)
	}

	req.Header.Add("Authorization", c.AuthHeader)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to list assets: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status for ListAssets: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body for ListAssets: %w", err)
	}

	var assetsResponse AssetCollectionResponse
	if err := json.Unmarshal(body, &assetsResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response for ListAssets: %w", err)
	}

	return &assetsResponse, nil
}

// GetAsset retrieves information for a specific asset of a specific fleet.
func (c *Client) GetAsset(fleetId string, assetId string) (*AssetResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/fleets/%s/assets/%s?features.gps,features.mileage,features.speed,features.trip,features.uic", c.BaseURL, fleetId, assetId), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request for GetAsset: %w", err)
	}

	req.Header.Add("Authorization", c.AuthHeader)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to get asset: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status for GetAsset: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body for GetAsset: %w", err)
	}

	var assetResponse AssetResponse
	if err := json.Unmarshal(body, &assetResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response for GetAsset: %w", err)
	}

	return &assetResponse, nil
}
