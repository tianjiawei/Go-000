package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	errgroup, ctx := errgroup.WithContext(context.Background())
	stop := make(chan struct{})
	sign := make(chan os.Signal)

	signal.Notify(sign, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
	server := &http.Server{Addr: ":8080"}

	errgroup.Go(func() error {
		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world!"))
		})
		return server.ListenAndServe()
	})

	errgroup.Go(func() error {
		select {
		case <-stop:
			return server.Shutdown(ctx)
		}
	})

	errgroup.Go(func() error {
		select {
		case <-sign:
		}
		server.Shutdown(ctx)
		fmt.Println("http server drop out success！")
		return nil
	})
	errgroup.Wait()
}
