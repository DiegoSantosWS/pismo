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
	"time"

	// Used pg drive on sqlx
	_ "github.com/lib/pq"
)

func init() {
	utils.Load()
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	ctx, cancel, srv := start()
	defer shutdown(ctx, cancel, wait, srv)

}

func start() (ctx context.Context, cancel context.CancelFunc, srv *http.Server) {
	ctx, cancel = context.WithCancel(context.Background())
	srv = router.NewRouter()
	connection.Load(ctx)
	return
}

func shutdown(ctx context.Context, cancel context.CancelFunc, wait time.Duration, srv *http.Server) {
	cancel()
	connection.Close()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel = context.WithTimeout(context.Background(), wait)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
}
