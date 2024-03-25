package presenter

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-michi/michi"
	"github.com/o-ga09/go-michi-server/app/internal/middleware"
)

type Server struct {
	Port string
}

func (s *Server) Run(ctx context.Context) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", s.Port))
	if err != nil {
		return err
	}

	r := michi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AddID)
	r.Use(middleware.RequestLogger)

	r.Route("/v1", func(sub *michi.Router) {
		sub.Group(func(sub *michi.Router) {
			sub.HandleFunc("GET /health", health)
			sub.HandleFunc("GET /health/deep", DBhealth)
		})
	})

	slog.Info("Server Starting ...")
	go func() {
		err := http.Serve(listen, r)
		if err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("Server Sttoping ...")

	return nil
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Println("healthy !")
}

func DBhealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DB healthy !")
}
