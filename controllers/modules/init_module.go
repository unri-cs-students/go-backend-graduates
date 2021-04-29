package modules

import (
	"github.com/SemmiDev/go-backend/controllers/middlewares"
	"github.com/SemmiDev/go-backend/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func Init(db *gorm.DB, g *gin.Engine, devmode bool) {
	NewCORSModule(g, devmode)
	NewLimiterModule(g)
	NewMajorModule(db, g)
	NewGraduatesModule(db, g)
	NewMessageModule(db, g)
	NewOrgzModule(db, g)
	NewContentModule(db, g)

	_ = db.AutoMigrate(&entities.View{})

	g.GET("/reset", middlewares.ResetAuth, func(c *gin.Context) {
		Reset(db, g)
	})

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Pong!")
	})
}

func Reset(db *gorm.DB, g *gin.Engine) {
	ResetMajor(db)
	ResetGraduates(db)
	ResetMessage(db)
	ResetOrgz(db)
	ResetContent(db)
	os.Exit(0)
}