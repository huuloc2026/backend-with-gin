package logger

import (
	"os"

	"github.com/huuloc2026/go-backend/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *Logger {

	//logFile, err := os.OpenFile("log/main.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	logWriter := &lumberjack.Logger{
		Filename:   "storages/log/main.log",
		MaxSize:    config.Max_Size, // megabytes
		MaxBackups: config.Max_Backups,
		MaxAge:     config.Max_Age,  //days
		Compress:   config.Compress, // disabled by default
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	//fileWriter := zapcore.AddSync(logFile)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logWriter)),
		// fileWriter,
		zapcore.InfoLevel,
	)

	logger := zap.New(core)
	defer logger.Sync()

	// logger.Info("Application started", zap.String("author", "Huuloc"))
	// logger.Warn("This is a warning message")
	// logger.Error("An error occurred", zap.Int("error_code", 500))
	return &Logger{logger}

}
