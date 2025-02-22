package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Cấu hình ghi log vào file
	logFile, err := os.OpenFile("log/main.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic("Failed to open log file")
	}

	// Tạo encoder cho log (định dạng JSON)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	fileWriter := zapcore.AddSync(logFile)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		fileWriter,
		zapcore.InfoLevel,
	)

	logger := zap.New(core)
	defer logger.Sync()

	logger.Info("Application started", zap.String("author", "Huuloc"))
	logger.Warn("This is a warning message")
	logger.Error("An error occurred", zap.Int("error_code", 500))
}
