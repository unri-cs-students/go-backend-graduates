package databases

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
func MysqlConnect(debug bool) *gorm.DB {
	if dbConnection == nil {

		_ = godotenv.Load()
		host := os.Getenv("MYSQL_HOST")
		port := os.Getenv("MYSQL_PORT")
		dbname := os.Getenv("MYSQL_DATABASE")
		user := os.Getenv("MYSQL_USERNAME")
		password := os.Getenv("MYSQL_PASSWORD")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
		config := &gorm.Config{Logger: logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}

		if !debug {
			config = &gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
			}
		}

		db, err := gorm.Open(mysql.Open(dsn), config)
		if err != nil {
			panic(err)
		}

		dbConnection = db
	}
	return dbConnection
}