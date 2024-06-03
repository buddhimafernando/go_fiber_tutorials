package main

import (
	"fmt"
	"net/http"
	"os"
	//"github.com/gofiber/fiber/v2"
)

type Server struct{
	http.Handler

	Address 			string
	DocumentRoot 		string
	DirectoryListing 	bool
	DirectoryIndex 		[]string
	AliasMap 			map[string]string
	
}



// func FileServer() *Server {
// }

func main() {
	fmt.Println("Hello, World!")

	dir, _ :=os.Getwd()

	//app := fiber.New()

	http.ListenAndServe(":3000",http.FileServer(http.Dir(dir)))
}