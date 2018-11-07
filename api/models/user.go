package models

import (
	"log"
	"orchestra/config"
)

// La structure de données
type User struct {
	Id   int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name string `gorm:"not null" form:"name" json:"name"`
}

// Migration
func MigrateUser() {
	db := config.GetDB("default")
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	}
}

// Seed
func SeedUser() {
	db := config.GetDB("default")
	if db.HasTable(&User{}) {
		users := []User{
			{Id: 1, Name: "Aramis"},
			{Id: 2, Name: "Athos"},
			{Id: 3, Name: "Porthos"},
			{Id: 4, Name: "D'Artagnan"},
		}
		for _, user := range users {
			db.Create(&user)
			log.Println(user)
		}
	}
}
