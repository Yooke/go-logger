package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const logPrefix = ""

var (
	logger   *log.Logger
	logLevel = 5
	logFile *os.File
	levelMap map[string]int = map[string]int{
		"DEBUG": 5,
		"INFO":  4,
		"WARN":  3,
		"ERROR": 2,
		"FATAL": 1,
	}
)

func init() {
	logger = log.New(os.Stdout, logPrefix, log.LstdFlags | log.Lshortfile)
}

func SetLevel(level string) {
	level = strings.ToUpper(level)
	if v, ok := levelMap[level]; ok {
		logLevel = v
	} else {
		panic(fmt.Sprintf("The log level %s is invalid", level))
	}
}

func SetLogFile(file string) {
	var err error
	logFile, err = os.OpenFile(file, os.O_APPEND | os.O_RDWR, 0644)
	if os.IsNotExist(err) {
		logFile, err = os.Create(file)
		if err != nil {
			panic(fmt.Sprintf("Create log file %s error: %s\n", file, err.Error()))
		}
	} else if err != nil {
		panic(fmt.Sprintf("Open log file %s error: %s\n", file, err.Error()))
	}
	logger.SetOutput(logFile)
}

func SetFlag(flag int) {
	logger.SetFlags(flag)
}

func Flush() {
	logFile.Sync()
	logFile.Close()
}

func Debug(v ...interface{}) {
	if logLevel >= levelMap["DEBUG"] {
		logger.Output(2, "DEBUG " + fmt.Sprintln(v...))
	}
}

func Debugf(format string, v ...interface{}) {
	if logLevel >= levelMap["DEBUG"] {
		logger.Output(2, "DEBUG " + fmt.Sprintf(format, v...))
	}
}

func Info(v ...interface{}) {
	if logLevel >= levelMap["INFO"] {
		logger.Output(2, "INFO " + fmt.Sprintln(v...))
	}
}

func Infof(format string, v ...interface{}) {
	if logLevel >= levelMap["INFO"] {
		logger.Output(2, "INFO " + fmt.Sprintf(format, v...))
	}
}

func Warn(v ...interface{}) {
	if logLevel >= levelMap["WARN"] {
		logger.Output(2, "WARN " + fmt.Sprintln(v...))
	}
}

func Warnf(format string, v ...interface{}) {
	if logLevel >= levelMap["WARN"] {
		logger.Output(2, "WARN " + fmt.Sprintf(format, v...))
	}
}

func Error(v ...interface{}) {
	if logLevel >= levelMap["ERROR"] {
		logger.Output(2, "ERROR " + fmt.Sprintln(v...))
	}
}

func Errorf(format string, v ...interface{}) {
	if logLevel >= levelMap["ERROR"] {
		logger.Output(2, "ERROR " + fmt.Sprintf(format, v...))
	}
}

func Fatal(v ...interface{}) {
	defer func() {
		Flush()
		os.Exit(1)
	}()
	logger.Output(2, "FATAL " + fmt.Sprintln(v...))
}

func Fatalf(format string, v ...interface{}) {
	defer func() {
		Flush()
		os.Exit(1)
	}()
	logger.Output(2, "FATAL " + fmt.Sprintf(format, v...))
}
