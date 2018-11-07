package models

import (
	"log"
	"orchestra/config"
	uuid_func "orchestra/functions/Uuid"
	"time"

	"github.com/satori/go.uuid"
)

// La structure de données
type Node_request struct {
	Id        uuid.UUID `gorm:"primary_key;type:char(36);column:id" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Mapper    string `gorm:"text;default:'{}';not null" json:"mapper"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
	Url       string `gorm:"type:varchar(255);not null" json:"url"`
	Sync      bool   `gorm:"type:boolean;not null" json:"sync"`
	Method    string `gorm:"type:varchar(255);not null" json:"method"`
	Node_id   uuid.UUID
}

func (node_request Node_request) Validator() (err error) {
	/*if !validation.ValidateEmail(node_request.field) {
		err = errors.New("client email invalide")
	}*/
	return
}

func (node_request *Node_request) BeforeCreate() (err error) {
	if uuid.Equal(node_request.Id, uuid_func.Format("00000000-0000-0000-0000-000000000000")) {
		node_request.Id = uuid.Must(uuid.NewV4())
	}
	if node_request.Mapper == "" {
		node_request.Mapper = "{}"
	}
	err = node_request.Validator()
	return
}

// Migration
func MigrateNode_request() {
	db := config.GetDB("default")
	if !db.HasTable(&Node_request{}) {
		db.CreateTable(&Node_request{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Node_request{})
	}
}

// Seed
func SeedNode_request() {
	db := config.GetDB("default")
	if db.HasTable(&Node_request{}) {
		nodes_requests := []Node_request{
			{Id: uuid_func.Format("00000000-0000-0000-0000-000000000001"), Name: "request 1", Url: "node 2", Sync: false, Method: "GET", Node_id: uuid_func.Format("00000000-0000-0000-0000-000000000001")},
		}
		for _, node_request := range nodes_requests {
			db.Create(&node_request)
			log.Println(node_request)
		}
	}
}
