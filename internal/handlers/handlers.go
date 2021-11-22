// Package handlers holds HTTP/WS handlers chain along with separate middleware implementations.
package handlers

import "artion-api-graphql/internal/config"

var cfg *config.Config

// SetConfig sets the global configuration.
func SetConfig(c *config.Config) {
	cfg = c
}
