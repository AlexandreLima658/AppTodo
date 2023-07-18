package controllers

import (
	"github.com/AlexandreLima658/AppTodo/src/database"
	"github.com/AlexandreLima658/AppTodo/src/models"
	"github.com/gofiber/fiber/v2"
	"log"
)

func CreateTodo(ctx *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	newTodo := models.Todo{}
	if err := ctx.BodyParser(&newTodo); err != nil {
		return err
	}
	db.Exec("insert into todos(item) values ($1)", newTodo.Item)
	return ctx.Redirect("/")

}

func SearchTodos(ctx *fiber.Ctx) error {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	todos, err := db.Query("select * from todos")
	if err != nil {
		log.Fatal(err)
	}
	var todoList []models.Todo
	for todos.Next() {
		var todoItem models.Todo
		if err := todos.Scan(&todoItem.ID, &todoItem.Item); err != nil {
			return err
		}
		todoList = append(todoList, todoItem)

	}

	return ctx.Render("index", fiber.Map{
		"Todos": todoList,
	})
}

func UpdatedTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	updatedTodo := models.Todo{}
	if err := ctx.BodyParser(&updatedTodo); err != nil {
		return err
	}
	db.Exec("update todos set item = $2 where id = $1", id, updatedTodo.Item)
	return ctx.JSON(fiber.Map{
		"message": "Updated",
		"todo":    updatedTodo,
	})

}

func DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("delete from todos where id = $1", id)
	return ctx.JSON(fiber.Map{
		"message": "Deleted",
		"id":      id,
	})
}
