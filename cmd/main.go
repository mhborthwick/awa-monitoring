package main

/*
TODO
- Add tests to adding package
- Set up grafana
*/

import (
	"os"
	"path/filepath"

	"github.com/mhborthwick/awa-monitoring-v2/pkg/adding/hover"
	"github.com/mhborthwick/awa-monitoring-v2/pkg/adding/klaviyo"
	"github.com/mhborthwick/awa-monitoring-v2/pkg/adding/zendesk"
	"github.com/mhborthwick/awa-monitoring-v2/pkg/setup/db"
	"github.com/mhborthwick/awa-monitoring-v2/pkg/setup/env"
)

func main() {
	dir, _ := os.Getwd()
	pathToEnvFile := filepath.Join(dir, ".env")
	envVars := env.LoadEnv(env.GetEnv, pathToEnvFile)
	client := db.NewInfluxDBClient(envVars.URL, envVars.Token)
	klaviyo.AddDataPoint(client, envVars.Org, envVars.Bucket)
	hover.AddDataPoint(client, envVars.Org, envVars.Bucket)
	zendesk.AddDataPoint(client, envVars.Org, envVars.Bucket)
	defer client.Close()
}
