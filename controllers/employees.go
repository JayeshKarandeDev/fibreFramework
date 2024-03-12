package controllers

import "github.com/gofiber/fiber"

type employee struct {
	Name string `json:"name"`
	Age  int16  `json:"age"`
}

func GetUsers(c *fiber.Ctx) {
	emp := employee{"Jayesh", 37}
	c.JSON(emp)
}
