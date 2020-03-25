package main

import (
	"log"
	"os"

	"github.com/Flyewzz/group_preparation/db"
	"github.com/Flyewzz/group_preparation/handlers"
	"github.com/Flyewzz/group_preparation/store/db/pg"
	"github.com/spf13/viper"
)

func PrepareConfig() {
	viper.SetConfigFile(os.Args[1])
	// viper.SetConfigFile("config.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read a config file: %v\n", err)
	}
}

func PrepareHandlerData() *handlers.HandlerData {
	db, err := db.ConnectToDB(viper.GetString("db.host"),
		viper.GetString("db.user"), viper.GetString("db.password"), viper.GetString("db.database"))
	if err != nil {
		log.Fatalf("Error with database: %v\n", err)
	}
	universityController := pg.NewUniversityControllerPg(viper.GetInt("university.itemsPerPage"), db)
	subjectController := pg.NewSubjectControllerPg(viper.GetInt("subject.itemsPerPage"), db)
	materialFileController := pg.NewMaterialFileControllerPg(db)
	materialController := pg.NewMaterialControllerPg(viper.GetInt("material.itemsPerPage"), db,
		*materialFileController)
	return handlers.NewHandlerData(universityController,
		subjectController, materialController)
}
