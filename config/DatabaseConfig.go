package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	DBUser string `yaml:"db_user"`
	DBPass string `yaml:"db_pass"`
	DBHost string `yaml:"db_host"`
	DBPort int    `yaml:"db_port"`
	DBName string `yaml:"db_name"`
}

// LoadConfig loads the DB configuration based on APP_ENV
func LoadConfig() (*DBConfig, error) {
	// Get environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local" // Default to local if APP_ENV is not set
	}

	// Determine config file based on environment
	configFile := fmt.Sprintf("/Users/laykis/GolandProjects/devlog/config/DatabaseConfig.yaml")

	// Read YAML file
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse YAML into struct
	var configMap map[string]DBConfig
	err = yaml.Unmarshal(data, &configMap)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling YAML: %v", err)
	}

	config, ok := configMap[env]
	if !ok {
		return nil, fmt.Errorf("Invalid environment: %s", env)
	}

	return &config, nil
}

func GetDBInstance(config *DBConfig) (*gorm.DB, error) {
	// DB DSN(Data Source Name) 생성
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	fmt.Println(dsn)
	// GORM 연결
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Error opening DB: %v", err)
	}

	// DB 연결 확인
	if err := db.DB().Ping(); err != nil {
		return nil, fmt.Errorf("DB connection error: %v", err)
	}

	return db, nil
}
