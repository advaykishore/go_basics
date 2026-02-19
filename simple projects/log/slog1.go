package main

import "log/slog"

func main() {
	slog.Info("User logged in",
		"user", "admin",
		"role", "administrator",
	)
}
