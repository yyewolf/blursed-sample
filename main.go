package main

import (
	"strconv"

	"github.com/go-fuego/fuego"
)

func main() {
	s := fuego.NewServer()
	s.Addr = "0.0.0.0:9999"

	fuego.Get(s, "/", func(c fuego.ContextNoBody) (string, error) {
		return "Hello, World!", nil
	})

	fuego.Get(s, "/{code}", func(c fuego.ContextNoBody) (string, error) {
		code := c.PathParam("code")
		statusCode, err := strconv.Atoi(code)
		if err != nil {
			statusCode = 200
		}

		c.SetStatus(statusCode)

		return "Wowie, got: " + code, nil
	})

	s.Run()
}
