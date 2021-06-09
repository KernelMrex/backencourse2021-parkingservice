package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	parking "parkinglistservice/api/parkinglistservice"
	"parkinglistservice/pkg/parkinglistservice/infrastructure/repository"
	"parkinglistservice/pkg/parkinglistservice/infrastructure/transport"
)

const appID = "parkingservice"

type config struct {
	GRPCServeAddress string `envconfig:"serve_address"`
	DBUser           string `envconfig:"mysql_user"`
	DBPassword       string `envconfig:"mysql_password"`
	DBHost           string `envconfig:"mysql_host"`
	DBPort           int    `envconfig:"mysql_port"`
	DBDBName         string `envconfig:"mysql_database"`
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	var c config
	if err := envconfig.Process(appID, &c); err != nil {
		log.Fatalf("could not parsing config: %v", err)
	}

	log.Info("Connecting to database")
	dbConn, err := createDBConnection(&DBConnectionProperties{
		DriverName: "mysql",
		User:       c.DBUser,
		Password:   c.DBPassword,
		Host:       c.DBHost,
		Port:       c.DBPort,
		DBName:     c.DBDBName,
	})
	if err != nil {
		log.Fatalf("could not connect database: %v", err)
	}

	srv := transport.NewGRPCServer(repository.NewParkingRepository(dbConn))

	log.Info("Server is starting")
	grpcServer := startServer(c.GRPCServeAddress, srv)

	waitForKillSignal(getKillSignalChan())
	log.Info("Server is stopping")
	grpcServer.GracefulStop()
}

func startServer(serveURL string, srv *transport.GRPCServer) *grpc.Server {
	grpcServer := grpc.NewServer()
	parking.RegisterParkingListServiceServer(grpcServer, srv)

	listener, err := net.Listen("tcp", serveURL)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return grpcServer
}
