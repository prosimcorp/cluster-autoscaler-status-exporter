/*
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
*/

package server

import (
	//
	"net/http"
	"time"

	//
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	//
	"cluster-autoscaler-status-exporter/api/v1alpha1"
	"cluster-autoscaler-status-exporter/internal/prometheus"
)

const (
	// SynchronizationScheduleSeconds The time in seconds between synchronizations
	SynchronizationScheduleSeconds = 2
)

// SynchronizeStatus TODO
// Prometheus ref: https://prometheus.io/docs/guides/go-application/
func SynchronizeStatus(app *v1alpha1.ApplicationSpec) {

	for {

		app.Logger.Debug("Synchronizing status")

		// Get configmap from the cluster
		configMap, err := app.KubeRawClient.CoreV1().ConfigMaps(app.Config.StatusConfigMap.Namespace).Get(app.Context, app.Config.StatusConfigMap.Name, metav1.GetOptions{})
		if err != nil {
			app.Logger.Error("Error obtaining cluster-autoscaler status configmap from the cluster")

			time.Sleep(SynchronizationScheduleSeconds * time.Second)
			continue
		}

		// Parse configMap as YAML
		configMapData := v1alpha1.ClusterAutoscalerStatus{}
		err = yaml.Unmarshal([]byte(configMap.Data["status"]), &configMapData)
		if err != nil {
			app.Logger.Error("Error parsing status configmap (hint: syntax has changed between cluster-autoscaler versions?)")
		}

		// Update Prometheus metrics from NodeGroups type data
		err = prometheus.UpgradePrometheusMetrics(&configMapData)
		if err != nil {
			app.Logger.Error("Imposible to update prometheus metrics")
		}

		time.Sleep(SynchronizationScheduleSeconds * time.Second)
	}
}

func RunMetricsServer(app *v1alpha1.ApplicationSpec) error {

	// Create a new http server
	server := &http.Server{
		Addr:    app.Config.Server.ListenAddress,
		Handler: nil,
	}

	// Start the server
	app.Logger.Info("Starting server", zap.String("address", app.Config.Server.ListenAddress))

	go SynchronizeStatus(app)

	http.Handle(app.Config.Server.MetricsPath, promhttp.Handler())
	err := server.ListenAndServe()
	if err != nil {
		app.Logger.Error("Error starting server", zap.Error(err))
		return err
	}

	return nil
}
