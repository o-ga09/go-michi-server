package main

import (
	"context"

	"github.com/o-ga09/go-michi-server/app/internal/presenter"
)

func main() {
	server := presenter.Server{Port: "8081"}
	ctx := context.Background()
	if err := server.Run(ctx); err != nil {
		panic(err)
	}
}
