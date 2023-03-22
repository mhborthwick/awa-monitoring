package main

import "github.com/mhborthwick/awa-monitoring-v2/pkg/http/colly"

// TODO: Move to rest package
type ZendeskIncidents struct {
	// TODO: update struct
	Data []interface{} `json:"data"`
}

func main() {
	colly.ScrapeData(
		"https://status.klaviyo.com/",
		"Klaviyo",
	)
	colly.ScrapeData(
		"https://hoverstatus.com/",
		"Hover",
	)
}

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
