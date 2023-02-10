package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
)

const (

	// MetricsPrefix
	MetricsPrefix = "cluster_autoscaler_status_"
)

var (
	mReady = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "ready_total",
		Help: "number of nodes ready to schedule pods",
	}, []string{"nodegroup"})

	mUnready = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "unready_total",
		Help: "number of nodes not ready to schedule pods",
	}, []string{"nodegroup"})

	mNotStarted = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "notstarted_total",
		Help: "TODO", // TODO
	}, []string{"nodegroup"})

	mLongNotStarted = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "longnotstarted_total",
		Help: "TODO", // TODO
	}, []string{"nodegroup"})

	mRegistered = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "registered_total",
		Help: "TODO", // TODO
	}, []string{"nodegroup"})

	mUnregistered = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: MetricsPrefix + "unregistered_total",
		Help: "TODO", // TODO
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
)

// TODO
func upgradePrometheusMetrics(nodeGroups *NodeGroups) (err error) {

	for _, nodegroup := range *nodeGroups {

		// Convert all the parsed strings into float64
		// TODO abstract this section to a function, oh dirty Diana
		healthReady, err := strconv.ParseFloat(nodegroup.Health.Ready, 64)
		if err != nil {
			return err
		}

		healthUnready, err := strconv.ParseFloat(nodegroup.Health.Unready, 64)
		if err != nil {
			return err
		}

		healthNotStarted, err := strconv.ParseFloat(nodegroup.Health.NotStarted, 64)
		if err != nil {
			return err
		}

		healthLongNotStarted, err := strconv.ParseFloat(nodegroup.Health.LongNotStarted, 64)
		if err != nil {
			return err
		}

		healthRegistered, err := strconv.ParseFloat(nodegroup.Health.Registered, 64)
		if err != nil {
			return err
		}

		healthLongUnregistered, err := strconv.ParseFloat(nodegroup.Health.LongUnregistered, 64)
		if err != nil {
			return err
		}

		healthCloudProviderTarget, err := strconv.ParseFloat(nodegroup.Health.CloudProviderTarget, 64)
		if err != nil {
			return err
		}

		healthCloudProviderMinSize, err := strconv.ParseFloat(nodegroup.Health.CloudProviderMinSize, 64)
		if err != nil {
			return err
		}

		healthCloudProviderMaxSize, err := strconv.ParseFloat(nodegroup.Health.CloudProviderMaxSize, 64)
		if err != nil {
			return err
		}

		// Update all the metrics for this nodegroup
		mReady.WithLabelValues(nodegroup.Name).Set(healthReady)
		mUnready.WithLabelValues(nodegroup.Name).Set(healthUnready)
		mNotStarted.WithLabelValues(nodegroup.Name).Set(healthNotStarted)
		mLongNotStarted.WithLabelValues(nodegroup.Name).Set(healthLongNotStarted)
		mRegistered.WithLabelValues(nodegroup.Name).Set(healthRegistered)
		mUnregistered.WithLabelValues(nodegroup.Name).Set(healthLongUnregistered)
		mCloudProviderTarget.WithLabelValues(nodegroup.Name).Set(healthCloudProviderTarget)
		mCloudProviderMinSize.WithLabelValues(nodegroup.Name).Set(healthCloudProviderMinSize)
		mCloudProviderMaxSize.WithLabelValues(nodegroup.Name).Set(healthCloudProviderMaxSize)
	}

	return nil
}
