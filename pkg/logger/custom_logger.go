package logger

import (
	"fmt"
	"runtime"
	"time"
)

var (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorBlue  = "\033[34m"
	colorGrey  = "\033[90m"
)

const (
	debugLevel = iota + 1
	infoLevel
	errorLevel
)

type CustomLogger struct{}

func NewCustomLogger() Logger {
	return &CustomLogger{}
}

func (cl CustomLogger) printf(level int, format string, a ...any) {
	debugStr := fmt.Sprintf("%s%s%s", colorBlue, "DEBUG", colorRed)
	infoStr := fmt.Sprintf("%s%s%s", colorGreen, "INFO", colorRed)
	errorStr := fmt.Sprintf("%s%s%s", colorRed, "ERROR", colorRed)

	var logLevelStr string
	switch level {
	case debugLevel:
		logLevelStr = debugStr
	case infoLevel:
		logLevelStr = infoStr
	case errorLevel:
		logLevelStr = errorStr
	default:
		logLevelStr = debugStr
	}

	logMsg := fmt.Sprintf(format, a...)
	logStr := fmt.Sprintf("%s%s%s", colorGrey, logMsg, colorGrey)
	timeNow := time.Now().Format("2006-01-02 15:04:05.000")
	calldepth := 2
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
		line = 0
	}
	callerStr := fmt.Sprintf("%s:%d", file, line)

	msg := fmt.Sprintf("%s:%s %s %s", timeNow, callerStr, logLevelStr, logStr)
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
