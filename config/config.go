package config

import (
	"fmt"

	_ "github.com/lib/pq"
)

type ApiConfig struct {
	Url string
}

type DbConfig struct {
	DataSourceName string
}

type Config struct {
	ApiConfig
	DbConfig
}

func (c *Config) readConfig() {
	api := ":8888"
	dbHost := "localhost"
	dbName := "api_warung_makan"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "root"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	fmt.Println(dsn)
	c.ApiConfig = ApiConfig{
		Url: api,
	}
	c.DbConfig = DbConfig{
		DataSourceName: dsn,
	}
}

func NewConfig() Config {
	conf := Config{}
	conf.readConfig()
	return conf
}

// api := os.Getenv("API_URL")
// dbHost := os.Getenv("DB_HOST")
// dbName := os.Getenv("DB_NAME")
// dbPort := os.Getenv("DB_PORT")
// dbUser := os.Getenv("DB_USER")
// dbPassword := os.Getenv("DB_PASSWORD")
