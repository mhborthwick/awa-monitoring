package env

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MockGetEnv(name string) string {
	var env string
	if name == "DOCKER_INFLUXDB_INIT_ADMIN_TOKEN" {
		env = "token"
	}
	if name == "DOCKER_INFLUXDB_INIT_PORT" {
		env = "1234"
	}
	if name == "DOCKER_INFLUXDB_INIT_ORG" {
		env = "org"
	}
	if name == "DOCKER_INFLUXDB_INIT_BUCKET" {
		env = "bucket"
	}
	if name == "DOCKER_INFLUXDB_ENDPOINT" {
		env = "localhost"
	}
	return env
}

func Test_Load(t *testing.T) {
	dir, _ := os.Getwd()
	path := filepath.Join(dir, "..", "..", "..", ".env")
	fmt.Println(path)
	envVars := LoadEnv(MockGetEnv, path)
	expected := EnvVars{
		Token:  "token",
		URL:    "http://localhost:1234",
		Org:    "org",
		Bucket: "bucket",
	}
	assert.Equal(t, expected, envVars)
}
