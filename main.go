package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id        int
	Completed bool
	Body      string
}

func main() {
	app := fiber.New()

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	app.Post("api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is reqired"})
		}

		todo.Id = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	app.Put("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.Id) == id {
				todos[i].Completed = !todo.Completed
				return c.Status(201).JSON(fiber.Map{"msg": "Todo succesfully updated"})
			}
		}

		return c.Status(404).JSON(fiber.Map{"msg": "Id not found"})

	})

	app.Delete("api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.Id) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"msg": fmt.Sprintf("Todo %s deleted", id)})
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": fmt.Sprintf("Id %s is not found", id)})
	})

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
