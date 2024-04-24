package railigentx

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListAssets(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the request
		expectedPath := "/fleets/fleetID/assets"
		expectedQuery := "features.gps,features.mileage,features.speed,features.trip,features.uic"
		if r.URL.Path != expectedPath || r.URL.RawQuery != expectedQuery {
			t.Errorf("Expected request path %s with query %s, got %s with query %s", expectedPath, expectedQuery, r.URL.Path, r.URL.RawQuery)
		}
		// Respond with mock data
		_, _ = w.Write([]byte(`{"data":[{"id":"asset1","features":{"gps":{"timestamp":123456789,"value":{"latitude":12.34,"longitude":56.78}},"mileage":{"timestamp":123456789,"value":100},"speed":{"timestamp":123456789,"value":50},"trip":{"timestamp":123456789,"value":{"tripId":"trip1"}},"uic":{"value":"uic1"}}}]}`))
	}))
	defer mockServer.Close()

	// Create a new Railigent X client with the mock server's URL
	client := NewClient(mockServer.URL, "username", "password")

	// Call the method being tested
	assetResponse, err := client.ListAssets("fleetID")
	if err != nil {
		t.Errorf("Error while listing assets: %v", err)
	}

	// Perform assertions on the response
	if len(assetResponse.Data) != 1 {
		t.Errorf("Expected 1 asset, got %d", len(assetResponse.Data))
	}
	// Add more assertions as needed
}

func TestGetAsset(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the request
		expectedPath := "/fleets/fleetID/assets/assetID"
		expectedQuery := "features.gps,features.mileage,features.speed,features.trip,features.uic"
		if r.URL.Path != expectedPath || r.URL.RawQuery != expectedQuery {
			t.Errorf("Expected request path %s with query %s, got %s with query %s", expectedPath, expectedQuery, r.URL.Path, r.URL.RawQuery)
		}
		// Respond with mock data
		_, _ = w.Write([]byte(`{"data":{"id":"asset1","features":{"gps":{"timestamp":123456789,"value":{"latitude":12.34,"longitude":56.78}},"mileage":{"timestamp":123456789,"value":100},"speed":{"timestamp":123456789,"value":50},"trip":{"timestamp":123456789,"value":{"tripId":"trip1"}},"uic":{"value":"uic1"}}}}`))
	}))
	defer mockServer.Close()

	// Create a new Railigent X client with the mock server's URL
	client := NewClient(mockServer.URL, "username", "password")

	// Call the method being tested
	assetResponse, err := client.GetAsset("fleetID", "assetID")
	if err != nil {
		t.Errorf("Error while getting asset: %v", err)
	}

	// Perform assertions on the response
	if assetResponse.Data.ID != "asset1" {
		t.Errorf("Expected asset ID 'asset1', got %s", assetResponse.Data.ID)
	}
	// Add more assertions as needed
}
