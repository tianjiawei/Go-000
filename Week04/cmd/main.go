package main

import "Go-000/Week04/internal/server"

func main() {
	router := server.NewHttp()
	err := router.Run(":8080")
	if err != nil {
		panic("http shutdown")
	}
}
