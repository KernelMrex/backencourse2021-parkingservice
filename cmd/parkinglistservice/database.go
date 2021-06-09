package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type DBConnectionProperties struct {
	DriverName string
	User       string
	Password   string
	Host       string
	Port       int
	DBName     string
}

func createDBConnection(properties *DBConnectionProperties) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@%s:%d/%s",
		properties.User,
		properties.Password,
		properties.Host,
		properties.Port,
		properties.DBName,
	)

	dbConn, err := sql.Open(properties.DriverName, dataSourceName)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := dbConn.Ping(); err != nil {
		return nil, errors.WithStack(err)
	}

	return dbConn, nil
}
