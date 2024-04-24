package railigentx

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListFleets(t *testing.T) {
	// Create a mock HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the request
		if r.URL.Path != "/fleets" {
			t.Errorf("Expected request path /fleets, got %s", r.URL.Path)
		}
		// Respond with mock data
		_, _ = w.Write([]byte(`{"data":[{"id":"fleet1","assetIds":["asset1","asset2"]}]}`))
	}))
	defer mockServer.Close()

	// Create a new Railigent X client with the mock server's URL
	client := NewClient(mockServer.URL, "username", "password")

	// Call the method being tested
	fleetResponse, err := client.ListFleets()
	if err != nil {
		t.Errorf("Error while listing fleets: %v", err)
	}

	// Perform assertions on the response
	if len(fleetResponse.Data) != 1 {
		t.Errorf("Expected 1 fleet, got %d", len(fleetResponse.Data))
	}
	// Add more assertions as needed
}

// Similar tests can be implemented for other methods such as ListAssets, GetFleet, and GetAsset.
