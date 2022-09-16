package platform

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func RunApp() error {
	ctx, cancel := context.WithCancel(context.Background())

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		trapped := <-gracefulStop
		log.Printf("shutting down the HTTP server beacuse of %q signal\n", trapped.String())
		cancel()
	}()

	RunHTTPServer(ctx, NewRouter())

	return nil
}
