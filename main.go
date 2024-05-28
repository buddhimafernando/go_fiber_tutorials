package main

import (
	//"fmt"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// User data structure
type UserData struct{
	Id 		string `json:id`
	Name 	string `json:name`
	Email	string `json:email`
}

// Dummy data of userdata
var userInfo = []UserData{
	{Id: "1", Name:"Buddhima Fernando", Email:"buddhima@gmail.com"},
	{Id: "2", Name: "Dinesh Perera", Email:"dinesh@gmail.com"},
	{Id: "3", Name:"Sara De Alwis", Email: "sara@gmail.com"},
	{Id: "4", Name:"Bhagya Semage", Email: "bhagya@gmail.com"},
}

func main(){
	app := fiber.New()

	fmt.Println(userInfo)

	// GET requests to retrieve user info
	app.Get("/userInfo", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(userInfo)
	})

	// POST requests to add user info
	app.Post("/addNewUser", func(c *fiber.Ctx) error {
		var newUser UserData

		if err := c.BodyParser(&newUser); err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return err
		}

		userInfo = append(userInfo, newUser)
		return c.Status(http.StatusCreated).JSON(newUser)
	})

	// PUT requests to update user info
	app.Put("/updateUser/:id",func(c *fiber.Ctx) error {
		id:= c.Params("id")
		var updatedUser UserData
		updatedUser.Id = id

		if err := c.BodyParser(&updatedUser); err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return err
		}

		for i, lang := range userInfo{
			if lang.Id == updatedUser.Id{
				userInfo = append(userInfo[:i], userInfo[i+1:]...)
				userInfo = append(userInfo, updatedUser)
			}
		}

		return c.Status(http.StatusCreated).JSON(updatedUser)
	})

	// DELETE requests to delete user info

	app.Delete("/deleteUser/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, lang := range userInfo {
			if lang.Id == id {
				userInfo = append(userInfo[:i], userInfo[i+1:]...)
				break
			}
		}
		return c.Status(http.StatusNoContent).JSON(fiber.Map{"data": userInfo})

	})


	app.Listen(":3000")
}
