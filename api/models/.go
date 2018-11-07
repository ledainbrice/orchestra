package models

import (
	"log"
	"orchestra/config"

	"github.com/gobuffalo/uuid"
)

// La structure de données
type  struct {
	ID   uuid.UUID `gorm:"primary_key;type:char(36);column:id" form:"id" json:"id"`
}

// Migration
func Migrate() {
	db := config.GetDB("default")
	if !db.HasTable(&{}) {
		db.CreateTable(&{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&{})
	}
}

// Seed
func Seed() {
	db := config.GetDB("default")
	if db.HasTable(&{}) {
		 := []{
			{Name: " 1"},
			{Name: " 2"},
			{Name: " 3"},
			{Name: " 4"},
		}
		for _,  := range  {
			db.Create(&)
			log.Println()
		}
	}
}
