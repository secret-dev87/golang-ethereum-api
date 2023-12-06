package main

import (
	"fmt"
	"log"

	"github.com/secret-dev87/golang-ethereum-api/bootstrap"
	"github.com/secret-dev87/golang-ethereum-api/pkg/env"
)

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(
		fmt.Sprintf("%s:%s",
			env.GetEnv("APP_HOST", "localhost"),
			env.GetEnv("APP_PORT", "3000"),
		)))
}
