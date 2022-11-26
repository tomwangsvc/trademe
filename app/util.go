package app

import (
	"fmt"
	"log"
)

func ContainString(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

type logger interface {
	Print() string
}

func LogInfo(prefix string, args ...interface{}) {
	for i, arg := range args {
		if logger, ok := arg.(logger); ok {
			args[i] = logger.Print()
		}
	}
	log.Println(fmt.Sprintf(prefix, args...))
}
