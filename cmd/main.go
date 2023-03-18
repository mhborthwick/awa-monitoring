package main

import (
	"fmt"

	"github.com/mhborthwick/awa-monitoring-v2/pkg/http/rest"
)

type ZendeskIncidents struct {
	// TODO: update struct
	Data []interface{} `json:"data"`
}

func main() {
	body, err := rest.FetchData("https://catfact.ninja/fact")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var z ZendeskIncidents
	jsonErr := rest.GetJson(body, &z)
	if jsonErr != nil {
		fmt.Println("Error:", jsonErr)
		return
	}
	fmt.Println(z)
}
