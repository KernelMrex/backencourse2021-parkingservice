package transport

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	parking "parkinglistservice/api/parkinglistservice"
	"parkinglistservice/pkg/parkinglistservice/infrastructure/repository"
	"testing"
)

const (
	bufSize    = 1024 * 1024
	dialTarget = "bufnet"
)

func TestGetParking(t *testing.T) {
	// Prepare mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock db: %v", err)
	}
	parkingMockRow := sqlmock.NewRows([]string{"name", "description", "coordinates"}).
		AddRow("Test parking", "Test description", "41.40338, 2.17403")
	mock.ExpectQuery("SELECT `name`, `description`, `coordinates` FROM `parking` WHERE `id`=\\?").
		WithArgs("81db895e-11a3-4c7a-8054-237fd188f823").
		WillReturnRows(parkingMockRow)

	// Prepare client
	grpcSrv, listener := startParkingListServiceServer(NewGRPCServer(
		repository.NewParkingRepository(db),
	))
	defer grpcSrv.GracefulStop()

	client, conn := createClient(t, listener)
	defer stopClient(t, conn)

	// Execute request
	resp, err := client.GetParking(context.Background(), &parking.GetParkingRequest{Id: "81db895e-11a3-4c7a-8054-237fd188f823"})
	if err != nil {
		t.Fatalf("GetParking failed: %v", err)
	}

	// Check asserts
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, "Test parking", resp.GetTitle())
	assert.Equal(t, "Test description", resp.GetDescription())
	assert.Equal(t, "41.40338, 2.17403", resp.GetLocation())
}

func TestGetParkingErr(t *testing.T) {
	// Prepare mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock db: %v", err)
	}
	mock.ExpectQuery("SELECT `name`, `description`, `coordinates` FROM `parking` WHERE `id`=\\?").
		WithArgs("81db895e-11a3-4c7a-8054-237fd188f823").
		WillReturnError(sql.ErrNoRows)

	// Prepare client
	grpcSrv, listener := startParkingListServiceServer(NewGRPCServer(
		repository.NewParkingRepository(db),
	))
	defer grpcSrv.GracefulStop()

	client, conn := createClient(t, listener)
	defer stopClient(t, conn)

	// Execute request
	_, err = client.GetParking(context.Background(), &parking.GetParkingRequest{Id: "81db895e-11a3-4c7a-8054-237fd188f823"})

	// Check asserts
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddParking(t *testing.T) {
	// Prepare mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock db: %v", err)
	}
	mock.ExpectExec("INSERT INTO `parking`").
		//WithArgs("Testing name", "Testing description", "41.40338, 2.17403"). // TODO: Should I test DB write here?
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Prepare client
	grpcSrv, listener := startParkingListServiceServer(NewGRPCServer(
		repository.NewParkingRepository(db),
	))
	defer grpcSrv.GracefulStop()

	client, conn := createClient(t, listener)
	defer stopClient(t, conn)

	// Executing request
	resp, err := client.AddParking(context.Background(), &parking.AddParkingRequest{
		Title:       "Testing name",
		Description: "Testing description",
		Location:    "41.40338, 2.17403",
	})
	if err != nil {
		t.Fatalf("GetParking failed: %v", err)
	}

	// Check asserts
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.True(t, isValidUUID(resp.GetId()))
}

func getBufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, s string) (net.Conn, error) {
		return listener.Dial()
	}
}

func createClient(t *testing.T, listener *bufconn.Listener) (parking.ParkingListServiceClient, *grpc.ClientConn) {
	conn, err := grpc.DialContext(
		context.Background(),
		dialTarget,
		grpc.WithContextDialer(getBufDialer(listener)),
		grpc.WithInsecure(),
	)
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	return parking.NewParkingListServiceClient(conn), conn
}

func stopClient(t *testing.T, conn *grpc.ClientConn) {
	if err := conn.Close(); err != nil {
		t.Fatalf("could not close connection: %v", err)
	}
}

func startParkingListServiceServer(srv *GRPCServer) (*grpc.Server, *bufconn.Listener) {
	lis := bufconn.Listen(bufSize)
	s := grpc.NewServer()

	parking.RegisterParkingListServiceServer(s, srv)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	return s, lis
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
