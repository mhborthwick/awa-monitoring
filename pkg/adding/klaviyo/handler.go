package klaviyo

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mhborthwick/awa-monitoring-v2/pkg/http/colly"
)

type DataPoints struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]interface{}
	Time        time.Time
}

func LoadData() []colly.Item {
	data := colly.ScrapeData(
		"https://status.klaviyo.com/",
		"Klaviyo",
	)
	return data
}

func AddDataPoint(
	client influxdb2.Client,
	org string,
	bucket string,
) {
	data := LoadData()
	var dataPoints []DataPoints
	for _, i := range data {
		dataPoints = append(dataPoints, DataPoints{
			Measurement: "status",
			Tags:        map[string]string{"provider": i.Provider},
			Fields:      map[string]interface{}{i.Name: i.Status},
			Time:        time.Now(),
		})
	}
	writeAPI := client.WriteAPIBlocking(org, bucket)
	for _, dp := range dataPoints {
		p := influxdb2.NewPoint(dp.Measurement, dp.Tags, dp.Fields, dp.Time)
		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			panic(err)
		}
	}
}
