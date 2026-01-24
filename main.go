package main

import (
	"context"
	"fmt"
	"log"

	"golang-playground/internal/grpcserver"
	pb "golang-playground/proto"

	"github.com/gofiber/fiber/v2"
)

const (
	httpPort = ":3000"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int64  `json:"age"`
}

func main() {
	// In-Process gRPC Server
	grpcSrv := grpcserver.New()
	grpcSrv.Start()

	conn, err := grpcSrv.DialContext(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	greeterClient := pb.NewGreeterClient(conn)

	// Fiber HTTP Server
	app := fiber.New()

	app.Post("/hello", func(c *fiber.Ctx) error {
		user := new(User) // Initialize a new User struct

		// Parse the request body into the user struct
		if err := c.BodyParser(user); err != nil {
			// Log the error and return a 400 Bad Request status
			log.Println("Error parsing body:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Cannot parse JSON request",
				"error":   err.Error(),
			})
		}

		request := pb.HelloRequest{
			Name:  user.Name,
			Email: user.Email,
			Age:   user.Age,
		}

		resp, err := greeterClient.SayHello(context.Background(), &request)
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Failed to call gRPC service: %v", err))
		}

		// Return a JSON response with the received data
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"user":    user,
			"message": resp.Message,
		})
	})

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")

		resp, err := greeterClient.SayHello(context.Background(), &pb.HelloRequest{Name: name})
		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("Failed to call gRPC service: %v", err))
		}

		return c.SendString(resp.Message)
	})

	log.Printf("HTTP server listening on %s", httpPort)
	log.Fatal(app.Listen(httpPort))
}
