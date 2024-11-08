package dbsql

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// InitDB initializes the database connection
func InitDB() error {
	// Get database configuration from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE") // usually "require" or "disable"
	password := os.Getenv("DB_PASSWORD")
	config := Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
		SSLMode:  sslMode,
	}
	return InitDBWithConfig(&config)
}

func InitDBWithConfig(config *Config) error {
	// Construct DSN from environment variables
	tls := "&tls=tidb&parseTime=true"
	dsn := fmt.Sprintf("%s:%s@%s/%s?%s",
		config.User, config.Password, config.Host, config.DBName, tls)

	// Setup database connection
	_, err := SetupDB(dsn)
	return err
}

// SetupDB sets up the database connection
func SetupDB(dsn string) (*gorm.DB, error) {
	var err error

	// Initialize database connection
	err = mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "gateway01.ap-southeast-1.prod.aws.tidbcloud.com",
	})
	if err != nil {
		return nil, err
	}

	dbSQL, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = dbSQL.Ping()
	if err != nil {
		return nil, err
	}

	// Convert *sql.DB to *gorm.DB
	db, err = gorm.Open(gmysql.New(gmysql.Config{Conn: dbSQL}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// SetDB sets the database connection for testing
func SetDB(newDB *gorm.DB) {
	db = newDB
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}
