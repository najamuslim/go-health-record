package main

import (
	"fmt"
	"health-record/app/server"
	"health-record/db"
	"health-record/helpers"
	"health-record/model/properties"
	"health-record/src/handler"
	"health-record/src/middleware"
	"health-record/src/repository"
	"health-record/src/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	r := server.InitServer()

	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Error loading .env file:", err)
	// 	return
	// }

	dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")
    dbHost := os.Getenv("DB_HOST")
    dbUsername := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbParams := os.Getenv("DB_PARAMS")

    // Construct the connection string
    connectionString := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?%s",
        dbUsername, dbPassword, dbHost, dbPort, dbName, dbParams,
    )

	fmt.Println("connectionString>> ", connectionString)
	fmt.Println("os.Getenv(DATABASE_URL)>> ", os.Getenv("DATABASE_URL"))
	postgreConfig := properties.PostgreConfig{
		DatabaseURL: connectionString,
	}

	db := db.InitPostgreDB(postgreConfig)
	//run migrations
	m, err := migrate.New(os.Getenv("MIGRATION_PATH"), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error creating migration instance: ", err)
	}

	//Run the migration up to the latest version
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error applying migrations:", err)
	}

	fmt.Println("Migration successfully applied")

	helper := helpers.NewHelper()

	// MIDDLEWARE
	middleware := middleware.NewMiddleware(helper)
	// REPOSITORY
	userRepository := repository.NewUserRepository(db)

	// USECASE
	authUsecase := usecase.NewAuthUsecase(userRepository, helper)

	// HANDLER
	authHandler := handler.NewAuthHandler(authUsecase)

	// ROUTE
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.POST("/v1/user/it/register", authHandler.Register)
	r.POST("/v1/user/it/login", authHandler.Login)

	authorized := r.Group("")
	authorized.Use(middleware.AuthMiddleware)

	r.Run()
}
