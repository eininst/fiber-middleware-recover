package main

import (
	"errors"
	recovers "github.com/eininst/fiber-middleware-recover"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(recovers.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		panic(errors.New("my is error"))
		return ctx.Next()
	})
	_ = app.Listen(":8080")
}
