package models

import (
	"errors"
	"log"
	"orchestra/config"
	validation "orchestra/functions/Validation"
	"time"

	"github.com/satori/go.uuid"
)

// La structure de données
type Client struct {
	Id        uuid.UUID `gorm:"primary_key;type:char(36);column:id" form:"id" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Email     string `gorm:"type:varchar(255)";not null`
}

func (client Client) Validator() (err error) {
	if !validation.ValidateEmail(client.Email) {
		err = errors.New("client email invalide")
	}
	return
}

func (client *Client) BeforeCreate() (err error) {
	client.Id = uuid.Must(uuid.NewV4())
	err = client.Validator()
	return
}

// Migration
func MigrateClient() {
	db := config.GetDB("default")
	if !db.HasTable(&Client{}) {
		db.CreateTable(&Client{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Client{})
	}
}

// Seed
func SeedClient() {
	db := config.GetDB("default")
	if db.HasTable(&Client{}) {
		clients := []Client{}
		for _, client := range clients {
			db.Create(&client)
			log.Println(client)
		}
	}
}
