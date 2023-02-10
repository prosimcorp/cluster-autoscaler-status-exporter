package main

import (
	"context"
	"flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/homedir"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

const (
	// SynchronizationScheduleSeconds The time in seconds between synchronizations
	SynchronizationScheduleSeconds = 2

	// Info messages
	GenerateRestClientMessage = "generating rest client to connect to kubernetes"

	// Error messages
	GenerateRestClientErrorMessage = "error connecting to kubernetes api: %s"
	ConfigmapRetrieveErrorMessage  = "error obtaining cluster-autoscaler status configmap from the cluster"
	ConfigMapParseErrorMessage     = "error parsing status configmap (hint: syntax has changed between cluster-autoscaler versions?)"
	MetricsUpdateErrorMessage      = "imposible to update prometheus metrics"
)

// SynchronizeStatus TODO
// Prometheus ref: https://prometheus.io/docs/guides/go-application/
func SynchronizeStatus(client *kubernetes.Clientset, namespace string, configmap string) (err error) {

	for {
		// Get configmap from the cluster
		configMap, err := client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), configmap, metav1.GetOptions{})
		if err != nil {
			log.Print(ConfigmapRetrieveErrorMessage)
			break
		}

		// Look for all the NG names and health arguments
		nodeGroupsNames := GetNodeGroupsNames(configMap.Data["status"])
		nodeGroupsHealthArgs := GetNodeGroupsHealthArguments(configMap.Data["status"])

		// Retrieve data as a Go struct
		nodeGroups, err := GetNodeGroupsObject(nodeGroupsNames, nodeGroupsHealthArgs)
		if err != nil {
			log.Print(ConfigMapParseErrorMessage)
		}

		// Update Prometheus metrics from NodeGroups type data
		err = upgradePrometheusMetrics(nodeGroups)
		if err != nil {
			log.Print(MetricsUpdateErrorMessage)
		}

		time.Sleep(SynchronizationScheduleSeconds * time.Second)
	}

	return err
}

func main() {
	// Get the values from flags
	connectionMode := flag.String("connection-mode", "kubectl", "(optional) What type of connection to use: incluster, kubectl")
	kubeconfig := flag.String("kubeconfig", filepath.Join(homedir.HomeDir(), ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	namespaceFlag := flag.String("namespace", "kube-system", "Kubernetes Namespace where to synchronize the certificates")
	configmapFlag := flag.String("configmap", "cluster-autoscaler-status", "Name of the cluster-autoscaler's status configmap")
	metricsPortFlag := flag.String("metrics-port", "2112", "Port where metrics web-server will run")
	metricsHostFlag := flag.String("metrics-host", "0.0.0.0", "Host where metrics web-server will run")
	flag.Parse()

	// Generate the Kubernetes client to modify the resources
	log.Printf(GenerateRestClientMessage)
	client, err := GetKubernetesClient(*connectionMode, *kubeconfig)
	if err != nil {
		log.Printf(GenerateRestClientErrorMessage, err)
	}

	// Parse Cluster Autoscaler's status configmap in the background
	go SynchronizeStatus(client, *namespaceFlag, *configmapFlag)

	// Start a webserver for exposing metrics endpoint
	metricsHost := *metricsHostFlag + ":" + *metricsPortFlag
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(metricsHost, nil)
}
