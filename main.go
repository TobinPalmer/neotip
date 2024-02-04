package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/heroku/x/hmetrics/onload"
	"log"
	"os"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	port := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}
