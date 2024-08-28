package main

import (
	"vlad_xp/internal/delivery"
	"vlad_xp/pkg/config"
	"vlad_xp/pkg/database"
	"vlad_xp/pkg/log"
)

func main() {
	logger, infoFile, errorFile := log.InitLogger()
	defer infoFile.Close()
	defer errorFile.Close()

	config.InitConfig()
	logger.Info("config init success")

	db := database.GetDB()
	logger.Info("db init success")

	delivery.Start(db, logger)
}
