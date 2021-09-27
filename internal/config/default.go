package config

import (
	"github.com/spf13/viper"
	"time"
)

// Default values of configuration options
const (
	// this defines default application name
	defApplicationName = "Artion GraphQL API Server"

	// defSelfAddress is a default address used as a placeholder
	// for actual API server identification.
	// Please make sure to configure your real key for your API server on the wild.
	defSelfAddress    = "0xE8E2ab527D1fDbCe570B221977BB5c3f12dFa1DA"
	defSelfPrivateKey = "0xaa682338447d15ac4462d938716c120d085a0db81d3945b18017ae0788a121a7"

	// EmptyAddress defines an empty address
	EmptyAddress = "0x0000000000000000000000000000000000000000"

	// defServerBind holds default API server binding address
	defServerBind = "localhost:16761"

	// default set of timeouts for the server
	defReadTimeout     = 2
	defWriteTimeout    = 15
	defIdleTimeout     = 1
	defHeaderTimeout   = 1
	defResolverTimeout = 30

	// defServerDomain holds default API server domain address
	defServerDomain = "localhost:16761"

	// defLoggingLevel holds default Logging level
	// See `godoc.org/github.com/op/go-logging` for the full format specification
	// See `golang.org/pkg/time/` for time format specification
	defLoggingLevel = "INFO"

	// defLoggingFormat holds default format of the Logger output
	defLoggingFormat = "%{color}%{level:-8s} %{shortpkg}/%{shortfunc}%{color:reset}: %{message}"

	// defLachesisUrl holds default Lachesis connection string
	defLachesisUrl = "~/.lachesis/data/lachesis.ipc"

	// defIpfsUrl holds default IPFS connection string
	defIpfsUrl = "localhost:5001"

	// defMongoUrl holds default MongoDB connection string
	defMongoUrl = "mongodb://localhost:27017"

	// defMongoDatabase holds the default name of the API persistent database
	defMongoDatabase = "fantom"

	// defCacheEvictionTime holds default time for in-memory eviction periods
	defCacheEvictionTime = 15 * time.Minute

	// defCacheMax size represents the default max size of the cache in MB
	defCacheMaxSize = 4096

	// defSolCompilerPath represents the default SOL compiler path
	defSolCompilerPath = "/usr/bin/solc"

	// defApiStateOrigin represents the default origin used for API state syncing
	defApiStateOrigin = "https://localhost"

	// defBlockScanRescanDepth represents the amount of blocks re-scanned on server start
	defBlockScanRescanDepth = 200
)

// default list of API peers
var defApiPeers = []string{"https://localhost:16761/api"}

// defCorsAllowOrigins holds CORS default allowed origins.
var defCorsAllowOrigins = []string{"*"}

// applyDefaults sets default values for configuration options.
func applyDefaults(cfg *viper.Viper) {
	// set simple details
	cfg.SetDefault(keyAppName, defApplicationName)
	cfg.SetDefault(keyBindAddress, defServerBind)
	cfg.SetDefault(keyDomainAddress, defServerDomain)
	cfg.SetDefault(keySignatureAddress, defSelfAddress)
	cfg.SetDefault(keySignaturePrivateKey, defSelfPrivateKey)
	cfg.SetDefault(keyLoggingLevel, defLoggingLevel)
	cfg.SetDefault(keyLoggingFormat, defLoggingFormat)
	cfg.SetDefault(keyLachesisUrl, defLachesisUrl)
	cfg.SetDefault(keyIpfsUrl, defIpfsUrl)
	cfg.SetDefault(keyMongoUrl, defMongoUrl)
	cfg.SetDefault(keyMongoDatabase, defMongoDatabase)
	cfg.SetDefault(keyApiPeers, defApiPeers)
	cfg.SetDefault(keyApiStateOrigin, defApiStateOrigin)

	// in-memory cache
	cfg.SetDefault(keyCacheEvictionTime, defCacheEvictionTime)
	cfg.SetDefault(keyCacheMaxSize, defCacheMaxSize)

	// server timeouts
	cfg.SetDefault(keyTimeoutRead, defReadTimeout)
	cfg.SetDefault(keyTimeoutWrite, defWriteTimeout)
	cfg.SetDefault(keyTimeoutHeader, defHeaderTimeout)
	cfg.SetDefault(keyTimeoutIdle, defIdleTimeout)
	cfg.SetDefault(keyTimeoutResolver, defResolverTimeout)

	// cors
	cfg.SetDefault(keyCorsAllowOrigins, defCorsAllowOrigins)
}
