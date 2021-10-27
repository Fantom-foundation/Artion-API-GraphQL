// Package email provides emailing service interface.
package email

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
)

// config represents the configuration setup used by the repository
// to establish and maintain required connectivity to external services
// as needed.
var cfg *config.Config

// log represents the logger to be used by the repository.
var log logger.Logger

// SetConfig sets the configuration to be used to identify email keys and service access data.
func SetConfig(c *config.Config) {
	cfg = c
}

// SetLogger sets the logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l.ModuleLogger("email")
}
