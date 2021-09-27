// Package svc implements monitoring and scanning services of the API server.
package svc

// service represents a Service run by the Manager.
type service interface {
	// init initializes the service
	init()

	// run executes the service
	run()

	// close terminates the service
	close()

	// name provides a name of the service
	name() string
}
