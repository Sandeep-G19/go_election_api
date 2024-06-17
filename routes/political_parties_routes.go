package routes

import (
	"database/sql"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitPoliticalPartiesRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/political_parties", func(c *gin.Context) { controllers.GetPoliticalParties(c, db) })
	r.GET("/political_parties/:id", func(c *gin.Context) { controllers.GetPoliticalParty(c, db) })
	r.POST("/political_parties", func(c *gin.Context) { controllers.CreatePoliticalParty(c, db) })
	r.PUT("/political_parties/:id", func(c *gin.Context) { controllers.UpdatePoliticalParty(c, db) })
	r.DELETE("/political_parties/:id", func(c *gin.Context) { controllers.DeletePoliticalParty(c, db) })
}
