

package controllers

import (
    "strconv"
    "github.com/gofiber/fiber/v2"
    "todo-api/models"
)

var todos = []models.Todo{}
var idCounter = 1

func GetTodos(c *fiber.Ctx) error {
    return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
    todo := new(models.Todo)
    if err := c.BodyParser(todo); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
    }
    todo.ID = idCounter
    idCounter++
    todos = append(todos, *todo)
    return c.Status(201).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    for i, t := range todos {
        if t.ID == id {
            if err := c.BodyParser(&todos[i]); err != nil {
                return err
            }
            todos[i].ID = id
            return c.JSON(todos[i])
        }
    }
    return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
}

func DeleteTodo(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    for i, t := range todos {
        if t.ID == id {
            todos = append(todos[:i], todos[i+1:]...)
            return c.SendStatus(204)
        }
    }
    return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
}
