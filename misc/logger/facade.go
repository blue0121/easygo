package logger

var (
	log Logger = &consoleLogger{}
)

func SetLogger(logger Logger) {
	log = logger
}

func Debug(tpl string, args ...any) {
	log.Debug(tpl, args...)
}

func Info(tpl string, args ...any) {
	log.Info(tpl, args...)
}

func Warn(tpl string, args ...any) {
	log.Warn(tpl, args...)
}

func Error(tpl string, args ...any) {
	log.Error(tpl, args...)
}

func Fatal(tpl string, args ...any) {
	log.Fatal(tpl, args...)
}

func Panic(tpl string, args ...any) {
	log.Panic(tpl, args...)
}
