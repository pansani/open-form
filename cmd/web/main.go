package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/occult/pagode/pkg/handlers"
	"github.com/occult/pagode/pkg/log"
	"github.com/occult/pagode/pkg/services"
	"github.com/occult/pagode/pkg/tasks"
)

func main() {
	// Start a new container.
	c := services.NewContainer()
	defer func() {
		// Gracefully shutdown all services.
		fatal("shutdown failed", c.Shutdown())
	}()

	// Build the router.
	if err := handlers.BuildRouter(c); err != nil {
		fatal("failed to build the router", err)
	}

	// Register all task queues.
	// TODO: Tasks disabled for MySQL
	// tasks.Register(c)

	// Start the task runner to execute queued tasks.
	// TODO: Tasks disabled for MySQL
	// c.Tasks.Start(context.Background())

	// Register all job handlers.
	tasks.RegisterJobs(c)

	// Start the server.
	go func() {
		srv := http.Server{
			Addr:         fmt.Sprintf("%s:%d", c.Config.HTTP.Hostname, c.Config.HTTP.Port),
			Handler:      c.Web,
			ReadTimeout:  c.Config.HTTP.ReadTimeout,
			WriteTimeout: c.Config.HTTP.WriteTimeout,
			IdleTimeout:  c.Config.HTTP.IdleTimeout,
		}

		if c.Config.HTTP.TLS.Enabled {
			certs, err := tls.LoadX509KeyPair(c.Config.HTTP.TLS.Certificate, c.Config.HTTP.TLS.Key)
			fatal("cannot load TLS certificate", err)

			srv.TLSConfig = &tls.Config{
				Certificates: []tls.Certificate{certs},
			}
		}

		if err := c.Web.StartServer(&srv); errors.Is(err, http.ErrServerClosed) {
			fatal("shutting down the server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the web server and task runner.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	<-quit
}

// fatal logs an error and terminates the application, if the error is not nil.
func fatal(msg string, err error) {
	if err != nil {
		log.Default().Error(msg, "error", err)
		os.Exit(1)
	}
}
