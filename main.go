package main

import (
	"log"
	"net/http"
	"strconv"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// User data structure
type UserData struct{
	Id 		string `json:"id"`
	Name 	string `json:"name"`
	Email	string `json:"email"`
	Birthyear string `json:"birthyear"`
}

// Dummy data of userdata
var userInfo = []UserData{
	{Id: "1", Name:"Buddhima Fernando", Email:"buddhima@gmail.com", Birthyear:"1997"},
	{Id: "2", Name: "Dinesh Perera", Email:"dinesh@gmail.com" , Birthyear:"1998"},
	{Id: "3", Name:"Sara De Alwis", Email: "sara@gmail.com"	, Birthyear:"1999"},
	{Id: "4", Name:"Bhagya Semage", Email: "bhagya@gmail.com", Birthyear:"1996"},
}

func handleWebSocket(c *websocket.Conn) {
    for {
        // Read message from client
        mt, msg, err := c.ReadMessage()
        if err != nil {
            log.Println("read:", err)
            break
        }
        log.Printf("recv: %s", msg)

        // Write message back to client
        err = c.WriteMessage(mt, msg)
        if err != nil {
            log.Println("write:", err)
            break
        }
    }
}


func setupRoutes(app *fiber.App){
	app.Get("/users",GetAllUsers)
	app.Get("/getUser/:id",GetUser)
	app.Post("/addUser",AddUser)
	app.Put("/updateUser/:id",UpdateUser)
	app.Delete("/deleteUser/:id",DeleteUser)
	app.Get("/getAge/:id",GetAge)

	app.Get("/ws", websocket.New(handleWebSocket))
}

// GetAge function to calculate the age of the user
func GetAge(c *fiber.Ctx) error{
	id:=c.Params("id")

	for _, user := range userInfo {
		if user.Id == id{
			age, err := strconv.ParseInt(user.Birthyear,10,64)
			if err != nil{
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			}
			trueAge := 2024 - age
			return c.Status(http.StatusOK).JSON(fiber.Map{"User": user,"age": trueAge})
			
		}
	}
	return c.Status(http.StatusNotFound).JSON(fiber.Map{"error":"User not found"})
	
}

// GetAllUser function to retrieve all user details
func GetAllUsers(c *fiber.Ctx) error{
	return c.Status(http.StatusOK).JSON(fiber.Map{"All Users": userInfo})
}

// GetUser function to retrieve the user details
func GetUser(c *fiber.Ctx) error{
	id:=c.Params("id")

	for _, user := range userInfo {
		if user.Id == id{
			return c.Status(http.StatusOK).JSON(fiber.Map{"User": user})
		}
	}
	return c.Status(http.StatusNotFound).JSON(fiber.Map{"error":"User not found"})

}

// AddUser function to add a new user
func AddUser(c *fiber.Ctx) error{
	var newUser UserData

	if err := c.BodyParser(&newUser); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		return err
	}

	userInfo = append(userInfo, newUser)
	return c.SendString("New user added successfully")

}

// UpdateUser function to update the user details
func UpdateUser(c *fiber.Ctx) error{
	id:=c.Params("id")
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

}

// DeleteUser function to delete the user details
func DeleteUser(c *fiber.Ctx) error{
	id := c.Params("id")

	for i, lang := range userInfo {
		if lang.Id == id {
			userInfo = append(userInfo[:i], userInfo[i+1:]...)
			break
		}
	}
	return c.Status(http.StatusNoContent).JSON(fiber.Map{"data": userInfo})


}

func main(){
	// creates a new instance of fiber
	app := fiber.New()

	// setup routes
	setupRoutes(app)

	app.Listen(":3000")
}
