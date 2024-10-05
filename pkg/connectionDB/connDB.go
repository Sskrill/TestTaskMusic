package connDB

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"log"
)

type ConfigParamDB struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
	Sslmode  string
}

func NewDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error env")

	}
	var cfgDB ConfigParamDB
	if err := envconfig.Process("db", &cfgDB); err != nil {
		log.Fatal("cant take env")
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfgDB.Host, cfgDB.Port, cfgDB.User, cfgDB.Password, cfgDB.DBname, cfgDB.Sslmode))

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
