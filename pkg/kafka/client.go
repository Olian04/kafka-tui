package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

// Client wraps the kafka-go client and provides a simplified interface
// for interacting with Kafka clusters.
type Client struct {
	brokers []string
	dialer  *kafka.Dialer
	timeout time.Duration
}

// Config holds configuration for creating a Kafka client.
type Config struct {
	Brokers []string
	Timeout time.Duration
}

// NewClient creates a new Kafka client with the given configuration.
func NewClient(config Config) *Client {
	if config.Timeout == 0 {
		config.Timeout = 10 * time.Second
	}

	dialer := &kafka.Dialer{
		Timeout:       config.Timeout,
		DualStack:     true,
		SASLMechanism: nil, // TODO: Add SASL support in future
		TLS:           nil, // TODO: Add TLS support in future
	}

	return &Client{
		brokers: config.Brokers,
		dialer:  dialer,
		timeout: config.Timeout,
	}
}

// Connect establishes a connection to a Kafka broker and returns a Conn.
// This is used for metadata operations.
func (c *Client) Connect(ctx context.Context) (*kafka.Conn, error) {
	if len(c.brokers) == 0 {
		return nil, fmt.Errorf("no brokers configured")
	}

	// Try connecting to the first broker
	conn, err := c.dialer.DialContext(ctx, "tcp", c.brokers[0])
	if err != nil {
		return nil, fmt.Errorf("failed to connect to broker %s: %w", c.brokers[0], err)
	}

	return conn, nil
}

// Brokers returns the list of configured broker addresses.
func (c *Client) Brokers() []string {
	return c.brokers
}

// Close closes any open connections. Currently a placeholder for future use.
func (c *Client) Close() error {
	// TODO: Track and close open connections
	return nil
}
