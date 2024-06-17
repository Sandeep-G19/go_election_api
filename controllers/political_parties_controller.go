package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPoliticalParties(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, name, leader FROM political_parties")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	var parties []gin.H
	for rows.Next() {
		var id int
		var name, leader string
		if err := rows.Scan(&id, &name, &leader); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
			return
		}
		parties = append(parties, gin.H{
			"id":     id,
			"name":   name,
			"leader": leader,
		})
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error with rows"})
		return
	}

	log.Printf("Political Parties: %v", parties)
	c.JSON(http.StatusOK, parties)
}

func GetPoliticalParty(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	row := db.QueryRow("SELECT id, name, leader FROM political_parties WHERE id = ?", id)

	var party gin.H
	var name, leader string
	if err := row.Scan(&id, &name, &leader); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No political party found with id: %v", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "No political party found"})
		} else {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
		}
		return
	}
	party = gin.H{
		"id":     id,
		"name":   name,
		"leader": leader,
	}
	log.Printf("Political Party: %v", party)
	c.JSON(http.StatusOK, party)
}

func CreatePoliticalParty(c *gin.Context, db *sql.DB) {
	var party struct {
		Name   string `json:"name"`
		Leader string `json:"leader"`
	}

	if err := c.BindJSON(&party); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := db.Exec("INSERT INTO political_parties (name, leader) VALUES (?, ?)", party.Name, party.Leader)
	if err != nil {
		log.Printf("Error inserting data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
		return
	}

	id, _ := result.LastInsertId()
	partyWithID := gin.H{
		"id":     id,
		"name":   party.Name,
		"leader": party.Leader,
	}
	log.Printf("Political Party Created: %v", partyWithID)
	c.JSON(http.StatusCreated, partyWithID)
}

func UpdatePoliticalParty(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var party struct {
		Name   string `json:"name"`
		Leader string `json:"leader"`
	}

	if err := c.BindJSON(&party); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.Exec("UPDATE political_parties SET name = ?, leader = ? WHERE id = ?", party.Name, party.Leader, id)
	if err != nil {
		log.Printf("Error updating data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	partyWithID := gin.H{
		"id":     id,
		"name":   party.Name,
		"leader": party.Leader,
	}
	log.Printf("Political Party Updated: %v", partyWithID)
	c.JSON(http.StatusOK, partyWithID)
}

func DeletePoliticalParty(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM political_parties WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	log.Printf("Political Party Deleted: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "Political party deleted"})
}
