package controllers

import "github.com/gofiber/fiber"

type account struct {
	EmpNumber  string `json:"name"`
	BankName   string `json:"bankName"`
	BankNumber string `json:"bankNumber"`
}

func GetAccounts(c *fiber.Ctx) {
	emp := account{"EmpNumber", "DSB Bank", "ASD122121212"}
	c.JSON(emp)
}
