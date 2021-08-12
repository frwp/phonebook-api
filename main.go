package main

import (
	"fmt"
	"log"

	"github.com/RianWardanaPutra/phonebook-api/controllers"
	_ "github.com/RianWardanaPutra/phonebook-api/docs"
	"github.com/RianWardanaPutra/phonebook-api/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	dsn := "host='redacted' user='redacted' dbname='redacted' port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db.AutoMigrate(models.Contact{})

	return db, nil
}

// @title Phonebook API
// @version 1.0
// @description This is a phonebook server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth
func main() {
	fmt.Println("Hello")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:3000"}
	// r.Use(cors.New(config))

	r.Use(cors.Default())

	db, err := ConnectDb()

	if err != nil {
		log.Fatal(err)
		return
	}

	c := controllers.NewController(db)

	v1 := r.Group("/api/v1")
	{
		contacts := v1.Group("/contacts")
		{
			contacts.GET("", c.ListContacts)
			contacts.POST("", c.AddContact)
			contacts.GET("/:id", c.FindContactById)
			contacts.PUT("/:id", c.UpdateContactById)
			contacts.DELETE("/:id", c.DeleteContactById)
		}
	}
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":8080")
}
