package db

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func NewInfluxDBClient(url string, token string) influxdb2.Client {
	client := influxdb2.NewClient(url, token)
	return client
}
