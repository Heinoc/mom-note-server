package main

import (
	"context"
	"log"
	"mom-note-server/routers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/**
 * author: heinoc
 */

func main() {

	//gin.SetMode(gin.ReleaseMode)

	router := routers.InitRouter()

	srv := &http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
