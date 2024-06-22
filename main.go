package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"sheriff-server/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {

	flag.Parse()
	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	database.SetupMongoDB()

	config := slogfiber.Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,
	}

	app.Use(recover.New())
	app.Use(slogfiber.NewWithConfig(logger, config))

	v1 := app.Group("/api/v1")

	v1.Get("/list", func(c *fiber.Ctx) error {
		return c.JSON("hello")
	})

	log.Fatal(app.Listen(*port))

}
