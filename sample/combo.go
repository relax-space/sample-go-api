package sample

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Father struct {
	Name string
	Sex  int
}

type Son struct {
	Father
	Name string
}

func InitSon() (son Son) {
	son = Son{
		Father: Father{
			Name: "fa",
			Sex:  18,
		},
		Name: "so",
	}
	return
}

var ErrLog = Logger(log.New(os.Stderr, "[mysql] ", log.Ldate|log.Ltime|log.Lshortfile))

// Logger is used to log critical error messages.
type Logger interface {
	Print(v ...interface{})
}

// SetLogger is used to set the logger for critical errors.
// The initial logger is os.Stderr.
func SetLogger(logger Logger) error {
	if logger == nil {
		return errors.New("logger is nil")
	}
	ErrLog = logger
	return nil
}

type MyLogger struct {
}

func (MyLogger) Print(v ...interface{}) {
	fmt.Println("=*=*=*=*=*=*=*=*=*=*=*=*=>", v)
}
