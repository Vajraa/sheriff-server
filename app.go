package main

import (
	"flag"
	"log"
	"sheriff-server/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

	database.SetupMongoDB()

	app.Use(recover.New())
	
	v1 := app.Group("/api/v1")

	v1.Get("/list", func(c *fiber.Ctx) error {
		return c.JSON("hello")
	})

	log.Fatal(app.Listen(*port))
}