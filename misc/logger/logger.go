package logger

type Logger interface {
	Debug(tpl string, args ...any)
	Info(tpl string, args ...any)
	Warn(tpl string, args ...any)
	Error(tpl string, args ...any)
	Fatal(tpl string, args ...any)
	Panic(tpl string, args ...any)
}
