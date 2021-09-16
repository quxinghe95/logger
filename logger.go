package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger = nil
var sugar *zap.SugaredLogger = nil

func init() {
	zapCfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "lvl",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, _ = zapCfg.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	sugar = logger.Sugar()
}

func SetLogFile(filename string) {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    50,
		MaxBackups: 5,
		MaxAge:     7,
		Compress:   false,
	}
	writter := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))

	core := zapcore.NewCore(getEncoder(), writter, zapcore.InfoLevel)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugar = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "ts"
	encoderConfig.LevelKey = "lvl"
	encoderConfig.NameKey = "logger"
	encoderConfig.CallerKey = "caller"
	encoderConfig.FunctionKey = zapcore.OmitKey
	encoderConfig.MessageKey = "msg"
	encoderConfig.StacktraceKey = "stacktrace"
	encoderConfig.LineEnding = zapcore.DefaultLineEnding
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// Debugf ...
func Debugf(template string, args ...interface{}) {
	sugar.Debugf("\x1b[0;34m"+template+"\x1b[0m", args...)
}

// Infof ...
func Infof(template string, args ...interface{}) {
	sugar.Infof("\x1b[0;32m"+template+"\x1b[0m", args...)
}

// Warnf ...
func Warnf(template string, args ...interface{}) {
	sugar.Warnf("\x1b[0;35m"+template+"\x1b[0m", args...)
}

// Errorf ...
func Errorf(template string, args ...interface{}) {
	sugar.Errorf("\x1b[0;31m"+template+"\x1b[0m", args...)
}

// Fatalf ...
func Fatalf(template string, args ...interface{}) {
	sugar.Fatalf(template, args...)
	panic("bad thing happened")
}
