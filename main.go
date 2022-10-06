package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"pismo/connection"
	"pismo/utils"
	"pismo/webservice"
	"syscall"
	"time"

	// Used pg drive on sqlx
	_ "github.com/lib/pq"
)

func init() {
	utils.Load()
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*3, "the duration for which the server gracefully wait for existing connections to finish - e.g. 3s or 1m")
	flag.Parse()

	cancel := start()
	defer shutdown(cancel, wait)
	waitShutdown()
}

func start() (cancel context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	webservice.NewRouter()
	connection.Load(ctx)
	return
}

func shutdown(cancel context.CancelFunc, wait time.Duration) {
	cancel()
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	connection.Close()
	webservice.Shutdown(ctx)
}

// waitShutdown waits until is going to die
func waitShutdown() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-sigc
	log.Printf("[ WaitShutdown ] signal received [%v] canceling everything", s)
}
