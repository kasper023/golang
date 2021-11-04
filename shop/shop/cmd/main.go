package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"shop/internal/transport/http"
	"time"
)

func main() {
	addr := os.Getenv("SERVICE_ADDR")

	_, cancel := context.WithCancel(context.Background())

	errCh := make(chan error, 1)

	go func() {
		sigCh := make(chan os.Signal)
		signal.Notify(sigCh)
		errCh <- fmt.Errorf("service is shutting down. Signal %d", <-sigCh)
	}()

	go http.StartServer(addr, errCh)

	log.Printf("Service started on port: %s\n", addr)

	err := <-errCh

	cancel()

	log.Println("Gracefully stop service")

	<- time.NewTicker(5 * time.Second).C

	log.Printf("Service terminated: %v\n", err)
}
