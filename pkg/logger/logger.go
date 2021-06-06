package logger

import (
	"context"
	"log"
)

const (
	INFO_LEVEL  = "info"
	DEBUG_LEVEL = "debug"
	WARN_LEVEL  = "warn"

	GFGLoggerConText = "gfgLogger"
)

type (
	Logger interface {
		WithFields(fields map[string]interface{}) Logger
		WithPrefix(prefix string) Logger

		Debugf(format string, args ...interface{})
		Infof(format string, args ...interface{})
		Printf(format string, args ...interface{})
		Warnf(format string, args ...interface{})
		Errorf(format string, args ...interface{})
		Panicf(format string, args ...interface{})

		Debug(args ...interface{})
		Info(args ...interface{})
		Print(args ...interface{})
		Warn(args ...interface{})
		Error(args ...interface{})
		Panic(args ...interface{})

		Debugln(args ...interface{})
		Infoln(args ...interface{})
		Println(args ...interface{})
		Warnln(args ...interface{})
		Errorln(args ...interface{})
		Panicln(args ...interface{})

		GFGDebugf(ctx context.Context, format string, args ...interface{})
		GFGInfof(ctx context.Context, format string, args ...interface{})
		GFGPrintf(ctx context.Context, format string, args ...interface{})
		GFGWarnf(ctx context.Context, format string, args ...interface{})
		GFGErrorf(ctx context.Context, format string, args ...interface{})
		GFGPanicf(ctx context.Context, format string, args ...interface{})

		GFGDebug(ctx context.Context, args ...interface{})
		GFGInfo(ctx context.Context, args ...interface{})
		GFGPrint(ctx context.Context, args ...interface{})
		GFGWarn(ctx context.Context, args ...interface{})
		GFGError(ctx context.Context, args ...interface{})
		GFGPanic(ctx context.Context, args ...interface{})

		GFGDebugln(ctx context.Context, args ...interface{})
		GFGInfoln(ctx context.Context, args ...interface{})
		GFGPrintln(ctx context.Context, args ...interface{})
		GFGWarnln(ctx context.Context, args ...interface{})
		GFGErrorln(ctx context.Context, args ...interface{})
		GFGPanicln(ctx context.Context, args ...interface{})
	}
)

var std Logger

func init() {
	var err error = nil
	std, err = newLogrusLogger(INFO_LEVEL)
	if err != nil {
		log.Panic(err)
	}
}

func WithFields(fields map[string]interface{}) Logger {
	return std.WithFields(fields)
}

func WithPrefix(prefix string) Logger {
	return std.WithPrefix(prefix)
}

func WithMetricType(metricType string) Logger {
	return std.WithFields(map[string]interface{}{
		"metric_type": metricType,
	})
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}
func Printf(format string, args ...interface{}) {
	std.Printf(format, args...)
}
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}
func Info(args ...interface{}) {
	std.Info(args...)
}
func Print(args ...interface{}) {
	std.Print(args...)
}
func Warn(args ...interface{}) {
	std.Warn(args...)
}
func Error(args ...interface{}) {
	std.Error(args...)
}
func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Debugln(args ...interface{}) {
	std.Debugln(args...)
}
func Infoln(args ...interface{}) {
	std.Infoln(args...)
}
func Println(args ...interface{}) {
	std.Println(args...)
}
func Warnln(args ...interface{}) {
	std.Warnln(args...)
}
func Errorln(args ...interface{}) {
	std.Errorln(args...)
}
func Panicln(args ...interface{}) {
	std.Panicln(args...)
}

func GFGDebugf(ctx context.Context, format string, args ...interface{}) {
	std.GFGDebugf(ctx, format, args...)
}
func GFGInfof(ctx context.Context, format string, args ...interface{}) {
	std.GFGInfof(ctx, format, args...)
}
func GFGPrintf(ctx context.Context, format string, args ...interface{}) {
	std.GFGPrintf(ctx, format, args...)
}
func GFGWarnf(ctx context.Context, format string, args ...interface{}) {
	std.GFGWarnf(ctx, format, args...)
}
func GFGErrorf(ctx context.Context, format string, args ...interface{}) {
	std.GFGErrorf(ctx, format, args...)
}
func GFGPanicf(ctx context.Context, format string, args ...interface{}) {
	std.GFGPanicf(ctx, format, args...)
}

func GFGDebug(ctx context.Context, args ...interface{}) {
	std.GFGDebug(ctx, args...)
}
func GFGInfo(ctx context.Context, args ...interface{}) {
	std.GFGInfo(ctx, args...)
}
func GFGPrint(ctx context.Context, args ...interface{}) {
	std.GFGPrint(ctx, args...)
}
func GFGWarn(ctx context.Context, args ...interface{}) {
	std.GFGWarn(ctx, args...)
}
func GFGError(ctx context.Context, args ...interface{}) {
	std.GFGError(ctx, args...)
}
func GFGPanic(ctx context.Context, args ...interface{}) {
	std.GFGPanic(ctx, args...)
}

func GFGDebugln(ctx context.Context, args ...interface{}) {
	std.GFGDebugln(ctx, args...)
}
func GFGInfoln(ctx context.Context, args ...interface{}) {
	std.GFGInfoln(ctx, args...)
}
func GFGPrintln(ctx context.Context, args ...interface{}) {
	std.GFGPrintln(ctx, args...)
}
func GFGWarnln(ctx context.Context, args ...interface{}) {
	std.GFGWarnln(ctx, args...)
}
func GFGErrorln(ctx context.Context, args ...interface{}) {
	std.GFGErrorln(ctx, args...)
}
func GFGPanicln(ctx context.Context, args ...interface{}) {
	std.GFGPanicln(ctx, args...)
}
