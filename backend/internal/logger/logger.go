package logger

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

var ServerLogger = New("URMC-HUB")

type Logger struct {
	info  *log.Logger
	debug *log.Logger
	err   *log.Logger
}

func New(appName string) *Logger {

	logDir := filepath.Join(os.Getenv("LOCALAPPDATA"), appName)
	os.MkdirAll(logDir, 0755)

	logPath := filepath.Join(logDir, "app.log")

	rotator := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    20,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}

	return &Logger{
		info:  log.New(rotator, "INFO  ", log.LstdFlags|log.Lshortfile),
		err:   log.New(rotator, "ERROR ", log.LstdFlags|log.Lshortfile),
		debug: log.New(rotator, "DEBUG ", log.LstdFlags|log.Lshortfile),
	}

}

func (l *Logger) Info(v ...any) {
	l.info.Println(v...)
}

func (l *Logger) Error(v ...any) {
	l.err.Println(v...)
}

func (l *Logger) Debug(v ...any) {
	l.debug.Println(v...)
}

func (l *Logger) Infof(format string, v ...any) {
	l.info.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...any) {
	l.err.Printf(format, v...)
}

func (l *Logger) Debugf(format string, v ...any) {
	l.debug.Printf(format, v...)
}
