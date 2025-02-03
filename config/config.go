package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

const (
	envDatabaseURL   = "DATABASE_URL"
	envHttpTimeout   = "HTTP_TIMEOUT"
	envApiToken      = "API_TOKEN"
	envApiUrl        = "CRM_API_URL"
	envDatabaseTable = "DATABASE_TABLE"
)

var (
	failVariableSet = "%s environment variable not set"
	failParse       = "failed to parse %s: %v"
)

type HttpConfig struct {
	Timeout time.Duration
}

type DatabaseConfig struct {
	Dsn   string
	Table string
}

type ApiConfig struct {
	Apikey string
	Url    string
}

type Config struct {
	Database *DatabaseConfig
	Api      *ApiConfig
	Http     *HttpConfig
}

func New(configPath string) (*Config, error) {
	if err := godotenv.Load(configPath); err != nil {
		return nil, err
	}

	httpCfg, err := loadHttp()
	if err != nil {
		return nil, err
	}

	dbCfg, err := loadDatabase()
	if err != nil {
		return nil, err
	}

	apiCfg, err := loadApi()
	if err != nil {
		return nil, err
	}

	//TODO: api env(dev, test, prod)

	return &Config{
		Http:     httpCfg,
		Database: dbCfg,
		Api:      apiCfg,
	}, nil
}

func loadDatabase() (*DatabaseConfig, error) {
	dsn := getEnv(envDatabaseURL, "")
	if dsn == "" {
		return nil, fmt.Errorf(failVariableSet, envDatabaseURL)
	}
	table := getEnv(envDatabaseTable, "")
	if table == "" {
		return nil, fmt.Errorf(failVariableSet, envDatabaseTable)
	}
	return &DatabaseConfig{Dsn: dsn, Table: table}, nil
}

func loadHttp() (*HttpConfig, error) {
	httpTimeoutEnv := getEnv(envHttpTimeout, "10")
	httpTimeout, err := strconv.ParseInt(httpTimeoutEnv, 10, 64)
	if err != nil {
		return nil, fmt.Errorf(failParse, envHttpTimeout, err)
	}
	return &HttpConfig{Timeout: time.Duration(httpTimeout) * time.Second}, nil
}

func loadApi() (*ApiConfig, error) {
	apiToken := getEnv(envApiToken, "")
	if apiToken == "" {
		return nil, fmt.Errorf(failVariableSet, envApiToken)
	}
	crmApiUrl := getEnv(envApiUrl, "")
	if crmApiUrl == "" {
		return nil, fmt.Errorf(failVariableSet, envApiUrl)
	}

	return &ApiConfig{Apikey: apiToken, Url: crmApiUrl}, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
