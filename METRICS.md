# Metrics

Metrics example:

```
# HELP docker_hub_image_last_updated last_updated
# TYPE docker_hub_image_last_updated gauge
docker_hub_image_last_updated{image="prometheus-rancher-exporter",user="infinityworks"} 1472731040.0
# HELP docker_hub_image_pulls_total pulls_total
# TYPE docker_hub_image_pulls_total counter
docker_hub_image_pulls_total{image="prometheus-rancher-exporter",user="infinityworks"} 188672.0
# HELP docker_hub_image_stars stars
# TYPE docker_hub_image_stars gauge
docker_hub_image_stars{image="prometheus-rancher-exporter",user="infinityworks"} 3.0
# HELP docker_hub_image_is_automated is_automated
# TYPE docker_hub_image_is_automated gauge
docker_hub_image_is_automated{image="prometheus-rancher-exporter",user="infinityworks"} 1.0
```
