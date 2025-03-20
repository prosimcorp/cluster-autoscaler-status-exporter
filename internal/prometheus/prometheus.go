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

package prometheus

import (
	//
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	//
	"cluster-autoscaler-status-exporter/api/v1alpha1"
)

const (
	MetricsPrefix = "cluster_autoscaler_status_"
)

// Prometheus metrics
var (
	mReady = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "ready_total",
		Help: "number of nodes ready to schedule pods",
	}, []string{"nodegroup"})

	mNotStarted = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "notstarted_total",
		Help: "number of registered nodes that are not started",
	}, []string{"nodegroup"})

	mRegisteredTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "registered_total",
		Help: "total number of registered nodes",
	}, []string{"nodegroup"})

	mUnregistered = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "unregistered_total",
		Help: "number of unregistered nodes",
	}, []string{"nodegroup"})

	mLongUnregistered = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "long_unregistered_total",
		Help: "number of nodes that have been unregistered for a long time",
	}, []string{"nodegroup"})

	mCloudProviderTarget = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "cloudprovidertarget_total",
		Help: "desired number of nodes in the provider",
	}, []string{"nodegroup"})

	mCloudProviderMinSize = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "cloudproviderminsize_total",
		Help: "minimum number of nodes in the provider",
	}, []string{"nodegroup"})

	mCloudProviderMaxSize = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "cloudprovidermaxsize_total",
		Help: "maximum number of nodes in the provider",
	}, []string{"nodegroup"})

	mHealthStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "health_status",
		Help: "health status of the nodegroup (1=Healthy, 0=Unhealthy)",
	}, []string{"nodegroup", "status"})

	mScaleUpStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "scaleup_status",
		Help: "scale up status of the nodegroup",
	}, []string{"nodegroup", "status"})

	mScaleDownStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "scaledown_status",
		Help: "scale down status of the nodegroup",
	}, []string{"nodegroup", "status"})
)

// UpgradePrometheusMetrics updates Prometheus metrics with the data from the ClusterAutoscalerStatus
func UpgradePrometheusMetrics(configMapData *v1alpha1.ClusterAutoscalerStatus) error {

	for _, nodegroup := range configMapData.NodeGroups {
		regTotal := float64(nodegroup.Health.NodeCounts.Registered.Total)
		regReady := float64(nodegroup.Health.NodeCounts.Registered.Ready)
		regNotStarted := float64(nodegroup.Health.NodeCounts.Registered.NotStarted)
		unregistered := float64(nodegroup.Health.NodeCounts.Unregistered)
		longUnregistered := float64(nodegroup.Health.NodeCounts.LongUnregistered)
		cloudProviderTarget := float64(nodegroup.Health.CloudProviderTarget)
		minSize := float64(nodegroup.Health.MinSize)
		maxSize := float64(nodegroup.Health.MaxSize)

		mRegisteredTotal.WithLabelValues(nodegroup.Name).Set(regTotal)
		mReady.WithLabelValues(nodegroup.Name).Set(regReady)
		mNotStarted.WithLabelValues(nodegroup.Name).Set(regNotStarted)
		mUnregistered.WithLabelValues(nodegroup.Name).Set(unregistered)
		mLongUnregistered.WithLabelValues(nodegroup.Name).Set(longUnregistered)
		mCloudProviderTarget.WithLabelValues(nodegroup.Name).Set(cloudProviderTarget)
		mCloudProviderMinSize.WithLabelValues(nodegroup.Name).Set(minSize)
		mCloudProviderMaxSize.WithLabelValues(nodegroup.Name).Set(maxSize)

		if nodegroup.Health.Status == "Healthy" {
			mHealthStatus.WithLabelValues(nodegroup.Name, nodegroup.Health.Status).Set(1)
		} else {
			mHealthStatus.WithLabelValues(nodegroup.Name, nodegroup.Health.Status).Set(0)
		}

		mScaleUpStatus.WithLabelValues(nodegroup.Name, nodegroup.ScaleUp.Status).Set(1)
		mScaleDownStatus.WithLabelValues(nodegroup.Name, nodegroup.ScaleDown.Status).Set(1)
	}

	return nil
}
