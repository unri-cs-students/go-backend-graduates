package databases

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func PostgresConnect(debug bool) *gorm.DB {
	if dbConnection == nil {

		_ = godotenv.Load()
		host := os.Getenv("PG_HOST")
		port := os.Getenv("PG_PORT")
		dbname := os.Getenv("PG_DATABASE")
		user := os.Getenv("PG_USERNAME")
		password := os.Getenv("PG_PASSWORD")

		psqlLoginInfo := fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
			host, port, user, password, dbname)

		config := &gorm.Config{Logger: logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}

		if !debug {
			config = &gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
			}
		}

		dTemp, err := gorm.Open(postgres.Open(psqlLoginInfo), config)
		dbConnection = dTemp

		if err != nil {
			panic(err)
		}
	}
	return dbConnection
}