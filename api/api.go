package api

// Les imports de librairies
import (
	"fmt"
	"log"
	"strings"
	"time"

	"orchestra/api/controllers"
	"orchestra/api/models"
	"orchestra/config"

	"github.com/auth0-community/auth0"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JWTKey = []byte("secret")
var validator *auth0.JWTValidator

// Migration de la BDD
func InitDb() {
	log.Println("migration start")
	models.MigrateUser()
	models.MigrateClient()
	models.MigrateNode()
	models.MigrateNode_request()
	log.Println("migration end")
}

// Seed de la BDD
func SeedDb() {
	log.Println("seed start")
	models.SeedUser()
	models.SeedClient()
	models.SeedNode()
	models.SeedNode_request()
	log.Println("seed end")
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func Handlers() *gin.Engine {
	// Initialisation du serveur MUX
	InitDb()
	SeedDb()
	//var reciever config.Reciever
	//reciever = config.NewReciever("brice@sportintown.com", map[string]interface{}{
	//	"floor": "azertyuiop",
	//})
	//config.SendEmail("default", "demarrage app", "petit poney %recipient.floor%", []config.Reciever{reciever})
	//for i := 0; i < 10; i++ {
	//	k := rand.Intn(999)
	//	config.AddTask(func() {
	//		fmt.Printf("I am worker! Number %d values %d \n", k, digits(k))
	//	})
	//}
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.Use(Cors())
	//r.Use(DbMiddleware())
	r.GET("down", controllers.Download)
	r.POST("login", controllers.Login)
	v1Users := r.Group("api/v1/users", TokenAuthMiddleware())
	{
		v1Users.POST("", controllers.PostUser)
		v1Users.GET("", controllers.GetUsers)
		v1Users.GET(":id", controllers.GetUser)
		v1Users.PUT(":id", controllers.EditUser)
		v1Users.DELETE(":id", controllers.DeleteUser)
		v1Users.OPTIONS("", controllers.OptionsUser)    // POST
		v1Users.OPTIONS(":id", controllers.OptionsUser) // PUT, DELETE
	}
	v1Clients := r.Group("api/v1/clients")
	{
		v1Clients.GET("", controllers.IndexClient)
		v1Clients.GET("post", controllers.AddClient)
		v1Clients.GET("delete/:id", controllers.DeleteClient)
		v1Clients.GET("view/:id", controllers.ViewClient)
	}
	v1Nodes := r.Group("api/v1/nodes")
	{
		v1Nodes.GET("", controllers.IndexNode)
		v1Nodes.GET(":id", controllers.ViewNode)
	}
	return r
}

func DbMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", config.GetDB("default"))
		c.Next()
	}
}

// Activation du CORS
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Récupération du paramètre "token" dans une variable
		authorizationHeader := c.GetHeader("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					c.JSON(403, gin.H{"error": error.Error()})
					c.Abort()
					return
				}
				if token.Valid {
					log.Println("token.Claims")
					log.Println(token.Claims)
					c.Set("user", token.Claims)
					c.Next()
				} else {
					c.JSON(403, gin.H{"error": "Invalid authorization token"})
					c.Abort()
					return
				}
			}
		} else {
			c.JSON(403, gin.H{"error": "An authorization header is required"})
			c.Abort()
			return
		}
	}
}
