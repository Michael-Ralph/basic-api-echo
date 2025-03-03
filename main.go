package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Displays Index page
func serveIndex(c echo.Context) error {
	return c.File("template/index.html")
}

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static files
	e.Static("/static", "static")

	// Routes
	e.GET("/", serveIndex)
	// Start server
	e.Logger.Fatal(e.Start(":9000"))

}
