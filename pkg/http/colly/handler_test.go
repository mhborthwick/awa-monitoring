package colly

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSelectors(t *testing.T) {
	klaviyoSelectors := GetSelectors("Klaviyo")
	assert.Equal(
		t,
		&Selector{
			Container: ".components-container .component-inner-container",
			Name:      ".name",
			Status:    ".component-status",
		},
		klaviyoSelectors,
	)
	hoverSelectors := GetSelectors("Hover")
	assert.Equal(
		t,
		&Selector{
			Container: "#statusio_components .component",
			Name:      ".component_name",
			Status:    ".component-status",
		},
		hoverSelectors,
	)
	notAllowed := GetSelectors("Not Allowed")
	assert.Nil(t, notAllowed)
}

func TestScrapeData(t *testing.T) {
	// Klaviyo
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
<title>Test Page</title>
</head>
<body>
<div class="components-container">
<div class="component-inner-container">
<div class="name">Test Name</div>
<div class="component-status">Test Status</div>
</div>
</div>
</body>
</html>
		`))
	}))
	defer server.Close()
	// TODO - Mock Get Allowed Domains - NEXT!
	// https://stackoverflow.com/questions/19167970/mock-functions-in-go
	// To think about tomorrow -  May need to just remove allowed domains
	actual := ScrapeData(server.URL, "Klaviyo")
	expected := []Item{
		{Provider: "Klaviyo", Name: "Test Name", Status: "Test Status"},
	}
	fmt.Println(actual)
	assert.Equal(t, expected, actual)
}
