package core

import (
	"database/sql"
	"log"
)

// AppContext is global namespace for all normally used application components
type AppContext struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger

	DB *sql.DB
}

/**
- dao 有的go 开发者将dao也作为全局共享组件附着在Application|AppContext 上
-


*/
