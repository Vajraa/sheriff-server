package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	configuration "sheriff-server/config"
	"sheriff-server/database"
	"syscall"
	"time"

	"sheriff-server/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	config := slogfiber.Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,
	}

	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	app.Use(recover.New())
	app.Use(slogfiber.NewWithConfig(logger, config))

	router.AuthRoutes(app)

	go func() {
		if err := app.Listen(*port); err != nil {
			slog.Error("Error Occured During Startup", err)
			panic(err)
		}
	}()

	slog.Info("Server up at port:3000")
	configuration.LoadEnv()
	client, dbContext, dbCancel := database.SetupMongoDB()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	slog.Info("Starting Shutdown Procedure")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		slog.Error("Error Occurred During Shutdown", err)
	}

	select {
	case <-ctx.Done():
		database.CloseConnection(client, dbContext, dbCancel)
		slog.Info("Server Shutdown Sequence Complete")
	}
}
