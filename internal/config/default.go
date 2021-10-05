package config

import (
	"github.com/spf13/viper"
	"time"
)

// Default values of configuration options
const (
	// this defines default application name
	defApplicationName = "Artion"

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
	defLoggingFormat = "%{color}%{level:-8s} %{module}/%{shortfile}::%{shortfunc}()%{color:reset}: %{message}"

	// defLachesisUrl holds default Lachesis connection string
	defLachesisUrl = "~/.lachesis/data/lachesis.ipc"

	// defIpfsUrl holds default IPFS connection string
	defIpfsUrl = "localhost:5001"

	// defSkipHttpGateways tells whether to skip known HTTP-to-IPFS gateways
	defSkipHttpGateways = true

	// defMongoUrl holds default MongoDB connection string for local database
	defMongoUrl = "mongodb://localhost:27017"

	// defMongoDatabase holds the default name of the MongoDB local database
	defMongoDatabase = "artion"

	// defSharedMongoUrl holds default MongoDB connection string for shared/replicated database
	defSharedMongoUrl = "mongodb://localhost:27017"

	// defSharedMongoDatabase holds the default name of the shared/replicated MongoDB database
	defSharedMongoDatabase = "artionshared"

	// defCacheEvictionTime holds default time for in-memory eviction periods
	defCacheEvictionTime = 15 * time.Minute

	// defCacheMax size represents the default max size of the cache in MB
	defCacheMaxSize = 2048

	// defApiStateOrigin represents the default origin used for API state syncing
	defApiStateOrigin = "https://localhost"
)

// defCorsAllowOrigins holds CORS default allowed origins.
var defCorsAllowOrigins = []string{"*"}

// applyDefaults sets default values for configuration options.
func applyDefaults(cfg *viper.Viper) {
	// set simple details
	cfg.SetDefault(keyAppName, defApplicationName)
	cfg.SetDefault(keyBindAddress, defServerBind)
	cfg.SetDefault(keyDomainAddress, defServerDomain)
	cfg.SetDefault(keyLoggingLevel, defLoggingLevel)
	cfg.SetDefault(keyLoggingFormat, defLoggingFormat)
	cfg.SetDefault(keyLachesisUrl, defLachesisUrl)
	cfg.SetDefault(keyIpfsUrl, defIpfsUrl)
	cfg.SetDefault(keySkipHttpGateways, defSkipHttpGateways)
	cfg.SetDefault(keyMongoUrl, defMongoUrl)
	cfg.SetDefault(keyMongoDatabase, defMongoDatabase)
	cfg.SetDefault(keySharedMongoUrl, defSharedMongoUrl)
	cfg.SetDefault(keySharedMongoDatabase, defSharedMongoDatabase)
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
