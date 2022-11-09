package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

const (
	debugLevel = iota + 1
	infoLevel
	errorLevel
	fatalLevel
)

type CustomLogger struct{}

func NewCustomLogger() Logger {
	return CustomLogger{}
}

func (cl CustomLogger) printf(level int, format string, a ...any) {
	debugStr := fmt.Sprintf("%s%s%s", colorBlue, "DEBUG", colorBlue)
	infoStr := fmt.Sprintf("%s%s%s", colorGreen, "INFO", colorGreen)
	errorStr := fmt.Sprintf("%s%s%s", colorRed, "ERROR", colorRed)
	fatalStr := fmt.Sprintf("%s%s%s", colorYellow, "FATAL", colorYellow)

	var logLevelStr string

	switch level {
	case debugLevel:
		logLevelStr = debugStr
	case infoLevel:
		logLevelStr = infoStr
	case errorLevel:
		logLevelStr = errorStr
	case fatalLevel:
		logLevelStr = fatalStr
	default:
		logLevelStr = debugStr
	}

	logMsg := fmt.Sprintf(format, a...)
	logStr := logMsg
	timeNow := time.Now().Format("2006-01-02 15:04:05.000")
	calldepth := 2
	_, file, line, ok := runtime.Caller(calldepth)

	if !ok {
		file = "???"
		line = 0
	}

	callerStr := fmt.Sprintf("%s:%d", file, line)

	msg := fmt.Sprintf("%s:%s | %s | %s", timeNow, callerStr, logLevelStr, logStr)
	fmt.Println(msg)
}

func (cl CustomLogger) Debug(a ...any) {
	cl.printf(debugLevel, "%s", a...)
}

func (cl CustomLogger) Debugf(format string, a ...any) {
	cl.printf(debugLevel, format, a...)
}

func (cl CustomLogger) Info(a ...any) {
	cl.printf(infoLevel, "%s", a...)
}

func (cl CustomLogger) Infof(format string, a ...any) {
	cl.printf(infoLevel, format, a...)
}

func (cl CustomLogger) Error(a ...any) {
	cl.printf(errorLevel, "%s", a...)
}

func (cl CustomLogger) Errorf(format string, a ...any) {
	cl.printf(errorLevel, format, a...)
}

func (cl CustomLogger) Fatal(a ...any) {
	cl.printf(fatalLevel, "%s", a...)
	os.Exit(1)
}

func (cl CustomLogger) Fatalf(format string, a ...any) {
	cl.printf(fatalLevel, format, a...)
	os.Exit(1)
}
