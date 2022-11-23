package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var (
	colorReset  = "\033[0m"
	boldText    = "\033[1m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorWhite  = "\033[37m"
)

var debugStr, infoStr, errorStr, fatalStr string

const (
	debugLevel = iota + 1
	infoLevel
	errorLevel
	fatalLevel
)

type CustomLogger struct{}

func New() Logger {
	return CustomLogger{}
}

func init() {
	debugStr = fmt.Sprintf("%s%s%s%s", colorBlue, boldText, "DEBUG", colorReset)
	infoStr = fmt.Sprintf("%s%s%s%s", boldText, colorWhite, "INFO", colorReset)
	errorStr = fmt.Sprintf("%s%s%s%s", colorRed, boldText, "ERROR", colorReset)
	fatalStr = fmt.Sprintf("%s%s%s%s", colorYellow, boldText, "FATAL", colorReset)
}

func (cl CustomLogger) printf(level int, format string, a ...any) {
	var currentColor, logLevelStr string

	switch level {
	case debugLevel:
		logLevelStr = debugStr
		currentColor = colorBlue
	case infoLevel:
		logLevelStr = infoStr
		currentColor = colorWhite
	case errorLevel:
		logLevelStr = errorStr
		currentColor = colorRed
	case fatalLevel:
		logLevelStr = fatalStr
		currentColor = colorYellow
	default:
		logLevelStr = debugStr
		currentColor = colorBlue
	}

	logMsg := fmt.Sprintf(format, a...)
	logStr := fmt.Sprintf("%s%s%s%s", currentColor, logMsg, currentColor, colorReset)

	timeNow := time.Now().Format("2006-01-02 15:04:05.000")
	calldepth := 2
	_, file, line, ok := runtime.Caller(calldepth)

	if !ok {
		file = "???"
		line = 0
	}

	callerStr := fmt.Sprintf("%s:%d", file, line)
	timeAndCallerStr := fmt.Sprintf("%s%s:%s%s", colorGreen, timeNow, callerStr, colorReset)

	msg := fmt.Sprintf("%s | %s - %s", timeAndCallerStr, logLevelStr, logStr)
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
