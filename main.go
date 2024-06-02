package main

import (
	"github.com/gofiber/fiber"
)

type BookBillings struct{
	BookName string `json:"bookName"`
	BillingId int `json:"billing_id"`
	Amount float64 `json:"amount`
}

type billinglist []BookBillings

var billings = billinglist{
	{BookName: "Game of Thrones I", BillingId: 1, Amount: 5600.00},
	{BookName: "Goosebumps II", BillingId: 2, Amount: 2000.00},
	{BookName: "Harry Potter III", BillingId: 3, Amount: 3500.00},
	{BookName: "The Alchemist", BillingId: 4, Amount: 1500.00},
} 


var mapBilling map[string]interface{}

func setRoutes(app *fiber.App){
	app.Get("/getTotal/:billing_id",getTotal)
}

func getTotal(c *fiber.Ctx){
	id := c.Params("billing_id")
	billId, ok := mapBilling["billing_id"].(string)

	if ok{} //check

	if id == billId{} 
	
}

func main(){
	app := fiber.New()
	setRoutes(app)

	app.Listen(":3000")
}