package main

// HealthStatus represents the status of a node group
type HealthStatus struct {
	Ready            string `json:"ready"`   // Number of nodes ready to schedule pods
	Unready          string `json:"unready"` // Number of nodes not ready to schedule pods in it
	NotStarted       string `json:"notStarted"`
	LongNotStarted   string `json:"longNotStarted"`
	Registered       string `json:"registered"`
	LongUnregistered string `json:"longUnregistered"`

	CloudProviderTarget  string `json:"cloudProviderTarget"` // Desired number of nodes in the provider
	CloudProviderMinSize string `json:"minSize"`             // Maximum number of nodes in the provider
	CloudProviderMaxSize string `json:"maxSize"`             // Minimum number of nodes in the provider
}

// NodeGroup represents available metrics for one node group
type NodeGroup struct {
	Name   string
	Health HealthStatus
}

// NodeGroups represents a group of node groups
type NodeGroups []NodeGroup
