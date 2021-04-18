package logging

import (
	"io"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type LogrusEchoLogger struct {
}

func (l *LogrusEchoLogger) Output() io.Writer {
	return logrus.StandardLogger().Out
}

func (l *LogrusEchoLogger) SetOutput(w io.Writer) {
	return
}

func (l *LogrusEchoLogger) Prefix() string {
	panic("Unimplemented")
}

func (l *LogrusEchoLogger) SetPrefix(p string) {
	panic("Unimplemented")
}

func (l *LogrusEchoLogger) Level() log.Lvl {
	panic("Unimplemented")
}

func (l *LogrusEchoLogger) SetLevel(v log.Lvl) {
	panic("Unimplemented")
}

func (l *LogrusEchoLogger) SetHeader(h string) {
	panic("Unimplemented")
}

func (l *LogrusEchoLogger) Print(i ...interface{}) {
	logrus.Print(i...)
}

func (l *LogrusEchoLogger) Printf(format string, args ...interface{}) {
	logrus.Printf(format, args...)
}

func (l *LogrusEchoLogger) Printj(j log.JSON) {
	panic("Unimplemented")
}

func (l *LogrusEchoLogger) Debug(i ...interface{}) {
	logrus.Debug(i...)
}

func (l *LogrusEchoLogger) Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func (l *LogrusEchoLogger) Debugj(j log.JSON) {
	panic("Unimplemented")
}

func (l *LogrusEchoLogger) Info(i ...interface{}) {
	logrus.Info(i...)
}

func (l *LogrusEchoLogger) Infof(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func (l *LogrusEchoLogger) Infoj(j log.JSON) {
	panic("Unimplemented")
}

func (l *LogrusEchoLogger) Warn(i ...interface{}) {
	logrus.Warn(i...)
}

func (l *LogrusEchoLogger) Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func (l *LogrusEchoLogger) Warnj(j log.JSON) {
	panic("Unimplemented")
}

func (l *LogrusEchoLogger) Error(i ...interface{}) {
	logrus.Error(i...)
}

func (l *LogrusEchoLogger) Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func (l *LogrusEchoLogger) Errorj(j log.JSON) {
	panic("Unimplemented")
}

func (l *LogrusEchoLogger) Fatal(i ...interface{}) {
	logrus.Fatal(i...)
}

func (l *LogrusEchoLogger) Fatalj(j log.JSON) {
	panic("Unimplemented")
}

func (l *LogrusEchoLogger) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func (l *LogrusEchoLogger) Panic(i ...interface{}) {
	logrus.Panic(i...)
}

func (l *LogrusEchoLogger) Panicj(j log.JSON) {
	panic("Unimplemented")
}

func (l *LogrusEchoLogger) Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}
