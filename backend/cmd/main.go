package main

import (
	"context"
	"log"
	"photo-upload-service/cmd/application"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("panic: %v", r)
		}
	}()

	app := application.New()

	if err := app.Run(ctx); err != nil {
		log.Fatalf("application run: %s", err)
	}
}
