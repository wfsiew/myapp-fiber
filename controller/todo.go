package controller

import (
    "app/model"
    "app/utils"
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
)

// GetTodos
//
// @Tags Todos
// @Produce json
// @Success 200 {array} int
// @Security BearerAuth
// @Router /app/todo [get]
func GetTodos(c *fiber.Ctx) error {
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    email := claims["email"].(string)
    fmt.Println(email)
    todo := []int{ 1, 2, 3}
    return c.JSON(todo)
}

// CreateTodo
//
// @Tags Todos
// @Param payload body model.Todo true "Todo payload"
// @Accept json
// @Produce json
// @Success 200
// @Security BearerAuth
// @Router /app/todo [post]
func CreateTodo(c *fiber.Ctx) error {
    payload := model.Todo{}
    if err := c.BodyParser(&payload); err != nil {
        utils.Logger.Err(err).Msg(err.Error())
        return err
    }

    errs := utils.ValidatePayload(payload, c)
    if errs != nil {
        return errs
    }

    return c.JSON(fiber.Map{"status": "ok", "data": payload})
}

// GetTodo
//
// @Tags Todos
// @Produce json
// @Param todoId path int true "todoId"
// @Success 200
// @Security BearerAuth
// @Router /app/todo/{todoId} [get]
func GetTodo(c *fiber.Ctx) error {
    return c.JSON((fiber.Map{"data": 1}))
}

// UpdateTodo
//
// @Tags Todos
// @Accept json
// @Produce json
// @Param todoId path int true "todoId"
// @Success 200
// @Security BearerAuth
// @Router /app/todo/{todoId} [put]
func UpdateTodo(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"status": "ok"})
}

// DeleteTodo
//
// @Tags Todos
// @Accept json
// @Produce json
// @Param todoId path int true "todoId"
// @Success 200
// @Security BearerAuth
// @Router /app/todo/{todoId} [delete]
func DeleteTodo(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"status": "ok"})
}

