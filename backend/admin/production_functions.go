package admin

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProduction(c *gin.Context, db *sql.DB) {
	var p Production

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}

	query := `INSERT INTO productions (id, title, description, genre, year)
			VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(query, p.ID, p.Title, p.Description, p.Genre, p.Year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})
}

func GetProductions(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, title, description, genre, year FROM productions WHERE id > 0")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var productions []Production
	for rows.Next() {
		var p Production
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Genre, &p.Year); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		productions = append(productions, p)
	}

	c.JSON(http.StatusOK, gin.H{"productions": productions})
}

func DeleteProduction(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	query := `DELETE FROM productions WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})
}

func GetProduction(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	var p Production
	query := `SELECT id, title, description, genre, year FROM productions WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&p.ID, &p.Title, &p.Description, &p.Genre, &p.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
		return
	}

	c.JSON(http.StatusOK, p)
}

func EditProduction(c *gin.Context, db *sql.DB) {
	var p Production

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
	query := `UPDATE productions SET title = $1, description = $2, genre = $3, year = $4 WHERE id = $5`
	_, err := db.Exec(query, p.Title, p.Description, p.Genre, p.Year, p.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})
}
