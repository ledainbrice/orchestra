package controllers

import (
	"orchestra/api/models"
	"orchestra/config"

	"github.com/gin-gonic/gin"
)

// Ajouter un client
func AddClient(c *gin.Context) {
	db := config.GetDB("default")

	var client models.Client
	c.Bind(&client)
	err := db.Create(&client).Error
	if err != nil {
		// Create failed, do something e.g. return, panic etc.
		c.JSON(422, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, gin.H{"success": client})
	}

}

// Obtenir la liste de tous les clients
func IndexClient(c *gin.Context) {
	db := config.GetDB("default")

	var clients []models.Client
	// SELECT * FROM clients
	db.Find(&clients)
	// Affichage des données
	c.JSON(200, clients)
}

// Obtenir un client par son id
func ViewClient(c *gin.Context) {
	db := config.GetDB("default")

	id := c.Params.ByName("id")
	var client models.Client
	uid0 := client.Id
	// SELECT * FROM clients WHERE id = id;
	db.Where("id = ?", id).First(&client)

	if client.Id != uid0 {
		// Affichage des données
		c.JSON(200, client)
	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "Client not found"})
	}
}

// Modifier un client
func EditClient(c *gin.Context) {
	db := config.GetDB("default")

	// Récupération de l'id dans une variable
	id := c.Params.ByName("id")
	var client models.Client
	// SELECT * FROM clients WHERE id = id;
	db.First(&client, id)

	if true {
		var client_update models.Client
		c.Bind(&client_update)
		client_update.Id = client.Id

		// UPDATE clients SET (field1, field2 ...) VALUES ('value 1', 'value 2' ...) WHERE id = client_update.Id;
		if db.Save(&client_update).Error != nil {
			// Update failed, do something e.g. return, panic etc.
			c.JSON(422, gin.H{"error": "Fields are empty"})
		}
		// Affichage des données modifiées
		c.JSON(200, gin.H{"success": client_update})
	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "Client not found"})
	}
}

// Supprimer un client
func DeleteClient(c *gin.Context) {
	db := config.GetDB("default")

	// Récupération de l'id dans une variable
	id := c.Params.ByName("id")
	var client models.Client
	uid0 := client.Id
	db.Where("id = ?", id).First(&client)

	if client.Id != uid0 {
		// DELETE FROM clients WHERE id = Client.Id
		if db.Delete(&client).Error != nil {
			// Delete failed, do something e.g. return, panic etc.
			c.JSON(422, gin.H{"error": "Cannot delete Client"})
		} else {
			// Affichage des données
			c.JSON(200, gin.H{"success": "Client #" + id + " deleted"})
		}

	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "Client not found"})
	}
}
