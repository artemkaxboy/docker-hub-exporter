# Prometheus Docker Hub Exporter

Exposes metrics of container pulls and stars from the Docker Hub API, to a Prometheus compatible endpoint. The exporter is capable of pulling down stats for individual images or for whole namespaces from DockerHub. This is based on the un-documented V2 Docker Hub API.

## Configuration

The image is setup to take parameters from environment variables or flags:

The available environment variables are:

* `IMAGES` Image you wish to monitor, format: `user1/image1, user2/image2`
* `NAMESPACES` Namespaces you wish to monitor, format: `namespace1, namespace2`
* `METRICS_PATH` Path under which to expose metrics. (default `/metrics`)

Below is a list of the available flags. You can also find this list by using the `--help` flag.

* `--image` Image you wish to monitor (can be used multiple times)
* `--namespace` Namespace you wish to monitor (can be used multiple times)
* `--metrics-path` Path under which to expose metrics. (default "/metrics")

## Install and deploy

Run manually from Docker Hub:

```
docker run -d --restart=always -p 9170:9170 artemkaxboy/docker-hub-exporter --image="infinityworks/docker-hub-exporter" --namespace="artemkaxboy"
```

## Known Issues

Currently, there is a known issue with this build where if you provide an image or list of images belonging to a namespace that has also been passed into the application then Prometheus will error during metrics gathering reporting that the metric was already collected with the same name and labels.

## Metrics

Metrics will be made available on port 9170 by default. An example of these metrics can be found in the [METRICS.md](./METRICS.md)

## Copyright

Core module and the original idea: MIT License, [infinityworks/docker-hub-exporter](https://github.com/infinityworks/docker-hub-exporter)
