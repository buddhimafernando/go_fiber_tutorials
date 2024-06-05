package main

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main(){
	app:=fiber.New()

	//configuring the rate limiting middleware
	app.Use(limiter.New(limiter.Config{
		// maximum number of requests when rate limit exceed
		Max: 10,

		// duration of the rate limit
		Expiration: 1*time.Minute,

		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests",
			})
		},
	}))

	//route to handle the incoming request
	app.Get("/request",func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Request received",
		})
	})

	app.Listen(":3000")

}