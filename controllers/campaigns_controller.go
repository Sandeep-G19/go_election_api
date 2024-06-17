package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCampaigns(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, candidate_id, district_id, start_date, end_date, budget FROM campaigns")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	var campaigns []gin.H
	for rows.Next() {
		var id, candidateID, districtID int
		var startDate, endDate string
		var budget float64
		if err := rows.Scan(&id, &candidateID, &districtID, &startDate, &endDate, &budget); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
			return
		}
		campaigns = append(campaigns, gin.H{
			"id":           id,
			"candidate_id": candidateID,
			"district_id":  districtID,
			"start_date":   startDate,
			"end_date":     endDate,
			"budget":       budget,
		})
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error with rows"})
		return
	}

	log.Printf("Campaigns: %v", campaigns)
	c.JSON(http.StatusOK, campaigns)
}

func GetCampaign(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	row := db.QueryRow("SELECT id, candidate_id, district_id, start_date, end_date, budget FROM campaigns WHERE id = ?", id)

	var campaign gin.H
	var candidateID, districtID int
	var startDate, endDate string
	var budget float64
	if err := row.Scan(&id, &candidateID, &districtID, &startDate, &endDate, &budget); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No campaign found with id: %v", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "No campaign found"})
		} else {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
		}
		return
	}
	campaign = gin.H{
		"id":           id,
		"candidate_id": candidateID,
		"district_id":  districtID,
		"start_date":   startDate,
		"end_date":     endDate,
		"budget":       budget,
	}
	log.Printf("Campaign: %v", campaign)
	c.JSON(http.StatusOK, campaign)
}

func CreateCampaign(c *gin.Context, db *sql.DB) {
	var campaign struct {
		CandidateID int     `json:"candidate_id"`
		DistrictID  int     `json:"district_id"`
		StartDate   string  `json:"start_date"`
		EndDate     string  `json:"end_date"`
		Budget      float64 `json:"budget"`
	}

	if err := c.BindJSON(&campaign); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := db.Exec("INSERT INTO campaigns (candidate_id, district_id, start_date, end_date, budget) VALUES (?, ?, ?, ?, ?)",
		campaign.CandidateID, campaign.DistrictID, campaign.StartDate, campaign.EndDate, campaign.Budget)
	if err != nil {
		log.Printf("Error inserting data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
		return
	}

	id, _ := result.LastInsertId()
	campaignWithID := gin.H{
		"id":           id,
		"candidate_id": campaign.CandidateID,
		"district_id":  campaign.DistrictID,
		"start_date":   campaign.StartDate,
		"end_date":     campaign.EndDate,
		"budget":       campaign.Budget,
	}
	log.Printf("Campaign Created: %v", campaignWithID)
	c.JSON(http.StatusCreated, campaignWithID)
}

func UpdateCampaign(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var campaign struct {
		CandidateID int     `json:"candidate_id"`
		DistrictID  int     `json:"district_id"`
		StartDate   string  `json:"start_date"`
		EndDate     string  `json:"end_date"`
		Budget      float64 `json:"budget"`
	}

	if err := c.BindJSON(&campaign); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.Exec("UPDATE campaigns SET candidate_id = ?, district_id = ?, start_date = ?, end_date = ?, budget = ? WHERE id = ?",
		campaign.CandidateID, campaign.DistrictID, campaign.StartDate, campaign.EndDate, campaign.Budget, id)
	if err != nil {
		log.Printf("Error updating data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	campaignWithID := gin.H{
		"id":           id,
		"candidate_id": campaign.CandidateID,
		"district_id":  campaign.DistrictID,
		"start_date":   campaign.StartDate,
		"end_date":     campaign.EndDate,
		"budget":       campaign.Budget,
	}
	log.Printf("Campaign Updated: %v", campaignWithID)
	c.JSON(http.StatusOK, campaignWithID)
}

func DeleteCampaign(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM campaigns WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	log.Printf("Campaign Deleted: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted"})
}
