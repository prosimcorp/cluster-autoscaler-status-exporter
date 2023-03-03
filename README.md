# CASE (Cluster Autoscaler Status Exporter)

## Description:

An exporter that parses Cluster Autoscaler's status and expose it for Prometheus

## How to release

Each release of this container is done following several steps carefully in order not to break the things for anyone.

1. Test the changes on the code:

    ```console
    make test
    ```

   > A release is not done if this stage fails

2. Define the package information

    ```console
    export VERSION="0.0.1"
    ```

3. Generate and push the Docker image (published on Docker Hub).

    ```console
    make docker-buildx
    ```

## How to deploy

TBD

## Flags

There are several flags that can be configured to change the behaviour of the
application. They are described in the following table:

| Name                | Description                                                       |           Default           | Example                         |
|:--------------------|:------------------------------------------------------------------|:---------------------------:|:--------------------------------|
| `--connection-mode` | Connect from inside or outside Kubernetes                         |          `kubectl`          | `--connection-mode incluster`   |
| `--kubeconfig`      | Path to the kubeconfig file                                       |      `~/.kube/config`       | `--kubeconfig "~/.kube/config"` |
| `--namespace`       | Namespace where to look for Cluster Autoscaler's status configmap |        `kube-system`        | `--namespace "default"`         |
| `--configmap`       | Name of Cluster Autoscaler's status configmap                     | `cluster-autoscaler-status` | `--configmap "another-cm"`      |
| `--metrics-port`    | Port to launch the metrics webserver                              |           `2112`            | `--metrics-port "6789"`         |
| `--metrics-host`    | Host to launch the metrics webserver                              |          `0.0.0.0`          | `--metrics-host "1.1.1.1"`      |
| `--help`            | Show this help message                                            |              -              | -                               |

## How to collaborate

We are open to external collaborations for this project: improvements, bugfixes, whatever.

For doing it, open an issue to discuss the need of the changes, then:

- Fork the repository
- Make your changes to the code
- Open a PR and wait for review

## License

Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.