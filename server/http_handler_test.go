package server

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHTTPHandler(t *testing.T) {
	app := fiber.New()

	// Setup route
	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("Hello, " + name + "!")
	})

	// Test cases
	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid name",
			path:           "/hello/World",
			expectedStatus: fiber.StatusOK,
			expectedBody:   "Hello, World!",
		},
		{
			name:           "Another valid name",
			path:           "/hello/Fiber",
			expectedStatus: fiber.StatusOK,
			expectedBody:   "Hello, Fiber!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBody, string(body))
		})
	}
}
