package controllers

import (
	"log"
	"orchestra/api/models"
	"orchestra/config"

	"github.com/gin-gonic/gin"
)

// Execution du node
func RunNode(c *gin.Context) {
	db := config.GetDB("default")

	id := c.Params.ByName("id")
	log.Println("Node is running", id)
	var node models.Node
	uid0 := node.Id
	// SELECT * FROM clients WHERE id = id;
	db.Where("id = ?", id).First(&node)

	if node.Id != uid0 {
		// Affichage des données
		c.JSON(200, node)
	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "Node not found"})
	}
}

// Ajouter un node
func AddNode(c *gin.Context) {
	db := config.GetDB("default")

	var node models.Node
	c.Bind(&node)
	err := db.Create(&node).Error
	if err != nil {
		// Create failed, do something e.g. return, panic etc.
		c.JSON(422, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, gin.H{"success": node})
	}

}

// Obtenir la liste de tous les nodes
func IndexNode(c *gin.Context) {
	db := config.GetDB("default")

	var nodes []models.Node
	// SELECT * FROM nodes
	db.Preload("Node_requests").Find(&nodes)
	// Affichage des données
	c.JSON(200, nodes)
}

// Obtenir un node par son id
func ViewNode(c *gin.Context) {
	db := config.GetDB("default")

	id := c.Params.ByName("id")
	var node models.Node
	uid0 := node.Id
	// SELECT * FROM nodes WHERE id = id;
	db.Preload("Node_requests").Where("id = ?", id).First(&node)

	if node.Id != uid0 {
		// Affichage des données
		c.JSON(200, node)
	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "Node not found"})
	}
}

// Modifier un node
func EditNode(c *gin.Context) {
	db := config.GetDB("default")

	// Récupération de l'id dans une variable
	id := c.Params.ByName("id")
	var node models.Node
	// SELECT * FROM nodes WHERE id = id;
	db.First(&node, id)

	if true {
		var node_update models.Node
		c.Bind(&node_update)
		node_update.Id = node.Id

		// UPDATE nodes SET (field1, field2 ...) VALUES ('value 1', 'value 2' ...) WHERE id = node_update.Id;
		if db.Save(&node_update).Error != nil {
			// Update failed, do something e.g. return, panic etc.
			c.JSON(422, gin.H{"error": "Fields are empty"})
		}
		// Affichage des données modifiées
		c.JSON(200, gin.H{"success": node_update})
	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "Node not found"})
	}
}

// Supprimer un node
func DeleteNode(c *gin.Context) {
	db := config.GetDB("default")

	// Récupération de l'id dans une variable
	id := c.Params.ByName("id")
	var node models.Node
	uid0 := node.Id
	db.Where("id = ?", id).First(&node)

	if node.Id != uid0 {
		// DELETE FROM nodes WHERE id = Node.Id
		if db.Delete(&node).Error != nil {
			// Delete failed, do something e.g. return, panic etc.
			c.JSON(422, gin.H{"error": "Cannot delete Node"})
		} else {
			// Affichage des données
			c.JSON(200, gin.H{"success": "Node #" + id + " deleted"})
		}

	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "Node not found"})
	}
}
