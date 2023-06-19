package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	Token  string
	URL    string
	Org    string
	Bucket string
}

type EnvGetter func(name string) string

func GetEnv(name string) string {
	return os.Getenv(name)
}

// TODO: Refactor later

func LoadProdEnv(envGetterFn EnvGetter) EnvVars {
	token := envGetterFn("DOCKER_INFLUXDB_INIT_ADMIN_TOKEN")
	port := envGetterFn("DOCKER_INFLUXDB_INIT_PORT")
	org := envGetterFn("DOCKER_INFLUXDB_INIT_ORG")
	bucket := envGetterFn("DOCKER_INFLUXDB_INIT_BUCKET")
	endpoint := envGetterFn("DOCKER_INFLUXDB_ENDPOINT")
	url := "http://" + endpoint + ":" + port
	envVars := EnvVars{}
	envVars.Token = token
	envVars.URL = url
	envVars.Org = org
	envVars.Bucket = bucket
	return envVars
}

func LoadEnv(envGetterFn EnvGetter, pathToEnvFile string) EnvVars {
	err := godotenv.Load(pathToEnvFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := envGetterFn("DOCKER_INFLUXDB_INIT_ADMIN_TOKEN")
	port := envGetterFn("DOCKER_INFLUXDB_INIT_PORT")
	org := envGetterFn("DOCKER_INFLUXDB_INIT_ORG")
	bucket := envGetterFn("DOCKER_INFLUXDB_INIT_BUCKET")
	endpoint := envGetterFn("DOCKER_INFLUXDB_ENDPOINT")
	url := "http://" + endpoint + ":" + port
	envVars := EnvVars{}
	envVars.Token = token
	envVars.URL = url
	envVars.Org = org
	envVars.Bucket = bucket
	return envVars
}
