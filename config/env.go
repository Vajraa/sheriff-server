package configuration

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" || goEnv == "development" {
		slog.Info("GO_ENV : Development")
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error Loading Env")
			panic(err)
		}
		slog.Info("Env Loaded successfully")
	}
}
