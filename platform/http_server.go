package platform

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

const (
	shutdownTimeout     = 3 * time.Second
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 10 * time.Second
	defaultAddr         = ":8000"
)

func RunHTTPServer(ctx context.Context, handler http.Handler) {
	server := http.Server{
		Handler:           handler,
		Addr:              defaultAddr,
		ReadTimeout:       defaultReadTimeout,
		ReadHeaderTimeout: defaultReadTimeout,
		WriteTimeout:      defaultWriteTimeout,
	}

	log.Println("starting the server")

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}

			log.Printf("failed to start HTTP server, error: %v", err)
		}
	}()

	<-ctx.Done()

	log.Println("stopping the server")

	shutdownCtx, httpCancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer httpCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("failed to shutdown HTTP server, error: %v", err)
	}
}
