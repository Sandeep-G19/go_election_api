package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDistricts(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, name, description FROM districts")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
		return
	}
	defer rows.Close()

	var districts []gin.H
	for rows.Next() {
		var id int
		var name, description string
		if err := rows.Scan(&id, &name, &description); err != nil {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
			return
		}
		districts = append(districts, gin.H{
			"id":          id,
			"name":        name,
			"description": description,
		})
	}
	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error with rows"})
		return
	}

	log.Printf("Districts: %v", districts)
	c.JSON(http.StatusOK, districts)
}

func GetDistrict(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	row := db.QueryRow("SELECT id, name, description FROM districts WHERE id = ?", id)

	var district gin.H
	var name, description string
	if err := row.Scan(&id, &name, &description); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No district found with id: %v", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "No district found"})
		} else {
			log.Printf("Error scanning row: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
		}
		return
	}
	district = gin.H{
		"id":          id,
		"name":        name,
		"description": description,
	}
	log.Printf("District: %v", district)
	c.JSON(http.StatusOK, district)
}

func CreateDistrict(c *gin.Context, db *sql.DB) {
	var district struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.BindJSON(&district); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := db.Exec("INSERT INTO districts (name, description) VALUES (?, ?)", district.Name, district.Description)
	if err != nil {
		log.Printf("Error inserting data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting data"})
		return
	}

	id, _ := result.LastInsertId()
	districtWithID := gin.H{
		"id":          id,
		"name":        district.Name,
		"description": district.Description,
	}
	log.Printf("District Created: %v", districtWithID)
	c.JSON(http.StatusCreated, districtWithID)
}

func UpdateDistrict(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var district struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.BindJSON(&district); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.Exec("UPDATE districts SET name = ?, description = ? WHERE id = ?", district.Name, district.Description, id)
	if err != nil {
		log.Printf("Error updating data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}

	districtWithID := gin.H{
		"id":          id,
		"name":        district.Name,
		"description": district.Description,
	}
	log.Printf("District Updated: %v", districtWithID)
	c.JSON(http.StatusOK, districtWithID)
}

func DeleteDistrict(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	_, err := db.Exec("DELETE FROM districts WHERE id = ?", id)
	if err != nil {
		log.Printf("Error deleting data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	log.Printf("District Deleted: %v", id)
	c.JSON(http.StatusOK, gin.H{"message": "District deleted"})
}
