package main

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	parking "parkinglistservice/api/parkinglistservice"
	"parkinglistservice/pkg/parkinglistservice/infrastructure/repository"
	"parkinglistservice/pkg/parkinglistservice/infrastructure/transport"
	"syscall"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	log.Info("Server is starting...")
	server := startServer(os.Getenv("SERVE_ADDRESS"), transport.NewGRPCServer(repository.NewParkingRepository(nil)))
	log.Info("Server is running")
	waitForKillSignal(getKillSignalChan())
	log.Info("Server is stopping...")
	server.GracefulStop()
	log.Info("Server has stopped")
}

func startServer(serveURL string, srv *transport.GRPCServer) *grpc.Server {
	baseServer := grpc.NewServer()
	parking.RegisterParkingListServiceServer(baseServer, srv)

	listener, err := net.Listen("tcp", serveURL)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}

	go func() {
		if err := baseServer.Serve(listener); err != nil {
			log.Errorf("failed to serve: %v", err)
		}
	}()

	return baseServer
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}
