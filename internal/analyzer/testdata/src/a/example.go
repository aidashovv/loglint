package a

import (
	"log/slog"

	"go.uber.org/zap"
)

func testSlog() {
	slog.Info("Starting server")    // want "log should be in lowercase"
	slog.Error("ошибка")            // want "log message should be in english"
	slog.Warn("stooop!!!")          // want "log message should not contain emojis or special symbols"
	slog.Debug("api_key is secret") // want "log message should not contain emojis or special symbols" "log contains sensitive data: api_key"
	slog.Info("server started")     // все четко
}

func testZap() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Starting database")       // want "log should be in lowercase"
	logger.Error("критическая ошибка")     // want "log message should be in english"
	logger.Warn("warning ⚠️")              // want "log message should be in english" "log message should not contain emojis or special symbols"
	logger.Debug("connection established") // все четко
}
