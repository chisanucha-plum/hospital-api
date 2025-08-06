package main

import (
	"hospital-api/database"
	"hospital-api/internal/configs"
	"hospital-api/internal/handlers"
	"log"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.InitSchema(db); err != nil {
		log.Fatal(err)
	}

	r := handlers.SetupRouter(db)
	r.Run(":" + configs.Envs.Port)
}
