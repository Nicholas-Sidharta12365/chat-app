package main

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go"
)

func main() {

	fs := http.FileServer(http.Dir("../frontend/dist"))
	http.Handle("/", fs)

    app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID: "1467387",
		Key: "8e5ce9a76a48c9529674",
		Secret: "a4aaf6aef13bfcb7c70e",
		Cluster: "ap1",
		Secure: true,
	  }

    app.Post("/api/messages", func(c *fiber.Ctx) error {

		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}
		pusherClient.Trigger("chat", "message", data)

        return c.JSON([]string{})
    })

    app.Listen(":8000")
}