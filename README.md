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
    export PLATFORM="linux/amd64"
    export VERSION="0.0.1"
    ```

3. Generate and push the Docker image (published on Docker Hub).

    ```console
    make docker-build docker-push
    ```

## How to deploy

TODO

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
