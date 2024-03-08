package mlog

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ZapLogger *zap.SugaredLogger
)

const (
	loggerId = iota + 1
)

type ZapsLog struct {
	LogPath string
	Size    int
	Backup  int
	Age     int
	Level   int
}

func Init(infoName, errorName string, maxSize, maxBackup, maxAge int) {

	infoWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   infoName,
		MaxSize:    maxSize, // 文件存储大小
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		LocalTime:  true,
		Compress:   true,
	})

	errWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   errorName,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackup,
		LocalTime:  true,
		Compress:   true,
	})

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zap.ErrorLevel && lvl >= zap.InfoLevel
	})

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString("[" + caller.TrimmedPath() + "]")
		},
		EncodeName: zapcore.FullNameEncoder,
	}

	infoCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		infoWriter,
		lowPriority,
	)

	errCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		errWriter,
		highPriority,
	)

	logger := zap.New(zapcore.NewTee(infoCore, errCore), zap.AddCaller())
	ZapLogger = logger.Sugar()
	defer logger.Sync()
}

func NewContext(ctx context.Context, fields ...interface{}) context.Context {
	return context.WithValue(ctx, loggerId, WithContext(ctx).With(fields...))
}

func WithContext(ctx context.Context) *zap.SugaredLogger {
	if ctx == nil {
		return ZapLogger
	}
	if ctxLogger, ok := ctx.Value(loggerId).(*zap.SugaredLogger); ok {
		return ctxLogger
	}
	return ZapLogger
}
