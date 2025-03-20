# CASE (Cluster Autoscaler Status Exporter)

## Description:

An exporter that parses Cluster Autoscaler's status and expose it for Prometheus

## Deployment

We have designed the deployment of this project to allow remote deployment using Kustomize. This way it is possible
to use it with a GitOps approach, using tools such as ArgoCD or FluxCD. Just make a Kustomization manifest referencing
the tag of the version you want to deploy as follows:

```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
- https://github.com/prosimcorp/cluster-autoscaler-status-exporter//deploy/?ref=main
```

## How to run

You can run the exporter using the following command:

```shell
$ go run cmd/cluster-autoscaler-status-exporter.go --config <config>
```

The `config` parameter is the path to the configuration file. The configuration file is a YAML file with the following
structure:

```yaml
---

# Logging configuration
logging:
  # debug, info, warn, error, dpanic, panic, fatal
  level: info
  # console or json
  format: json

# Server configuration
server:
  # Address to listen on
  listenAddress: ":8080"
  # Path to expose metrics on
  metricsPath: "/metrics"

# Configuration for the status config map
# By default if this config section is not provided, the exporter will look for cluster-autoscaler-status config map in
# kube-system namespace
statusConfigMap:
  namespace: "kube-system"
  name: "cluster-autoscaler-status"
```

## Metrics exported

The metrics exported are the following:

| Metric | Description | Type | Labels |
|--------|-------------|------|--------|
| `cluster_autoscaler_status_cloudprovidermaxsize_total` | Maximum number of nodes allowed in the provider. | Gauge | `nodegroup` |
| `cluster_autoscaler_status_cloudproviderminsize_total` | Minimum number of nodes allowed in the provider. | Gauge | `nodegroup` |
| `cluster_autoscaler_status_cloudprovidertarget_total` | Desired number of nodes in the provider. | Gauge | `nodegroup` |
| `cluster_autoscaler_status_health_status` | Health status of the node group (1 = Healthy, 0 = Unhealthy). | Gauge | `nodegroup`, `status` |
| `cluster_autoscaler_status_long_unregistered_total` | Number of nodes that have been unregistered for a long period. | Gauge | `nodegroup` |
| `cluster_autoscaler_status_notstarted_total` | Number of registered nodes that have not started. | Gauge | `nodegroup` |
| `cluster_autoscaler_status_ready_total` | Number of nodes that are ready to schedule pods. | Gauge | `nodegroup` |
| `cluster_autoscaler_status_registered_total` | Total number of registered nodes. | Gauge | `nodegroup` |
| `cluster_autoscaler_status_scaledown_status` | Scale-down status of the node group. | Gauge | `nodegroup`, `status` |
| `cluster_autoscaler_status_scaleup_status` | Scale-up status of the node group. | Gauge | `nodegroup`, `status` |
| `cluster_autoscaler_status_unregistered_total` | Number of unregistered nodes. | Gauge | `nodegroup` |

## How to collaborate

We are open to external collaborations for this project: improvements, bugfixes, whatever.

For doing it, open an issue to discuss the need of the changes, then:

- Fork the repository
- Make your changes to the code
- Open a PR and wait for review

## License

Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.