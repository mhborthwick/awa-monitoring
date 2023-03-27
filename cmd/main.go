package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mhborthwick/awa-monitoring-v2/pkg/db"
)

// TODO: Move to rest package
type ZendeskIncidents struct {
	// TODO: update struct
	Data []interface{} `json:"data"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("DOCKER_INFLUXDB_INIT_ADMIN_TOKEN")
	port := os.Getenv("DOCKER_INFLUXDB_INIT_PORT")
	url := "http://localhost:" + port
	client := db.NewInfluxDBClient(url, token)
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
