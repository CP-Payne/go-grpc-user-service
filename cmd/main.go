package main

import (
	"context"
	"fmt"
	"go/grpc/userservice/gen"
	"go/grpc/userservice/internal/controller/userdata"
	grpchandler "go/grpc/userservice/internal/handler/grpc"
	"go/grpc/userservice/internal/repository/memory"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	log.Printf("Starting the userdata service on port %s", port)

	repo := memory.New()

	ctrl := userdata.New(repo)

	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	gen.RegisterUserServiceServer(server, h)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup signal handling
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	// Wait for signal
	go func() {
		sig := <-signals
		log.Printf("Received shutdown signal: %v", sig)
		cancel()
	}()

	// Start server and check for success; otherwise cancel
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Printf("grpc server stopped unexpectedly: %v", err)
			cancel()
		}
	}()

	// Graceful shutdown
	go func() {
		<-ctx.Done()
		log.Println("Initiating graceful shutdown...")
		ch := make(chan struct{})

		// Attempts to gracefully stop the server
		go func() {
			server.GracefulStop()
			close(ch)
		}()

		// Waits 30 seconds for the server to stop gracefully; otherwise forcefully stop it
		select {
		case <-ch:
			log.Println("Server gracefully stopped")
		case <-time.After(30 * time.Second):
			log.Println("Graceful shutdown timed out. Stopping the server forcefully.")
			server.Stop()
		}
	}()

	<-ctx.Done()

}
