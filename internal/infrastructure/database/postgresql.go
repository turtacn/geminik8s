package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/turtacn/geminik8s/internal/pkg/errors"
	"github.com/turtacn/geminik8s/pkg/api"
)

// postgresClient implements the api.DBClient interface for PostgreSQL.
type postgresClient struct {
	connPool *pgxpool.Pool
	connStr  string
}

// NewPostgresClient creates a new PostgreSQL client.
// The connection string should be in the format: "postgres://user:password@host:port/database"
func NewPostgresClient(connectionString string) api.DBClient {
	return &postgresClient{
		connStr: connectionString,
	}
}

// Connect establishes a connection pool to the database.
func (c *postgresClient) Connect() error {
	pool, err := pgxpool.Connect(context.Background(), c.connStr)
	if err != nil {
		return errors.Wrap(err, errors.DatabaseError, "failed to connect to postgres")
	}
	c.connPool = pool
	return nil
}

// Close closes the database connection pool.
func (c *postgresClient) Close() error {
	if c.connPool != nil {
		c.connPool.Close()
	}
	return nil
}

// Execute runs a command on the database that does not return rows (e.g., INSERT, UPDATE, DELETE).
func (c *postgresClient) Execute(query string, args ...interface{}) error {
	if c.connPool == nil {
		return errors.New(errors.DatabaseError, "database connection is not initialized")
	}
	_, err := c.connPool.Exec(context.Background(), query, args...)
	if err != nil {
		return errors.Wrapf(err, errors.DatabaseError, "failed to execute query")
	}
	return nil
}

// Query runs a command on the database that is expected to return rows (e.g., SELECT).
// This is a simplified implementation. A real one would handle row scanning.
func (c *postgresClient) Query(query string, args ...interface{}) (interface{}, error) {
	if c.connPool == nil {
		return nil, errors.New(errors.DatabaseError, "database connection is not initialized")
	}
	rows, err := c.connPool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, errors.DatabaseError, "failed to execute query")
	}
	defer rows.Close()

	// A real implementation would scan the rows into a struct.
	// For now, we just return the raw rows object as a placeholder.
	return rows, nil
}

//Personal.AI order the ending
