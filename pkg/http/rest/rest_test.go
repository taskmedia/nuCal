package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// checkEndpointGetStatuscode tests for a given response code for GET method at defined endpoint (path)
// e.g. status code 404 (Not Found) at endpoint '/notexisting'
func checkEndpointGetStatuscode(t *testing.T, httpEndpoint string, expectedHttpStatuscode int) {
	r := SetupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, httpEndpoint, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, expectedHttpStatuscode, w.Code)
}

// Test version HTTP endpoint
func TestVersion(t *testing.T) {
	// 200 - OK
	checkEndpointGetStatuscode(t, "/version", http.StatusOK)
}
