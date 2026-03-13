package sample

import (
	"log/slog"
)

func testLogs() {
	// ==== slog - правильные сообщения (линейка OK) ====
	slog.Info("starting server")
	slog.Error("failed to connect to database")
	slog.Warn("something went wrong")
	slog.Debug("user authenticated successfully")

	// ==== slog - нарушения ====

	// Правило 1: первая буква должна быть строчной
	slog.Info("Starting server on port 8080") // want "must start with lowercase"

	// Правило 2: только английский язык
	slog.Error("ошибка подключения к базе данных") // want "must be in English only"

	// Правило 3: без спецсимволов и эмодзи
	slog.Warn("warning: something went wrong!!!🚀") // want "must not contain special chars or emojis"

	// Правило 4: без чувствительных данных
	slog.Info("user password: ") // want "must not contain sensitive data"

	// ==== zap - правильные сообщения ==== (В ТЗ речь о unit тестах, а не о костылях)
}
