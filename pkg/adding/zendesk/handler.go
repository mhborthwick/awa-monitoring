package zendesk

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mhborthwick/awa-monitoring-v2/pkg/adding"
)

type Response struct {
	Data []Service `json:"data"`
}

type Service struct {
	Attributes struct {
		Name       string `json:"name"`
		Deprecated bool   `json:"deprecated"`
	} `json:"attributes"`
}

func AddDataPoint(
	client influxdb2.Client,
	org string,
	bucket string,
) {
	url := "https://status.zendesk.com/api/ssp/services.json"
	var res Response
	adding.LoadJSONData(url, &res)
	var dataPoints []adding.DataPoints
	for _, d := range res.Data {
		status := adding.StatusFormatBoolToInt(d.Attributes.Deprecated, false)
		dataPoints = append(dataPoints, adding.DataPoints{
			Measurement: d.Attributes.Name,
			Tags:        map[string]string{"provider": "Zendesk"},
			Fields:      map[string]interface{}{d.Attributes.Name: status},
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
