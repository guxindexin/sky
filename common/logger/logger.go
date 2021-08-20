package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	log      *zap.SugaredLogger
	levelMap = map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
)

func Setup() {
	var (
		allCore    []zapcore.Core
		basePath   string
		level      zapcore.Level
		fileWriter zapcore.WriteSyncer
	)
	basePath, _ = os.Getwd()
	level = getLoggerLevel(viper.GetString(`log.level`))
	fileWriter = zapcore.AddSync(&lumberjack.Logger{
		Filename: fmt.Sprintf("%s/%s", basePath, viper.GetString(`log.path`)), // 日志文件名
		MaxSize:  viper.GetInt(`log.maxsize`),                                 // 日志文件大小
		//MaxAge:   viper.GetInt(`log.maxage`),  // 最长保存天数
		//MaxBackups:   viper.GetInt(`log.maxbackups`),  // 最多备份几个
		LocalTime: viper.GetBool(`log.localtime`), // 日志时间戳
		Compress:  viper.GetBool(`log.compress`),  // 是否压缩文件，使用gzip
	})

	consoleDebugging := zapcore.Lock(os.Stdout)

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	//consoleEncoder := zapcore.NewJSONEncoder(encoder) // json格式日志
	consoleEncoder := zapcore.NewConsoleEncoder(encoder) // 普通格式日志

	if viper.GetString(`log.level`) == "debug" {
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, consoleDebugging, zap.NewAtomicLevelAt(level)))
	}
	allCore = append(allCore, zapcore.NewCore(consoleEncoder, fileWriter, zap.NewAtomicLevelAt(level)))

	core := zapcore.NewTee(allCore...)

	logger := zap.New(core).WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))
	log = logger.Sugar()
}

func getLoggerLevel(l string) zapcore.Level {
	if level, ok := levelMap[l]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func DPanic(args ...interface{}) {
	log.DPanic(args...)
}

func DPanicf(format string, args ...interface{}) {
	log.DPanicf(format, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
