package log

import (
	"fmt"
	"github.com/showurl/openim-api-sdk/v20220317/conf"
	"log"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
	Error Level = "error"
	None  Level = "none"
)

var level Level = Level(conf.LogLevel)

func Errorf(format string, args ...interface{}) {
	if level != None {
		log.Printf(`{"level": "error", "msg": %v}`, fmt.Sprintf(format, args...))
	}
}
func Infof(format string, args ...interface{}) {
	if level == Error || level == Info {
		log.Printf(`{"level": "info", "msg": %v}`, fmt.Sprintf(format, args...))
	}
}
func Debugf(format string, args ...interface{}) {
	if level == Debug {
		log.Printf(`{"level": "debug", "msg": %v}`, fmt.Sprintf(format, args...))
	}
}
