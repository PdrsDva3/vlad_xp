package main

import (
	"vlad_xp/internal/delivery"
	"vlad_xp/pkg/config"
	"vlad_xp/pkg/database"
)

func main() {
	config.InitConfig()
	db := database.GetDB()

	delivery.Start(db)
}
