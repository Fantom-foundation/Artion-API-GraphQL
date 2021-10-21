// Package config handles API server configuration binding and loading.
package config

import (
	"crypto/ecdsa"
	"time"
)

// Config defines configuration options structure for Fantom API server.
type Config struct {
	// AppName holds the name of the application
	AppName string `mapstructure:"app_name"`

	// Server configuration
	Server Server `mapstructure:"server"`

	// Logger configuration
	Log Log `mapstructure:"log"`

	// Node represents the node structure
	Node Node `mapstructure:"node"`

	// IPFS represents the node structure
	Ipfs Ipfs `mapstructure:"ipfs"`

	// Mongo database configuration
	Db Database `mapstructure:"db"`

	// Shared Mongo database configuration
	SharedDb Database `mapstructure:"shared_db"`

	// Cache configuration
	Cache Cache `mapstructure:"cache"`

	// Auth configuration
	Auth Auth `mapstructure:"auth"`

	// RngOracle provides configuration for random feed oracle handling
	RngOracle RandomFeedOracle `mapstructure:"rng"`
}

// Server represents the GraphQL server configuration
type Server struct {
	BindAddress     string   `mapstructure:"bind"`
	CorsOrigin      []string `mapstructure:"cors"`
	ReadTimeout     int64    `mapstructure:"read_timeout"`
	WriteTimeout    int64    `mapstructure:"write_timeout"`
	IdleTimeout     int64    `mapstructure:"idle_timeout"`
	HeaderTimeout   int64    `mapstructure:"header_timeout"`
	ResolverTimeout int64    `mapstructure:"resolver_timeout"`
}

// Log represents the logger configuration
type Log struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// Node represents the Fantom Opera node access configuration
type Node struct {
	Url string `mapstructure:"url"`
}

// Ipfs represents the IPFS node access configuration
type Ipfs struct {
	// Url of the IPFS node
	Url string `mapstructure:"url"`

	// Skip known HTTP-to-IPFS gateways and use our IPFS node instead
	SkipHttpGateways bool `mapstructure:"skip_http_gateways"`

	// Gateway to process IPFS requests instead of IPFS node (like Pinata)
	Gateway string `mapstructure:"gateway"`

	// GatewayBearer represents API key (JWT) to be used for Gateway auth
	GatewayBearer string `mapstructure:"gateway_bearer"`
}

// Database represents the database access configuration.
type Database struct {
	Url    string `mapstructure:"url"`
	DbName string `mapstructure:"db"`
}

// Cache represents the cache sub-system configuration.
type Cache struct {
	Eviction time.Duration `mapstructure:"eviction"`
	MaxSize  int           `mapstructure:"size"`
}

// Auth represents the authentication configuration.
type Auth struct {
	BearerSecret string `mapstructure:"bearer_secret"`
	NonceSecret  string `mapstructure:"nonce_secret"`
}

// RandomFeedOracle configures random oracle feed.
type RandomFeedOracle struct {
	PrivateKey ecdsa.PrivateKey `mapstructure:"pk"`
	ChainID    string           `mapstructure:"chain"`
}
