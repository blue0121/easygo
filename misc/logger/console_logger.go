package logger

import (
	"fmt"
	"github.com/petermattis/goid"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type consoleLogger struct {
}

func (c *consoleLogger) Debug(tpl string, args ...any) {
	printf(LOG_DEBUG, tpl, args...)
}

func (c *consoleLogger) Info(tpl string, args ...any) {
	printf(LOG_INFO, tpl, args...)
}

func (c *consoleLogger) Warn(tpl string, args ...any) {
	printf(LOG_WARN, tpl, args...)
}

func (c *consoleLogger) Error(tpl string, args ...any) {
	printf(LOG_ERROR, tpl, args...)
}

func (c *consoleLogger) Fatal(tpl string, args ...any) {
	printf(LOG_FATAL, tpl, args...)
	os.Exit(1)
}

func (c *consoleLogger) Panic(tpl string, args ...any) {
	str := toString(LOG_PANIC, tpl, args...)
	panic(str)
}

func toString(level, tpl string, args ...any) string {
	_, file, line, ok := runtime.Caller(4)

	now := time.Now()
	sb := strings.Builder{}
	sb.WriteString(now.Format(LOG_DATETIME))
	sb.WriteString(fmt.Sprintf(" [%d] ", goid.Get()))
	if ok {
		sb.WriteString(fmt.Sprintf("%s(%d) ", path.Base(file), line))
	}
	sb.WriteString(fmt.Sprintf("[%s] ", level))

	sb.WriteString(fmt.Sprintf(tpl, args...))
	return sb.String()
}

func printf(level, tpl string, args ...any) {
	str := toString(level, tpl, args...)
	fmt.Println(str)
}
