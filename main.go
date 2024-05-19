package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-do/services"
	"go-do/system"
	"log"
	"os"
)

func main() {
	environment := initEnvironment()

	db := system.Database()
	services.InitServices(db)

	if environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	_ = engine.SetTrustedProxies(nil)
	system.Router(engine)
	err := engine.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func initEnvironment() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	environment := os.Getenv("ENV")
	if environment == "" {
		environment = "development"
		err := os.Setenv("ENV", environment)
		if err != nil {
			log.Println(err)
		}
	}
	if os.Getenv("PORT") == "" {
		err := os.Setenv("PORT", "8080")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("PORT: " + os.Getenv("PORT"))
	fmt.Println("ENVIRONMENT: " + environment)
	return environment
}
