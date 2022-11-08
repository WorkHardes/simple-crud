package logger

type Logger interface {
	Debug(a ...any)
	Debugf(format string, a ...any)
	Info(a ...any)
	Infof(format string, a ...any)
	Error(a ...any)
	Errorf(format string, a ...any)
	Fatal(a ...any)
	Fatalf(format string, a ...any)
}
