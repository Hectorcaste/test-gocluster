package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var pingError = false

func ClientMariaGormDTB(dbUsername, dbPassword, dbHost, dbTable string) (*gorm.DB, error) {
	if db == nil {
		err := connectDB(dbUsername, dbPassword, dbHost, dbTable)
		if err != nil {
			pingError = true
			return nil, err
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		pingError = true
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil || pingError {
		err = connectDB(dbUsername, dbPassword, dbHost, dbTable)
		if err != nil {
			pingError = true
			return nil, err
		}
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	pingError = false
	return db, nil
}

func connectDB(dbUsername, dbPassword, dbHost, dbTable string) error {
	connect := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbTable)
	fmt.Println("connect", connect)
	var err error
	db, err = gorm.Open(mysql.Open(connect), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})

	if err != nil {
		return err
	}

	return nil
}
