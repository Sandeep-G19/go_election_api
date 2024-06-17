package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVoters(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, name, address, dob FROM voters")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	var voters []gin.H
	for rows.Next() {
		var id int
		var name, address string
		var dob string
		if err := rows.Scan(&id, &name, &address, &dob); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
			return
		}
		voters = append(voters, gin.H{
			"id":      id,
			"name":    name,
			"address": address,
			"dob":     dob,
		})
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error with rows"})
		return
	}

	log.Printf("Voters: %v", voters)
	c.JSON(http.StatusOK, voters)
}

func GetVoter(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	row := db.QueryRow("SELECT id, name, address, dob FROM voters WHERE id = ?", id)

	var voter gin.H
	var name, address string
	var dob string
	if err := row.Scan(&id, &name, &address, &dob); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No voter found with id: %v", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "No voter found"})
		} else {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
		}
		return
	}
	voter = gin.H{
		"id":      id,
		"name":    name,
		"address": address,
		"dob":     dob,
	}
	log.Printf("Voter: %v", voter)
	c.JSON(http.StatusOK, voter)
}

func CreateVoter(c *gin.Context, db *sql.DB) {
	var voter struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		DOB     string `json:"dob"`
	}

	if err := c.BindJSON(&voter); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := db.Exec("INSERT INTO voters (name, address, dob) VALUES (?, ?, ?)", voter.Name, voter.Address, voter.DOB)
	if err != nil {
		log.Printf("Error inserting data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
		return
	}

	id, _ := result.LastInsertId()
	voterWithID := gin.H{
		"id":      id,
		"name":    voter.Name,
		"address": voter.Address,
		"dob":     voter.DOB,
	}
	log.Printf("Voter Created: %v", voterWithID)
	c.JSON(http.StatusCreated, voterWithID)
}

func UpdateVoter(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var voter struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		DOB     string `json:"dob"`
	}

	if err := c.BindJSON(&voter); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.Exec("UPDATE voters SET name = ?, address = ?, dob = ? WHERE id = ?", voter.Name, voter.Address, voter.DOB, id)
	if err != nil {
		log.Printf("Error updating data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	voterWithID := gin.H{
		"id":      id,
		"name":    voter.Name,
		"address": voter.Address,
		"dob":     voter.DOB,
	}
	log.Printf("Voter Updated: %v", voterWithID)
	c.JSON(http.StatusOK, voterWithID)
}

func DeleteVoter(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM voters WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	log.Printf("Voter Deleted: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "Voter deleted"})
}
