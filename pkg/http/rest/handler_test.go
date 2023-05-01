package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockData struct {
	Data string `json:"data"`
}

func TestGetJson(t *testing.T) {
	var actual MockData
	body := []byte(`{"data": "Mock data response"}`)
	err := GetJSON(body, &actual)
	if err != nil {
		t.Fatal(err)
	}
	expected := MockData{Data: "Mock data response"}
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestFetchData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data": "Mock data response"}`))
	}))
	defer server.Close()
	data, err := FetchData(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"data": "Mock data response"}`
	if string(data) != expected {
		t.Errorf("Unexpected response body: got %s want %s", data, expected)
	}
}
