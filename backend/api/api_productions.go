package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	Здесь хотим хранить функции работающие с произведениями
*/

type Production struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Year        int    `json:"year"`
}

// компонента: TopMainPage. Хотим вернуть топ-5 произведений по рейтингу
func Top5Productions(c *gin.Context, db *sql.DB) {
	query := `
        SELECT p.id, p.title, COALESCE(AVG(r.rating), 0) as average_rating
        FROM productions p
        LEFT JOIN ratings r ON r.entity_id = p.id AND r.type = 'production'
		WHERE p.id > 0
        GROUP BY p.id
        ORDER BY average_rating DESC
        LIMIT 5
    `

	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var productions []gin.H

	for rows.Next() {
		var id, title string
		var avgRating float64

		if err := rows.Scan(&id, &title, &avgRating); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}

		production := gin.H{
			"id":             id,
			"title":          title,
			"average_rating": avgRating,
		}

		productions = append(productions, production)
	}

	c.JSON(http.StatusOK, gin.H{"productions": productions})
}

// компонента MainPage. Возращаем id случайного произведения
func RandomProduction(c *gin.Context, db *sql.DB) {
	query := `
        SELECT id 
        FROM productions 
		WHERE id > 0
        ORDER BY RANDOM() 
        LIMIT 1
    `

	row := db.QueryRow(query)

	var productionID string

	err := row.Scan(&productionID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": productionID,
	})
}

type ProductionWithAverageRating struct {
	ID            int64   `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Genre         string  `json:"genre"`
	Year          int     `json:"year"`
	AverageRating float32 `json:"average_rating"`
}

// компонента TopProductions. Получение всех произведений в порядке возрастания рейтинга
func TopProductions(c *gin.Context, db *sql.DB) {
	var productions []ProductionWithAverageRating

	query := `
		SELECT p.id, p.title, p.description, p.genre, p.year, COALESCE(AVG(r.rating), 0) as average_rating
		FROM productions p
		LEFT JOIN ratings r ON r.entity_id = p.id AND r.type = 'production'
		WHERE p.id > 0
		GROUP BY p.id
		ORDER BY average_rating DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var production ProductionWithAverageRating

		if err := rows.Scan(&production.ID, &production.Title, &production.Description, &production.Genre, &production.Year, &production.AverageRating); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}

		productions = append(productions, production)
	}

	if err := rows.Err(); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"top_productions": productions,
	})
}

// аналогична админке
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

// аналогична админке
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
