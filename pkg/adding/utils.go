package adding

import (
	"time"

	"github.com/mhborthwick/awa-monitoring-v2/pkg/http/colly"
	"github.com/mhborthwick/awa-monitoring-v2/pkg/http/rest"
)

type DataPoints struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]interface{}
	Time        time.Time
}

func LoadScrapeData(url string, provider string) []colly.Item {
	data := colly.ScrapeData(url, provider)
	return data
}

func LoadJSONData[T any](url string, res *T) *T {
	body, _ := rest.FetchData(url)
	rest.GetJSON(body, &res)
	return res
}
