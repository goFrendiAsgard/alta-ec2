package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
)

func main() {
	port := os.Getenv("HTTP_PORT")
	e := echo.New()
	e.GET("/", hello)
	e.GET("/:name", helloName)
	fmt.Println(port)
	if err := e.Start(port); err != nil {
		fmt.Println(err)
	}
}

func hello(c echo.Context) error {
	return c.String(200, "Hello world")
}

func helloName(c echo.Context) error {
	name := c.Param("name")
	return c.String(200, fmt.Sprintf("Hello %s", name))
}
