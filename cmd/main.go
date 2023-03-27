package main

import (
	"fmt"

	"github.com/mhborthwick/awa-monitoring-v2/pkg/setup/db"
	"github.com/mhborthwick/awa-monitoring-v2/pkg/setup/env"
)

// TODO: Move to rest package
type ZendeskIncidents struct {
	// TODO: update struct
	Data []interface{} `json:"data"`
}

// func main() {
// 	klaviyo.WriteDataPoint()
// }

func main() {
	envVars := env.LoadEnv(env.GetEnv)
	client := db.NewInfluxDBClient(envVars.URL, envVars.Token)
	writeAPI := client.WriteAPI(envVars.Org, envVars.Bucket)
	fmt.Println(writeAPI)
	defer client.Close()
}

// func main() {
// 	colly.ScrapeData(
// 		"https://status.klaviyo.com/",
// 		"Klaviyo",
// 	)
// 	colly.ScrapeData(
// 		"https://hoverstatus.com/",
// 		"Hover",
// 	)
// }

// func main() {
// 	body, err := rest.FetchData("https://catfact.ninja/fact")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	var z ZendeskIncidents
// 	jsonErr := rest.GetJson(body, &z)
// 	if jsonErr != nil {
// 		fmt.Println("Error:", jsonErr)
// 		return
// 	}
// 	fmt.Println(z)
// }