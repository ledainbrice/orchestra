package controllers

import (
	"fmt"
	"path/filepath"

	"orchestra/api/models"
	"orchestra/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	fileName := "hermine.jpg"
	//fileName := ctx.Param("filename")
	targetPath := filepath.Join("tmp/", fileName)
	//This ckeck is for example, I not sure is it can prevent all possible filename attacks - will be much better if real filename will not come from user side. I not even tryed this code
	//if !strings.HasPrefix(filepath.Clean(targetPath), "tmp/") {
	//	c.String(403, "Look like you attacking me")
	//	return
	//}
	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)
}

// Ajouter un utilisteur
func PostUser(c *gin.Context) {
	db := config.GetDB("default")

	var user models.User
	c.Bind(&user)

	err := db.Create(&user).Error
	if err != nil {
		// Create failed, do something e.g. return, panic etc.
		c.JSON(422, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, gin.H{"success": user})
	}
}

// Obtenir la liste de tous les utilisateurs
func GetUsers(c *gin.Context) {
	db := config.GetDB("default")
	//out := csv.Read("assets/people.csv")
	//log.Println(out)
	//log.Println("Executing finalHandler", c.Keys["user"])

	var users []models.User
	// SELECT * FROM users
	db.Find(&users)
	// Affichage des données
	c.JSON(200, users)
}

// Obtenir un utilisateur par son id
func GetUser(c *gin.Context) {
	db := config.GetDB("default")

	id := c.Params.ByName("id")
	var user models.User
	// SELECT * FROM users WHERE id = id;
	db.First(&user, id)

	if user.Id != 0 {
		// Affichage des données
		c.JSON(200, user)
	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

// Modifier un utilisateur
func EditUser(c *gin.Context) {
	db := config.GetDB("default")

	// Récupération de l'id dans une variable
	id := c.Params.ByName("id")
	var user models.User
	// SELECT * FROM users WHERE id = id;
	db.First(&user, id)

	if user.Name != "" {
		if user.Id != 0 {
			var json models.User
			c.Bind(&json)

			result := models.User{
				Id:   user.Id,
				Name: json.Name,
			}

			// UPDATE users SET name='json.Name' WHERE id = user.Id;
			db.Model(&user).Update("name", result.Name)
			// Affichage des données modifiées
			c.JSON(200, gin.H{"success": result})
		} else {
			// Affichage de l'erreur
			c.JSON(404, gin.H{"error": "User not found"})
		}

	} else {
		// Affichage de l'erreur
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

// Supprimer un utilisateur
func DeleteUser(c *gin.Context) {
	db := config.GetDB("default")

	// Récupération de l'id dans une variable
	id := c.Params.ByName("id")
	var user models.User
	db.First(&user, id)

	if user.Id != 0 {
		// DELETE FROM users WHERE id = user.Id
		db.Delete(&user)
		// Affichage des données
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		// Affichage de l'erreur
		c.JSON(404, gin.H{"error": "User not found"})
	}
}

func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}

func Login(c *gin.Context) {
	db := config.GetDB("default")

	json := struct {
		email    string
		password string
	}{
		email:    "",
		password: "",
	}
	c.Bind(&json)
	var user models.User
	db.First(&user, 3)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Name,
	})
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	c.JSON(200, config.JwtToken{Token: tokenString})

}
