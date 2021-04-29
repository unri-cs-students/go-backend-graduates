package main

import (
	"fmt"
	"github.com/SemmiDev/go-backend/configurations/connections/databases"
	"github.com/SemmiDev/go-backend/controllers/middlewares"
	"github.com/SemmiDev/go-backend/controllers/modules"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
	"strings"
)

func main() {
	_ = godotenv.Load()

	fmt.Println("Starting server...")

	development := true
	if strings.EqualFold(os.Getenv("GIN_MODE"), "release") {
		development = false
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	var db *gorm.DB

	if strings.EqualFold(os.Getenv("DBMS"), "mysql") {
		db = databases.MysqlConnect(development)
	} else {
		db = databases.PostgresConnect(development)
	}

	middlewares.InitErrorHandler(r)
	modules.Init(db, r, development)

	// Development Endpoint
	if development {
		modules.Development(db, r)
	}

	_ = r.Run(":9090")
}