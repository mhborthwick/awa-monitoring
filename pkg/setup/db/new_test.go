package db

import (
	"net/http"
	"net/http/httptest"
	"testing"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/stretchr/testify/assert"
)

func TestNewInfluxDBClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()
	mockToken := "my-mock-token"
	actualClient := NewInfluxDBClient(server.URL, mockToken)
	defer actualClient.Close()
	expectedClient := influxdb2.NewClient(server.URL, mockToken)
	defer expectedClient.Close()
	assert.IsType(t, expectedClient, actualClient)
}
