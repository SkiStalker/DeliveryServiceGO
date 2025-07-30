package main

import (
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
	pb "user-service/grpc_build/user"

	"google.golang.org/grpc"

	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {

	host := os.Getenv("USER_SERVICE_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("USER_SERVICE_PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	user_server := &UserServer{}
	pb.RegisterUserServiceServer(s, user_server)

	healthServer := health.NewServer()
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, healthServer)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("User Service is listening on :50051")
		if err := s.Serve(lis); err != nil && !errors.Is(err, net.ErrClosed) {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	<-stop
	log.Println("Shutting down...")

	server_shutdown := make(chan struct{})

	go func() {
		s.GracefulStop()
		user_server.Close()
		server_shutdown <- struct{}{}
	}()

	select {
	case <-server_shutdown:
	case <-time.After(10 * time.Second):
		s.Stop()
	}
	log.Println("User service exited cleanly.")
}
