package utils

import (
	"errors"
	"runtime"
)

// CurrentFile return current file path of the caller
func CurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("can not get current file info"))
	}
	return file
}
