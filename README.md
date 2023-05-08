# AWA Monitoring

> 🚨 This project is currently in development

## Setup

- Create `.env` file

- Add the following to `.env` file

```sh
DOCKER_INFLUXDB_INIT_MODE=setup
DOCKER_INFLUXDB_INIT_USERNAME=change_me
DOCKER_INFLUXDB_INIT_PASSWORD=change_me # min 8 characters
DOCKER_INFLUXDB_INIT_ADMIN_TOKEN=change_me # openssl rand -hex 32
DOCKER_INFLUXDB_INIT_ORG=change_me # org name
DOCKER_INFLUXDB_INIT_BUCKET=change_me # bucket name
DOCKER_INFLUXDB_INIT_RETENTION=change_me # ex. 4d
DOCKER_INFLUXDB_INIT_PORT=8086
DOCKER_INFLUXDB_INIT_HOST=influxdb
GRAFANA_PORT=3000
```

## Checklist

- [x] Add handlers

  - [x] Zendesk - http

  - [x] Klaviyo - scraper

  - [x] Hover - scraper

- [x] Set up InfluxDB

- [x] Set up Grafana

- [ ] Set up Kubernetes CronJob

- [ ] Deploy
