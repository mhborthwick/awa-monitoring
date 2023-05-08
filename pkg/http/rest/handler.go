package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetJSON(body []byte, p interface{}) error {
	return json.Unmarshal(body, p)
}

func FetchData(url string) ([]byte, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	fmt.Println("Visiting", url)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
