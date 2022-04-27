package database

import (
	"log"

	"github.com/henryhiraki22/lifts/entity"
	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Database connected")
	return nil
}

func Migrate(table *entity.LiftA, tableb *entity.LiftB) {
	Connector.AutoMigrate(&table)
	Connector.AutoMigrate(&tableb)
	log.Println("Table migrated")
}
