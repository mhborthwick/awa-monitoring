package zendesk

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mhborthwick/awa-monitoring-v2/pkg/http/rest"
)

/*
- Load data from https://status.zendesk.com/api/ssp/services.json
- Add data point to InfluxDB
*/

type Response struct {
	Data []Service `json:"data"`
}

type Service struct {
	Attributes struct {
		Name       string `json:"name"`
		Deprecated bool   `json:"deprecated"`
	} `json:"attributes"`
}

type DataPoints struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]interface{}
	Time        time.Time
}

func LoadData() []Service {
	url := "https://status.zendesk.com/api/ssp/services.json"
	body, _ := rest.FetchData(url)
	var r Response
	rest.GetJson(body, &r)
	// fmt.Printf("%+v\n", r)
	return r.Data
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
			Tags:        map[string]string{"provider": "Zendesk"},
			Fields:      map[string]interface{}{i.Attributes.Name: i.Attributes.Deprecated},
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
