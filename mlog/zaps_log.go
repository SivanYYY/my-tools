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

func Init(infoName, errorName string, maxSize, maxBackUp, maxAge int, level zapcore.Level) {

	infoWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   infoName,
		MaxSize:    maxSize, // 文件存储大小
		MaxAge:     maxAge,
		MaxBackups: maxBackUp,
		LocalTime:  true,
		Compress:   true,
	})

	errWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   errorName,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		MaxBackups: maxBackUp,
		LocalTime:  true,
		Compress:   true,
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
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	infoCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		infoWriter,
		zap.InfoLevel,
	)

	errCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		errWriter,
		zap.ErrorLevel,
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
