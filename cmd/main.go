package main

import (
	"context"

	"github.com/dyaksa/dating-app/cmd/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	ctx := context.Background()
	api := api.NewApi()

	api.Start(ctx)
}
