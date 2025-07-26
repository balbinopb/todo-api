package main

import (
    "todo-api/routes"
    "github.com/gofiber/fiber/v2"
    "todo-api/database"
)

func main() {
    database.Connect()
    app := fiber.New()
    routes.Setup(app)
    app.Listen(":3000")
}