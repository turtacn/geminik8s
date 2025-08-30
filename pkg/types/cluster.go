package types

// ClusterStatus represents the status of the cluster.
type ClusterStatus string

const (
	StatusCreating    ClusterStatus = "Creating"
	StatusRunning     ClusterStatus = "Running"
	StatusDegraded    ClusterStatus = "Degraded"
	StatusFailed      ClusterStatus = "Failed"
	StatusUnknown     ClusterStatus = "Unknown"
	StatusUpgrading   ClusterStatus = "Upgrading"
	StatusReconciling ClusterStatus = "Reconciling"
)

// NodeRole defines the role of a node in the cluster.
type NodeRole string

const (
	RoleLeader   NodeRole = "Leader"
	RoleFollower NodeRole = "Follower"
	RoleUnknown  NodeRole = "Unknown"
)

// ClusterConfig represents the complete configuration for a geminik8s cluster.
type ClusterConfig struct {
	APIVersion string      `yaml:"apiVersion" json:"apiVersion"`
	Kind       string      `yaml:"kind" json:"kind"`
	Metadata   Metadata    `yaml:"metadata" json:"metadata"`
	Spec       ClusterSpec `yaml:"spec" json:"spec"`
}

// Metadata holds metadata about the cluster.
type Metadata struct {
	Name string `yaml:"name" json:"name"`
}

// ClusterSpec defines the desired state of the cluster.
type ClusterSpec struct {
	Network NetworkConfig `yaml:"network" json:"network"`
	Nodes   []NodeInfo    `yaml:"nodes" json:"nodes"`
	Storage StorageConfig `yaml:"storage" json:"storage"`
}

// NetworkConfig holds the network configuration for the cluster.
type NetworkConfig struct {
	VIP string `yaml:"vip" json:"vip"`
}

// NodeInfo contains basic information about a node in the cluster.
type NodeInfo struct {
	IP   string   `yaml:"ip" json:"ip"`
	Role NodeRole `yaml:"role" json:"role"`
}

// StorageConfig holds the storage configuration for the cluster.
type StorageConfig struct {
	// For now, this is a placeholder. We can add PostgreSQL/Kine specific configs here.
	Type string `yaml:"type" json:"type"` // e.g., "postgresql"
}

//Personal.AI order the ending
