package http_server

import (
	"base-api/app/http/routers"
	infra "base-api/infra/context"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

type httpServer struct {
	infraContext infra.InfraContextInterface
}

type HTTPServer interface {
	RunHTTP(cmd *cobra.Command, args []string) error
}

func New() HTTPServer {
	return &httpServer{}
}

func (h httpServer) initializeHandler() *mux.Router {
	r := mux.NewRouter()
	return routers.InitialRouter(h.infraContext, r)
}

func (h httpServer) SetGracefulTimeout() time.Duration {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*time.Duration(h.infraContext.Config().Server.GraceFulTimeout), "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	return wait
}

func (h httpServer) RunHTTP(cmd *cobra.Command, args []string) error {
	h.infraContext = infra.New()

	// Initialize The Server
	httpServer := &http.Server{
		Handler:      h.initializeHandler(),
		Addr:         h.infraContext.Config().Server.Addr,
		WriteTimeout: time.Duration(h.infraContext.Config().Server.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(h.infraContext.Config().Server.ReadTimeout) * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), h.SetGracefulTimeout())
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	_ = httpServer.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

	return nil
}
