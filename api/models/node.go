package models

import (
	"errors"
	"log"
	"orchestra/config"
	uuid_func "orchestra/functions/Uuid"
	"time"

	"github.com/satori/go.uuid"
)

// La structure de données
type Node struct {
	Id            uuid.UUID `gorm:"primary_key;type:char(36);column:id" json:"id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Name          string         `gorm:"type:varchar(255);not null" json:"name"`
	Node_requests []Node_request `gorm:"foreignkey:node_id" json:"node_requests"`
}

func (node Node) Validator() (err error) {
	if node.Name == "" {
		err = errors.New("Node must have a name")
	}
	return
}

func (node *Node) BeforeCreate() (err error) {
	if uuid.Equal(node.Id, uuid_func.Format("00000000-0000-0000-0000-000000000000")) {
		node.Id = uuid.Must(uuid.NewV4())
	}
	err = node.Validator()
	return
}

// Migration
func MigrateNode() {
	db := config.GetDB("default")
	if !db.HasTable(&Node{}) {
		db.CreateTable(&Node{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Node{})
	}
}

// Seed
func SeedNode() {
	db := config.GetDB("default")
	if db.HasTable(&Node{}) {
		nodes := []Node{
			{Id: uuid_func.Format("00000000-0000-0000-0000-000000000001"), Name: "Node 1"},
			{Id: uuid_func.Format("00000000-0000-0000-0000-000000000002"), Name: "Node 2"},
			{Id: uuid_func.Format("00000000-0000-0000-0000-000000000003"), Name: "Node 3"},
			{Id: uuid_func.Format("00000000-0000-0000-0000-000000000004"), Name: "Node 4"},
		}
		for _, node := range nodes {
			db.Create(&node)
			log.Println(node)
		}
	}
}

//Execution of node
func (node *Node) Execute(data string) (err error) {
	return
}
