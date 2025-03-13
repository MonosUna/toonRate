package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	Здесь хотим хранить функции взаимодействия с пользователем
*/

type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Pfp         string `json:"pfp"`
	Description string `json:"description"`
}

// компонента Profile. Количество сущностей созданных пользователем
func Statistics(c *gin.Context, db *sql.DB) {
	username := c.Param("username")

	query := `
        SELECT 
            COALESCE(SUM(CASE WHEN t.type = 'rating' AND t.entity_type = 'production' THEN 1 ELSE 0 END), 0) AS ratings_count,
            COALESCE(SUM(CASE WHEN t.type = 'discussion' THEN 1 ELSE 0 END), 0) AS discussions_count,
            COALESCE(SUM(CASE WHEN t.type = 'review' THEN 1 ELSE 0 END), 0) AS reviews_count,
            COALESCE(SUM(CASE WHEN t.type = 'collection' THEN 1 ELSE 0 END), 0) AS collections_count
        FROM (
            SELECT 'rating' AS type, 'production' AS entity_type
            FROM ratings
            WHERE author = $1
            UNION ALL
            SELECT 'discussion', NULL
            FROM discussions
            WHERE author = $1
            UNION ALL
            SELECT 'review', NULL
            FROM reviews
            WHERE author = $1
            UNION ALL
            SELECT 'collection', NULL
            FROM collections
            WHERE author = $1
        ) t
    `

	var ratingsCount, discussionsCount, reviewsCount, collectionsCount int
	err := db.QueryRow(query, username).Scan(&ratingsCount, &discussionsCount, &reviewsCount, &collectionsCount)
	if err != nil {
		log.Println("Error fetching statistics:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ratings_count":     ratingsCount,
		"discussions_count": discussionsCount,
		"reviews_count":     reviewsCount,
		"collections_count": collectionsCount,
	})
}

type NewUserDescription struct {
	Username    string `json:"username"`
	Description string `json:"description"`
}

// компонента Profile. Изменение описания профиля
func UpdateUserDescription(c *gin.Context, db *sql.DB) {
	var requestData struct {
		ID          int64  `json:"ID"`
		Description string `json:"Description"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}

	query := `UPDATE users SET description = $1 WHERE id = $2`
	result, err := db.Exec(query, requestData.Description, requestData.ID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})
}

// компонента Profile. Изменение картинки профиля
func UpdateUserPfp(c *gin.Context, db *sql.DB) {
	var requestData struct {
		ID  int64  `json:"ID"`
		Pfp string `json:"Pfp"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}

	if requestData.Pfp == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}

	query := `UPDATE users SET pfp = $1 WHERE id = $2`
	result, err := db.Exec(query, requestData.Pfp, requestData.ID)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})
}

// аналогична админке
func AddUser(c *gin.Context, db *sql.DB) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
	query := `INSERT INTO users (id, username, password, email, pfp, description) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(query, u.ID, u.Username, u.Password, u.Email, u.Pfp, u.Description)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}

// аналогична админке
func GetUsers(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, username, password, email, pfp, description FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.Pfp, &u.Description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		users = append(users, u)
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
