
package main

import "github.com/labstack/echo"

var server = ":8081"

func main() {
	s := echo.New()
	s.Static("/", "SPA_Example")
	println("Server en http://localhost" + server)
	s.Start(":8081")
}
