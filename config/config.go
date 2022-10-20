package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
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
	api := os.Getenv("API_URL")
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(errEnv.Error())
	}
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", dbHost, dbUser, dbPassword, dbName, dbPort)
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
