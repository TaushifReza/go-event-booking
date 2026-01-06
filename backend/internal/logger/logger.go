package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

func logWriter(filename string) zapcore.WriteSyncer {
    return zapcore.AddSync(&lumberjack.Logger{
        Filename:   "logs/" + filename,
        MaxSize:    10,
        MaxBackups: 5,
        MaxAge:     30,
        Compress:   true,
    })
}

func InitLogger() {
    encoderCfg := zap.NewProductionEncoderConfig()
    encoderCfg.TimeKey = "timestamp"
    encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

    encoder := zapcore.NewJSONEncoder(encoderCfg)

    // INFO LOG
    infoCore := zapcore.NewCore(
        encoder,
        logWriter("info.log"),
        zap.LevelEnablerFunc(func(level zapcore.Level) bool {
            return level == zapcore.InfoLevel
        }),
    )

    // WARN LOG
    warnCore := zapcore.NewCore(
        encoder,
        logWriter("warn.log"),
        zap.LevelEnablerFunc(func(level zapcore.Level) bool {
            return level == zapcore.WarnLevel
        }),
    )

    // ERROR LOG
    errorCore := zapcore.NewCore(
        encoder,
        logWriter("error.log"),
        zap.LevelEnablerFunc(func(level zapcore.Level) bool {
            return level >= zapcore.ErrorLevel
        }),
    )

    // Combine all cores
    core := zapcore.NewTee(infoCore, warnCore, errorCore)

    Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
