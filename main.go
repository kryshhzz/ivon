package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/django"
    "os"
)

func main() {

    engine := django.New("./views", ".html")

    app := fiber.New(fiber.Config{
        Views: engine,
    })
    app.Get("/", func(c *fiber.Ctx) error {

        return c.Render("index", fiber.Map{
            "student_name": "Steve",
        })
    })

    log.Fatal(app.Listen(":"+os.Getenv("PORT")))
}
