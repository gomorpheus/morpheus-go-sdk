// Provides logging functions for use like logDebug("foobar") and logInfo("baz")
// Use `export MORPHEUS_LOG_LEVEL=DEBUG to print more output
// The default log level is INFO
package morpheus

import (
	"fmt"
	"os"
	"strings"
)

var (
	LogLevels = map[string]int {
		"DEBUG": 0,
		"INFO": 1,
		"WARN": 2,
		"ERROR": 3,
		"QUIET": 9,
	}
	DefaultLogLevel = "INFO"
	logLevel string
	logLevelValue int
	// MORPHEUS_LOG_LEVEL = strings.ToUpper(os.Getenv("MORPHEUS_LOG_LEVEL"))

	// DEBUG = (MORPHEUS_LOG_LEVEL == "DEBUG")
	// logger = Logger{Level:"DEBUG"}
)

func init() {
	envLogLevel := strings.ToUpper(os.Getenv("MORPHEUS_LOG_LEVEL"))
	if envLogLevel != "" {
		SetLogLevel(envLogLevel)
	} else {
		SetLogLevel(DefaultLogLevel)
	}
}

type Logger struct {
	Level string
}

func SetLogLevel(newLevel string) int {
	if _, ok := LogLevels[newLevel]; ok {
	    logLevel = newLevel
	} else {
		if logLevel == "" {
			logLevel = DefaultLogLevel
		}
	}
	logLevelValue = LogLevels[logLevel]
	return logLevelValue
}

func GetLogLevel() string {
	return logLevel
}

func log(messages ...interface{}) {
	for _, msg := range messages {
		fmt.Println(msg)
	}
}

func logDebug(messages ...interface{}) {
	if logLevelValue <= LogLevels["DEBUG"] {
		log(messages...)
	}
}

func logInfo(messages ...interface{}) {
	if logLevelValue <= LogLevels["INFO"] {
		log(messages...)
	}
}

func logWarn(messages ...interface{}) {
	if logLevelValue <= LogLevels["WARN"] {
		log(messages...)
	}
}

func logError(messages ...interface{}) {
	if logLevelValue <= LogLevels["ERROR"] {
		log(messages...)
	}
}
