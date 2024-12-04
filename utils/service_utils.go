package utils

import (
	"github.com/google/uuid"

	"net/http"
	"log"
	"os/signal"
	"context"
	"time"
	"syscall"
)

/*
	Ctrl+C shutdown
	
	-------------------------------
	https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go

*/
func ManualShutdown(server *http.Server) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func GenerateID() uuid.UUID {
	var id = uuid.New()
	log.Printf("UUID Generated. %s\n", id)
	return id
}