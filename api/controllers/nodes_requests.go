package controllers

import (
	"orchestra/api/models"
	"orchestra/config"

	"github.com/gin-gonic/gin"
)

// Ajouter un node_request
func AddNode_request(c *gin.Context) {
	db := config.GetDB("default")

	var node_request models.Node_request
	c.Bind(&node_request)
	err := db.Create(&node_request).Error
	if err != nil {
		// Create failed, do something e.g. return, panic etc.
		c.JSON(422, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, gin.H{"success": node_request})
	}

}

// Obtenir la liste de tous les nodes_requests
func IndexNode_request(c *gin.Context) {
	db := config.GetDB("default")

	var nodes_requests []models.Node_request
	// SELECT * FROM nodes_requests
	db.Find(&nodes_requests)
	// Affichage des données
	c.JSON(200, nodes_requests)
}

// Obtenir un node_request par son id
func ViewNode_request(c *gin.Context) {
	db := config.GetDB("default")

	id := c.Params.ByName("id")
	var node_request models.Node_request
	uid0 := node_request.Id
	// SELECT * FROM nodes_requests WHERE id = id;
	db.Where("id = ?", id).First(&node_request)

	if node_request.Id != uid0 {
		// Affichage des données
		c.JSON(200, node_request)
	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "Node_request not found"})
	}
}

// Modifier un node_request
func EditNode_request(c *gin.Context) {
	db := config.GetDB("default")

	// Récupération de l'id dans une variable
	id := c.Params.ByName("id")
	var node_request models.Node_request
	// SELECT * FROM nodes_requests WHERE id = id;
	db.First(&node_request, id)

	if true {
		var node_request_update models.Node_request
		c.Bind(&node_request_update)
		node_request_update.Id = node_request.Id

		// UPDATE nodes_requests SET (field1, field2 ...) VALUES ('value 1', 'value 2' ...) WHERE id = node_request_update.Id;
		if db.Save(&node_request_update).Error != nil {
			// Update failed, do something e.g. return, panic etc.
			c.JSON(422, gin.H{"error": "Fields are empty"})
		}
		// Affichage des données modifiées
		c.JSON(200, gin.H{"success": node_request_update})
	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "Node_request not found"})
	}
}

// Supprimer un node_request
func DeleteNode_request(c *gin.Context) {
	db := config.GetDB("default")

	// Récupération de l'id dans une variable
	id := c.Params.ByName("id")
	var node_request models.Node_request
	uid0 := node_request.Id
	db.Where("id = ?", id).First(&node_request)

	if node_request.Id != uid0 {
		// DELETE FROM nodes_requests WHERE id = Node_request.Id
		if db.Delete(&node_request).Error != nil {
			// Delete failed, do something e.g. return, panic etc.
			c.JSON(422, gin.H{"error": "Cannot delete Node_request"})
		} else {
			// Affichage des données
			c.JSON(200, gin.H{"success": "Node_request #" + id + " deleted"})
		}

	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "Node_request not found"})
	}
}
