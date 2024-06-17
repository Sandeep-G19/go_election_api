package routes

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, db *sql.DB) {
	log.Println("Initializing Election Routes")
	InitElectionRoutes(r, db)

	log.Println("Initializing Political Parties Routes")
	InitPoliticalPartiesRoutes(r, db)

	log.Println("Initializing Campaigns Routes")
	InitCampaignsRoutes(r, db)

	log.Println("Initializing Candidates Routes")
	InitCandidatesRoutes(r, db)

	log.Println("Initializing Districts Routes")
	InitDistrictsRoutes(r, db)

	log.Println("Initializing Results Routes")
	InitResultsRoutes(r, db)

	log.Println("Initializing Voters Routes")
	InitVotersRoutes(r, db)
}
