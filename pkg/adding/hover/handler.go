package hover

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mhborthwick/awa-monitoring-v2/pkg/adding"
)

func AddDataPoint(
	client influxdb2.Client,
	org string,
	bucket string,
) {
	url := "https://hoverstatus.com/"
	provider := "Hover"
	data := adding.LoadScrapeData(url, provider)
	var dataPoints []adding.DataPoints
	for _, i := range data {
		dataPoints = append(dataPoints, adding.DataPoints{
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
