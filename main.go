package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pismo/connection"
	"pismo/router"
	"pismo/utils"
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

	ctx, cancel, srv := start()
	defer shutdown(ctx, cancel, wait, srv)
	waitShutdown()
}

func start() (ctx context.Context, cancel context.CancelFunc, srv *http.Server) {
	ctx, cancel = context.WithCancel(context.Background())
	srv = router.NewRouter()
	connection.Load(ctx)
	return
}

func shutdown(ctx context.Context, cancel context.CancelFunc, wait time.Duration, srv *http.Server) {
	cancel()
	// Create a deadline to wait for.
	ctx, cancel = context.WithTimeout(ctx, wait)
	defer cancel()

	connection.Close()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
}

// waitShutdown waits until is going to die
func waitShutdown() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	s := <-sigc
	log.Printf("[ WaitShutdown ] signal received [%v] canceling everything", s)
}
