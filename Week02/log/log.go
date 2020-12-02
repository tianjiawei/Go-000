package log

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	LoggerInfo  *zap.Logger
	LoggerDebug *zap.Logger
	LoggerWarn  *zap.Logger
	LoggerError *zap.Logger
}

var Logger Log

func init() {
	//初始化4个等级的日志
	Root, _ := os.Getwd()
	logBasePath := Root + "/storage/log/"
	var err error
	Logger.LoggerDebug, err = InitLog(logBasePath+"debug.log", "debug")
	Logger.LoggerInfo, err = InitLog(logBasePath+"info.log", "info")
	Logger.LoggerWarn, err = InitLog(logBasePath+"warn.log", "warn")
	Logger.LoggerError, err = InitLog(logBasePath+"error.log", "error")

	if err != nil {
		panic("初始化日志组件失败")
	}
}

func InitLog(logPath string, logLevel string) (*zap.Logger, error) {

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    500,     // megabytes
		MaxBackups: 1,       // 最多保留300个备份
		MaxAge:     1,       // days
		Compress:   false,   // 是否压缩 disabled by default
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var level zapcore.Level
	switch logLevel {
	case "info":
		level = zapcore.InfoLevel
	case "debug":
		level = zapcore.DebugLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
		level)
	logger := zap.New(core, zap.AddCaller())
	return logger, nil
}

func (l *Log) Debug(msg string, fields ...zapcore.Field) {
	l.LoggerDebug.Debug(msg, fields...)
	defer l.LoggerInfo.Sync()
}

func (l *Log) Info(msg string, fields ...zapcore.Field) {
	l.LoggerInfo.Info(msg, fields...)
	defer l.LoggerInfo.Sync()
}

func (l *Log) Warn(msg string, fields ...zapcore.Field) {
	l.LoggerWarn.Warn(msg, fields...)
	defer l.LoggerInfo.Sync()
}

func (l *Log) Error(msg string, fields ...zapcore.Field) {
	l.LoggerError.Error(msg, fields...)
	defer l.LoggerInfo.Sync()
}
