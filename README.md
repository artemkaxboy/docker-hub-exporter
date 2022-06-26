# LEGACY VERSION

It's a legacy version of code fully compatible with [infinityworks/docker-hub-exporter](https://github.com/infinityworks/docker-hub-exporter). It's not recommended to use it anymore. For new setup use actual version from [artemkaxboy/docker-hub-exporter](https://github.com/artemkaxboy/docker-hub-exporter).

## HOW TO MIGRATE

* Stop and remove old container
* Change `infinityworks/docker-hub-exporter` to `artemkaxboy/docker-hub-exporter:legacy` in your original run command

# Prometheus Docker Hub Exporter

Exposes metrics of container pulls and stars from the Docker Hub API, to a Prometheus compatible endpoint. The exporter is capable of pulling down stats for individual images or for whole namespaces from DockerHub. This is based on the un-documented V2 Docker Hub API.

## Configuration

The image is setup to take parameters from environment variables or flags:

The available environment variables are:

* `IMAGES` Images you wish to monitor: expected format 'user/image1,user/image2'
* `ORGS` Organisations/Users you wish to monitor: expected format 'org1,org2'
* `BIND_PORT` Address on which to expose metrics and web interface. (default ":9170")

Below is a list of the available flags.

* `-images` Images you wish to monitor: expected format 'user/image1,user/image2'
* `-organisations` Organisations/Users you wish to monitor: expected format 'org1,org2'
* `-telemetry-path` Path under which to expose metrics. (default "/metrics")
* `-connection-timeout` Connection timeout in seconds.  (default 5)
* `-connection-retries` Connection retries until failure is raised.  (default 3)
* `-listen-address` Address on which to expose metrics and web interface. (default ":9170")

## Install and deploy

### Run with `docker run` command

Using environment variables:

```shell
docker run -d --restart=always -p 9170:9170 \
  --env IMAGES="infinityworks/docker-hub-exporter,infinityworks/build-tools" \
  --env ORGS="artemkaxboy,nginx" \
  artemkaxboy/docker-hub-exporter:legacy
```

Using flags:

```shell
docker run -d --restart=always -p 9170:9170 \
  artemkaxboy/docker-hub-exporter:legacy \
  -images="infinityworks/docker-hub-exporter,infinityworks/build-tools" \
  -organisations="artemkaxboy,nginx"
```

### Run with `docker compose`

Create `compose.yml`/`docker-compose.yml`:

```yml
services:
  docker-hub-exporter:
    image: artemkaxboy/docker-hub-exporter:legacy
    environment:
      IMAGES: "infinityworks/docker-hub-exporter,infinityworks/build-tools"
      ORGS: "artemkaxboy,nginx"
    ports:
      - "9170:9170"
    restart: unless-stopped
```

Run with docker compose v2:

```shell
docker compose up -d
```

## Known Issues

Currently, there is a known issue with this build where if you provide an image or list of images belonging to a namespace that has also been passed into the application then Prometheus will error during metrics gathering reporting that the metric was already collected with the same name and labels.

## Metrics

Metrics will be made available on port 9170 by default. An example of these metrics can be found in the [METRICS.md](./METRICS.md)

## Copyright

Core module and the original idea: MIT License, [infinityworks/docker-hub-exporter](https://github.com/infinityworks/docker-hub-exporter)
