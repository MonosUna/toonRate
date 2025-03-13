package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	Здесь хотим хранить функции работающие с оценками и статусами произведений
*/

// общая функция: получить все оценки произведения
func GetProductionRatings(c *gin.Context, db *sql.DB) {
	id := c.Param("id")

	query := `SELECT rating FROM ratings WHERE type = 'production' AND entity_id = $1`
	rows, err := db.Query(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var ratings []int

	for rows.Next() {
		var rating int
		if err := rows.Scan(&rating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		ratings = append(ratings, rating)
	}

	c.JSON(http.StatusOK, gin.H{
		"ratings": ratings,
	})
}

// компонента ProductionStatus. Получаем информацию о произведениях из листов
type ProductionFromProductionStatus struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

// компонента ProductionStatus. Получение списков статусов произведений
func AllProductionStatus(c *gin.Context, db *sql.DB) {
	username := c.Param("username")

	var viewed []ProductionFromProductionStatus
	var watching []ProductionFromProductionStatus
	var toWatch []ProductionFromProductionStatus

	rows, err := db.Query(`
        SELECT r.entity_id, p.title, p.year, r.rating
        FROM ratings r
        JOIN productions p ON r.entity_id = p.id
        WHERE r.author = $1 AND r.type = 'PS'`, username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var entityID, title string
		var year, rating int
		if err := rows.Scan(&entityID, &title, &year, &rating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}

		production := ProductionFromProductionStatus{
			ID:    entityID,
			Title: title,
			Year:  year,
		}

		switch rating {
		case 1:
			viewed = append(viewed, production)
		case 2:
			watching = append(watching, production)
		case 3:
			toWatch = append(toWatch, production)
		}
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"viewed":   viewed,
		"watching": watching,
		"to_watch": toWatch,
	})
}

type NewRatingOrStatus struct {
	Author string `json:"author"`
	ID     int64  `json:"ID"`
	Rating int32  `json:"rating"`
}

// компонента Product. Оценка произведения.
func SetNewProductionRating(c *gin.Context, db *sql.DB) {
	var newRating NewRatingOrStatus

	if err := c.ShouldBindJSON(&newRating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}

	if newRating.Rating < 1 || newRating.Rating > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}

	var existingRatingID int64
	query := `
		SELECT id FROM ratings
		WHERE entity_id = $1 AND author = $2 AND type = 'production'
	`

	err := db.QueryRow(query, newRating.ID, newRating.Author).Scan(&existingRatingID)

	if err == sql.ErrNoRows {
		var maxID int
		err = db.QueryRow("SELECT MAX(id) FROM ratings").Scan(&maxID)
		if err != nil {
			maxID = 0
		}

		insertQuery := `
			INSERT INTO ratings (id, type, rating, entity_id, author)
			VALUES ($1, 'production', $2, $3, $4)
		`
		_, err = db.Exec(insertQuery, maxID+1, newRating.Rating, newRating.ID, newRating.Author)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": ""})
		return
	} else if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	updateQuery := `
		UPDATE ratings 
		SET rating = $1
		WHERE id = $2 AND type = 'production'
	`
	_, err = db.Exec(updateQuery, newRating.Rating, existingRatingID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})
}

// компонента Product. Выставляем новый статус.
func SetNewProductionStatus(c *gin.Context, db *sql.DB) {
	var newRating NewRatingOrStatus

	if err := c.ShouldBindJSON(&newRating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}

	var existingRatingID int64
	query := `
		SELECT id FROM ratings
		WHERE entity_id = $1 AND author = $2 AND type = 'PS'
	`
	err := db.QueryRow(query, newRating.ID, newRating.Author).Scan(&existingRatingID)

	if err == sql.ErrNoRows {
		var maxID int
		err = db.QueryRow("SELECT MAX(id) FROM ratings").Scan(&maxID)
		if err != nil {
			log.Println(err.Error())
			maxID = 0
		}
		insertQuery := `
			INSERT INTO ratings (id, type, rating, entity_id, author)
			VALUES ($1, 'PS', $2, $3, $4)
		`
		_, err = db.Exec(insertQuery, maxID+1, newRating.Rating, newRating.ID, newRating.Author)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": ""})
		return
	} else if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	updateQuery := `
		UPDATE ratings 
		SET rating = $1
		WHERE id = $2 AND type = 'PS'
	`
	_, err = db.Exec(updateQuery, newRating.Rating, existingRatingID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})
}

// компонента Product. Получаем текущий статус
func GetProductionStatus(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	username := c.Param("username")

	query := `SELECT rating FROM ratings WHERE type = 'PS' AND entity_id = $1 AND author = $2`
	rows, err := db.Query(query, id, username)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var status int

	if rows.Next() {
		var rating int
		if err := rows.Scan(&rating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		status = rating
	} else {
		status = 0
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
