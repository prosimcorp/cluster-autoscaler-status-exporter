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

package v1alpha1

// ClusterAutoscalerStatus is the struct of the cluster-autoscaler-status configMap
type ClusterAutoscalerStatus struct {
	Time             string      `yaml:"time"`
	AutoscalerStatus string      `yaml:"autoscalerStatus"`
	ClusterWide      ClusterWide `yaml:"clusterWide"`
	NodeGroups       []NodeGroup `yaml:"nodeGroups"`
}

type ClusterWide struct {
	Health    HealthStatus  `yaml:"health"`
	ScaleUp   ScalingStatus `yaml:"scaleUp"`
	ScaleDown ScalingStatus `yaml:"scaleDown"`
}

type NodeGroup struct {
	Name      string           `yaml:"name"`
	Health    NodeHealthStatus `yaml:"health"`
	ScaleUp   ScalingStatus    `yaml:"scaleUp"`
	ScaleDown ScalingStatus    `yaml:"scaleDown"`
}

type HealthStatus struct {
	Status             string     `yaml:"status"`
	NodeCounts         NodeCounts `yaml:"nodeCounts"`
	LastProbeTime      string     `yaml:"lastProbeTime"`
	LastTransitionTime string     `yaml:"lastTransitionTime"`
}

type NodeHealthStatus struct {
	Status              string     `yaml:"status"`
	NodeCounts          NodeCounts `yaml:"nodeCounts"`
	LastProbeTime       string     `yaml:"lastProbeTime"`
	LastTransitionTime  string     `yaml:"lastTransitionTime"`
	CloudProviderTarget int        `yaml:"cloudProviderTarget"`
	MinSize             int        `yaml:"minSize"`
	MaxSize             int        `yaml:"maxSize"`
}

type NodeCounts struct {
	Registered       RegisteredNodes `yaml:"registered"`
	LongUnregistered int             `yaml:"longUnregistered"`
	Unregistered     int             `yaml:"unregistered"`
}

type RegisteredNodes struct {
	Total        int `yaml:"total"`
	Ready        int `yaml:"ready"`
	NotStarted   int `yaml:"notStarted"`
	BeingDeleted int `yaml:"beingDeleted"`
}

type ScalingStatus struct {
	Status             string `yaml:"status"`
	LastProbeTime      string `yaml:"lastProbeTime"`
	LastTransitionTime string `yaml:"lastTransitionTime"`
}
