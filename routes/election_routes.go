package routes

import (
	"database/sql"
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func InitElectionRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("/elections", func(c *gin.Context) { controllers.GetElections(c, db) })
	r.GET("/elections/:id", func(c *gin.Context) { controllers.GetElection(c, db) })
	r.POST("/elections", func(c *gin.Context) { controllers.CreateElection(c, db) })
	r.PUT("/elections/:id", func(c *gin.Context) { controllers.UpdateElection(c, db) })
	r.DELETE("/elections/:id", func(c *gin.Context) { controllers.DeleteElection(c, db) })
}
