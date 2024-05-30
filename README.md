Simple RESTful API: Build a basic RESTful API with CRUD (Create, Read, Update, Delete) operations on resources. Use Go Fiber for routing and handling HTTP requests/responses.
- Define your data model (resources) and CRUD operations (Create, Read, Update, Delete).

Use Go Fiber for routing: 
- Define routes for each CRUD operation (e.g., app.POST("/users", createUser))
- Implement handler functions for each route to handle requests and responses (e.g., func createUser(c *fiber.Ctx) error { ... })

Implement logic for CRUD operations in your handler functions:
- Use libraries for database interaction (if needed)
- Encode/decode data using JSON format
