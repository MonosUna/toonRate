package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	Здесь хотим хранить функции работающие с обзорами
*/

type Review struct {
	ID         int64  `json:"id"`
	Production int64  `json:"production"`
	Author     string `json:"author"`
	Topic      string `json:"topic"`
	Message    string `json:"message"`
}

// компонента MainPage. Возращаем id случайного обзора
func RandomReview(c *gin.Context, db *sql.DB) {
	query := `
        SELECT id 
        FROM reviews 
        ORDER BY RANDOM() 
        LIMIT 1
    `

	row := db.QueryRow(query)

	var reviewID string

	err := row.Scan(&reviewID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": reviewID,
	})
}

// аналогична админке
func AddReview(c *gin.Context, db *sql.DB) {
	var r Review
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
	query := `INSERT INTO reviews (id, production_id, author, topic, message) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(query, r.ID, r.Production, r.Author, r.Topic, r.Message)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}

// аналогична админке
func GetReviews(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, production_id, author, topic, message FROM reviews")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var r Review
		if err := rows.Scan(&r.ID, &r.Production, &r.Author, &r.Topic, &r.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		reviews = append(reviews, r)
	}
	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

// аналогична админке
func GetReview(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var r Review
	query := `SELECT id, production_id, author, topic, message FROM reviews WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&r.ID, &r.Production, &r.Author, &r.Topic, &r.Message)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
		return
	}
	c.JSON(http.StatusOK, r)
}
