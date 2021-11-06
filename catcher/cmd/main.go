package main

import (
	"catcher/internal/http"
	"catcher/internal/store/inmemory"
	"context"
	"log"
)

func main() {
	store := inmemory.NewDB()

	srv := http.NewServer(context.Background(), ":8080", store)
	if err := srv.Run(); err != nil {
		log.Println(err)
	}

	srv.WaitForGracefulTermination()
}
