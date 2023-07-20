package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wesleywu/go-lifespan/example/bar"
	"github.com/wesleywu/go-lifespan/lifespan"
)

// main demonstrating the usage of OnBoostrap and OnShutdown
func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	lifespan.OnBootstrap(ctx)

	err := bar.CallFoo()
	if err != nil {
		log.Printf(err.Error())
	}

	initSignals(ctx, lifespan.OnShutdown)
}

func initSignals(ctx context.Context, onShutdown func(ctx2 context.Context)) {
	var (
		sigChan = make(chan os.Signal, 1)
		sig     os.Signal
	)
	signal.Notify(
		sigChan,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGABRT,
	)
	for {
		select {
		case <-ctx.Done():
			log.Print("context timeout reached, gracefully shutting down\n")
			onShutdown(ctx)
			return
		case sig = <-sigChan:
			log.Printf("signal received: %s, gracefully shutting down\n", sig.String())
			onShutdown(ctx)
			return
		}
	}
}
