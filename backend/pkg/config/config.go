package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const (
	DBName     = "DB_NAME"
	DBUser     = "DB_USER"
	DBPassword = "DB_PASSWORD"
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
)

func InitConfig() {
	envPath, _ := os.Getwd()
	envPath = filepath.Join(envPath, "..")
	envPath = filepath.Join(envPath, "/backend")
	envPath = filepath.Join(envPath, "/deploy")

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(envPath)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to init config. Error:%v", err.Error()))
	}
}
