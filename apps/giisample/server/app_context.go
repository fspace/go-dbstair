package server

import "log"

// AppContext is global namespace for all normally used application components
type AppContext struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}
