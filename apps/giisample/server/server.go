package server

// @see https://github.com/nytimes/gizmo/blob/master/server/server.go
// Server is the basic interface that defines what to expect from any server.
type Server interface {
	// Register(Service) error
	Start() error
	Stop() error
}
