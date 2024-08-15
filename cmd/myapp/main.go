package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/atipugin/golang-templ-htmx-app/internal/app"
)

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}
