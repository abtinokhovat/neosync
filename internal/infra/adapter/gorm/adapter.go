package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"neosync/internal/infra/adapter/mariadb"
	"time"
)

// DB is a wrapper around the GORM DB instance.
type DB struct {
	config mariadb.Config
	db     *gorm.DB
}

// New creates a new DB instance and connects to the database using GORM.
func New(config mariadb.Config) *DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Username, config.Password, config.Host, config.Port, config.DBName)

	// Open a connection to the database using GORM and the mysql driver.
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // GORM logger with info level for debugging.
	})
	if err != nil {
		panic(fmt.Errorf("can't open MariaDB connection: %v", err))
	}

	// Configure connection pool settings.
	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(fmt.Errorf("can't get sql.DB from GORM: %v", err))
	}
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	// Log successful connection.
	logger.Default.Info(nil, "connected to MariaDB")

	return &DB{config: config, db: gormDB}
}

// Conn returns the underlying GORM DB connection.
func (m *DB) Conn() *gorm.DB {
	return m.db
}
