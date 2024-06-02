package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type BookBillings struct {
	BookName  string  `json:"bookName"`
	BillingId int     `json:"billing_id"`
	Amount    float64 `json:"amount"`
}

type billinglist []BookBillings

var billings = billinglist{
	{BookName: "Game of Thrones I", BillingId: 1, Amount: 5600.00},
	{BookName: "Goosebumps II", BillingId: 2, Amount: 2000.00},
	{BookName: "Harry Potter III", BillingId: 3, Amount: 3500.00},
	{BookName: "The Alchemist", BillingId: 4, Amount: 1500.00},
}

var mapBilling map[int]BookBillings

func init() {
	// Initialize the map
	mapBilling := make(map[int]BookBillings)
	for _, billing := range billings {
		mapBilling[billing.BillingId] = billing
	}
}

func setRoutes(app *fiber.App) {
	app.Get("/getTotal/:billing_id", getTotal)
}

func getTotal(c *fiber.Ctx) error {
	// Get the billing_id parameter from the request
	id := c.Params("billing_id")
	billId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid billing_id"})
	}

	// Retrieve the billing details from the map
	billing, ok := mapBilling[billId]
	if !ok {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Billing not found"})
	}

	// Return the billing details as JSON
	return c.JSON(billing)
}

func main() {
	app := fiber.New()
	setRoutes(app)

	app.Listen(":3000")
}
