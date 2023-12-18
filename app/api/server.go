package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Rhiadc/ms-base-go/config"
	"github.com/go-chi/chi"
)

type Server struct {
	Config config.Config
}

func NewServer(opts ...func(server *Server) error) (*Server, error) {
	server := &Server{}
	for _, opt := range opts {
		if err := opt(server); err != nil {
			return nil, err
		}
	}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	r := chi.NewRouter()
	server.router(r)
	s := &http.Server{Addr: fmt.Sprintf("0.0.0.0:%s", server.Config.APIPort), Handler: r}

	go func() {
		<-sig
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Println("Graceful shutdown timed out. Forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := s.Shutdown(shutdownCtx)
		if err != nil {
			log.Printf("Error during graceful shutdown: %v", err)
		}
		serverStopCtx()
	}()

	go func() {
		log.Printf("Server started on %s, pid: %d", s.Addr, os.Getpid())
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error starting server: %v", err)
		}
	}()

	// Wait for server to finish
	<-serverCtx.Done()

	log.Println("Server stopped gracefully.")
	return server, nil
}

func WithConfig(appConfig config.Config) func(server *Server) error {
	return func(server *Server) error {
		server.Config = appConfig
		return nil
	}
}

func WithService() func(server *Server) error {
	return func(server *Server) error {
		return nil
	}
}
