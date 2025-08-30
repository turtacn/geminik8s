package storage

import (
	"context"
	"fmt"
	"time"

	custom_errors "github.com/turtacn/geminik8s/internal/pkg/errors"
)

// ReplicationStatus represents the status of PostgreSQL logical replication.
type ReplicationStatus string

const (
	ReplicationActive   ReplicationStatus = "Active"
	ReplicationInactive ReplicationStatus = "Inactive"
	ReplicationError    ReplicationStatus = "Error"
	ReplicationUnknown  ReplicationStatus = "Unknown"
)

// Storage a- represents the storage configuration for the cluster.
type Storage struct {
	ID          string
	Postgres    *PostgresConfig
	Kine        *KineConfig
	Replication *Replication
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// PostgresConfig holds the configuration for a PostgreSQL instance.
type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

// ConnectionString returns the lib/pq-compatible connection string.
func (c *PostgresConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode)
}

// KineConfig holds the configuration for a Kine instance.
type KineConfig struct {
	Endpoint string // The endpoint Kine should listen on (e.g., "unix://... or "tcp://...")
}

// Replication represents the state of logical replication.
type Replication struct {
	MasterNodeID   string
	ReplicaNodeID  string
	Status         ReplicationStatus
	LastSyncTime   time.Time
	ReplicationLag time.Duration
}

// Repository defines the interface for storage configuration persistence.
type Repository interface {
	Save(ctx context.Context, storage *Storage) error
	FindByID(ctx context.Context, id string) (*Storage, error)
}

// NewStorage creates a new Storage entity.
func NewStorage(id string, pgConfig *PostgresConfig, kineConfig *KineConfig) (*Storage, error) {
	if pgConfig == nil || kineConfig == nil {
		return nil, custom_errors.New(custom_errors.ValidationError, "Postgres and Kine configs cannot be nil")
	}
	return &Storage{
		ID:       id,
		Postgres: pgConfig,
		Kine:     kineConfig,
		Replication: &Replication{
			Status: ReplicationUnknown,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// UpdateReplicationStatus updates the replication status and lag.
func (s *Storage) UpdateReplicationStatus(status ReplicationStatus, lag time.Duration) {
	s.Replication.Status = status
	s.Replication.ReplicationLag = lag
	s.Replication.LastSyncTime = time.Now().Add(-lag) // Approximate
	s.UpdatedAt = time.Now()
}

// IsReplicationHealthy checks if the replication is active and lag is within a tolerance.
func (s *Storage) IsReplicationHealthy(tolerance time.Duration) bool {
	return s.Replication.Status == ReplicationActive && s.Replication.ReplicationLag <= tolerance
}

//Personal.AI order the ending
