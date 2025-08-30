package types

import "time"

// NodeStatusType represents the health status of a node.
type NodeStatusType string

const (
	NodeStatusHealthy   NodeStatusType = "Healthy"
	NodeStatusUnhealthy NodeStatusType = "Unhealthy"
	NodeStatusUnknown   NodeStatusType = "Unknown"
)

// Node represents a single machine in the geminik8s cluster.
type Node struct {
	Config NodeConfig `yaml:"config" json:"config"`
	Status NodeStatus `yaml:"status" json:"status"`
}

// NodeConfig holds the configuration for a single node.
type NodeConfig struct {
	Name string   `yaml:"name" json:"name"`
	IP   string   `yaml:"ip" json:"ip"`
	Role NodeRole `yaml:"role" json:"role"`
}

// NodeStatus represents the observed state of a node.
type NodeStatus struct {
	Status            NodeStatusType      `yaml:"status" json:"status"`
	Message           string              `yaml:"message" json:"message"`
	LastHeartbeatTime time.Time           `yaml:"lastHeartbeatTime" json:"lastHeartbeatTime"`
	Services          []ServiceStatus     `yaml:"services" json:"services"`
	HealthChecks      []HealthCheckResult `yaml:"healthChecks" json:"healthChecks"`
}

// ServiceStatus represents the status of a service running on the node.
type ServiceStatus struct {
	Name   string `yaml:"name" json:"name"` // e.g., "k3s", "postgresql", "kine"
	Active bool   `yaml:"active" json:"active"`
	Error  string `yaml:"error,omitempty" json:"error,omitempty"`
}

// HealthCheckResult holds the outcome of a single health check.
type HealthCheckResult struct {
	CheckName  string    `yaml:"checkName" json:"checkName"`
	Success    bool      `yaml:"success" json:"success"`
	Message    string    `yaml:"message" json:"message"`
	Timestamp  time.Time `yaml:"timestamp" json:"timestamp"`
	DurationMs int64     `yaml:"durationMs" json:"durationMs"`
}

// HostMeta is the metadata stored on each node to describe the cluster topology
// from its own perspective. This is crucial for the failover mechanism.
type HostMeta struct {
	// MyID identifies the current node.
	MyID NodeIdentity `yaml:"myId" json:"myId"`
	// PeerID identifies the other node in the cluster.
	PeerID NodeIdentity `yaml:"peerId" json:"peerId"`
	// VIP is the virtual IP for the cluster.
	VIP string `yaml:"vip" json:"vip"`
	// LastModified is the timestamp of the last modification to this file.
	// Used as a simple fencing mechanism during network partitions.
	LastModified time.Time `yaml:"lastModified" json:"lastModified"`
}

// NodeIdentity holds the identifying information for a node.
type NodeIdentity struct {
	Name string   `yaml:"name" json:"name"`
	IP   string   `yaml:"ip" json:"ip"`
	Role NodeRole `yaml:"role" json:"role"`
}

//Personal.AI order the ending
