package mariadb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"neosync/internal/logger"
	"time"
)

type DB struct {
	config Config
	db     *sql.DB
}

type Scanner interface {
	Scan(dest ...interface{}) error
}

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Port     int    `koanf:"port"`
	Host     string `koanf:"host"`
	DBName   string `koanf:"db_name"`
}

func (c Config) String() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", c.Username, c.Password, c.Host, c.Port, c.DBName)
}

func (m *DB) Conn() *sql.DB {
	return m.db
}

func New(config Config) *DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		config.Username, config.Password, config.Host, config.Port, config.DBName),
	)
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	logger.L().Info("connected to mariadb")

	return &DB{config: config, db: db}
}
