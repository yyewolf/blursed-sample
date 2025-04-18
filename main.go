package main

import "github.com/go-fuego/fuego"

func main() {
	s := fuego.NewServer()
	s.Addr = "0.0.0.0:9999"

	fuego.Get(s, "/", func(c fuego.ContextNoBody) (string, error) {
		return "Hello, World!", nil
	})

	s.Run()
}
