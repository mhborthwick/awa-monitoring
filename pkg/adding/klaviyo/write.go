package klaviyo

import (
	"fmt"

	"github.com/mhborthwick/awa-monitoring-v2/pkg/http/colly"
)

func LoadData() []colly.Item {
	data := colly.ScrapeData(
		"https://status.klaviyo.com/",
		"Klaviyo",
	)
	return data
}

func WriteDataPoint() {
	data := LoadData()
	fmt.Println(data)
}
