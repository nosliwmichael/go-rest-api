package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nosliwmichael/go-rest-api/internal/configuration"
	"log"
	"os"
	"sync"
)

var (
	instance *sql.DB
	once     sync.Once
)

func GetDbInstance(config *configuration.Config) *sql.DB {
	once.Do(func() {
		username, password := getCredentials()
		connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, config.Database.Host, config.Database.Port, config.Database.SID)
		var err error
		instance, err = sql.Open("mysql", connectionStr)
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		if err = instance.Ping(); err != nil {
			log.Fatalf("Error pinging database: %v", err)
		}
	})
	return instance
}

func getCredentials() (username string, password string) {
	// TODO: Setup proper credential management
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	return
}
