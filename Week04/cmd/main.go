package main

import (
	"Go-000/Week04/internal/server"
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
	sign := make(chan os.Signal)

	signal.Notify(sign, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
	server := &http.Server{Addr: ":8080"}

	errgroup.Go(func() error {
		return startServer()
	})

	errgroup.Go(func() error {
		return shutDownServer(ctx, sign, server)
	})
	errgroup.Wait()
}

func startServer() error {
	router := server.NewHttp()
	return router.Run(":8080")
}

func shutDownServer(ctx context.Context, sign chan os.Signal, server *http.Server) error {
	select {
	case <-sign:
	}
	e := server.Shutdown(ctx)
	if e == nil {
		fmt.Println("http server drop out success！")
	} else {
		fmt.Println("http server drop out failed！")
	}
	//todo 释放资源
	return nil
}
