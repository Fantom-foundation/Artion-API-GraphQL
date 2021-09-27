// Package config handles API server configuration binding and loading.
package config

import (
	"crypto/ecdsa"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Config defines configuration options structure for Fantom API server.
type Config struct {
	// AppName holds the name of the application
	AppName string `mapstructure:"app_name"`

	// MySignature represents a signature of the server on blockchain.
	MySignature ServerSignature `mapstructure:"me"`

	// Server configuration
	Server Server `mapstructure:"server"`

	// Logger configuration
	Log Log `mapstructure:"log"`

	// Lachesis represents the node structure
	Lachesis Lachesis `mapstructure:"node"`

	// IPFS represents the node structure
	Ipfs Ipfs `mapstructure:"ipfs"`

	// Database configuration
	Db Database `mapstructure:"db"`

	// Cache configuration
	Cache Cache `mapstructure:"cache"`

	// Cache configuration
	Compiler Compiler `mapstructure:"compiler"`

	// TokenLogoFilePath contains the path to JSON file with the map
	// of known ERC20 tokens to their logo URLs.
	// The file will be loaded on configuration loading.
	TokenLogoFilePath string `mapstructure:"erc20_tokens_file"`

	// TokenLogo is a list of known ERC20 tokens
	// mapped to URL addresses of their logos.
	TokenLogo map[common.Address]string

	// ReScanBlocks represents the number of blocks to be re-scanned.
	RepoCommand RepoCmd `mapstructure:"cmd"`
}

// RepoCmd represents a repository command configuration.
type RepoCmd struct {
	BlockScanReScan uint64
	RestoreStake    string
}

// Server represents the GraphQL server configuration
type Server struct {
	BindAddress     string   `mapstructure:"bind"`
	DomainAddress   string   `mapstructure:"domain"`
	Origin          string   `mapstructure:"origin"`
	Peers           []string `mapstructure:"peers"`
	CorsOrigin      []string `mapstructure:"cors_origins"`
	ReadTimeout     int64    `mapstructure:"read_timeout"`
	WriteTimeout    int64    `mapstructure:"write_timeout"`
	IdleTimeout     int64    `mapstructure:"idle_timeout"`
	HeaderTimeout   int64    `mapstructure:"header_timeout"`
	ResolverTimeout int64    `mapstructure:"resolver_timeout"`
}

// ServerSignature represents the signature used by this server
// on sending requests to the blockchain, especially signed requests.
type ServerSignature struct {
	Address    common.Address   `mapstructure:"address"`
	PrivateKey ecdsa.PrivateKey `mapstructure:"pkey"`
}

// Log represents the logger configuration
type Log struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// Lachesis represents the Lachesis node access configuration
type Lachesis struct {
	Url string `mapstructure:"url"`
}

// Ipfs represents the IPFS node access configuration
type Ipfs struct {
	Url string `mapstructure:"url"`
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

// Compiler represents the contract compilers configuration.
type Compiler struct {
	CompilerTempPath       string `mapstructure:"temp"`
	DefaultSolCompilerPath string `mapstructure:"sol"`
}
