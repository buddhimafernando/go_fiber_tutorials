package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct{
	http.Handler

	Address 			string
	DocumentRoot 		string
	DirectoryListing 	bool
	DirectoryIndex 		[]string
	AliasMap 			map[string]string
	
}

// creating a nes file server structure
func FileServer() *Server {
	server:= &Server{
		Address: ":3000",
		DocumentRoot: "./public", // current directory
		DirectoryListing: false,
		DirectoryIndex: []string{"index.html"},
		AliasMap: map[string]string{},
	}
	return server
}

func main() {
	fmt.Println("Hello, World!")

	app := fiber.New()

	// use logger middleware
	app.Use(logger.New())

	// serve static files from the public directory
	app.Static("/files","./public")

	// route to download specific files
	app.Get("/download/:file", func(c *fiber.Ctx) error {
		fileName := c.Params("file")
		filePath := "./public/" + fileName
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).SendString("File not found")
		}
		return c.Download(filePath)
	})


	log.Fatal(app.Listen(":3000"))
}