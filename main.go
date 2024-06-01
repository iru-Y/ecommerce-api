package main

import (
	"fmt"
	"log/slog"

	"main.go/repositories"
	"main.go/routes"
)

func main() {
	db, err := repositories.ConnectDatabase()
	if err != nil {
		slog.Error(fmt.Sprintf("main error: %v", err))
		panic("uau")
	}
	routes.Init("0.0.0.0", "39998", db)
}
