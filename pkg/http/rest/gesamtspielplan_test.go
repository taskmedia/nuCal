package rest

import (
	"net/http"
	"testing"
)

func TestAddRouterGesamtspielplan(t *testing.T) {
	// check Content-Type
	checkEndpointPostStatuscode(t, "/rest/v1/gesamtspielplan", "payload", "no-content-type", http.StatusBadRequest)

	// check invalid payload
	checkEndpointPostStatuscode(t, "/rest/v1/gesamtspielplan", "payload", "application/json", http.StatusBadRequest)
	checkEndpointPostStatuscode(t, "/rest/v1/gesamtspielplan", "[{\"invalid\": \"json structure\"}]", "application/json", http.StatusBadRequest)
}
