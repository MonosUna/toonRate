package api

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
	Здесь хотим хранить функции работающие с подборками
*/

type Collection struct {
	ID      int64  `json:"id"`
	Author  string `json:"author"`
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

type ProductionShort struct {
	Title   string
	Genre   string
	Comment string
	Year    int
}

// компонента Collections. Вернуть все произведения коллекции
func GetCollectionProductions(c *gin.Context, db *sql.DB) {
	collectionID := c.Param("id")
	id, _ := strconv.ParseInt(collectionID, 10, 64)
	rows, err := db.Query(`
		SELECT p.title, p.genre, p.year, cp.comment
		FROM collectionsProductions cp
		JOIN productions p ON cp.production_id = p.id
		WHERE cp.collection_id = $1`, id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var productions []ProductionShort
	for rows.Next() {
		var production ProductionShort
		if err := rows.Scan(&production.Title, &production.Genre, &production.Year, &production.Comment); err != nil {
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

	c.JSON(http.StatusOK, gin.H{"productions": productions})
}

type CollectionProduction struct {
	ID           int64  `json:"id"`
	CollectionID int64  `json:"collection_id"`
	ProductionID int64  `json:"production_id"`
	Comment      string `json:"comment"`
}

// компонента Collection. Добавляем в базу данных произведение из коллекции.
func AddCollectionProduction(c *gin.Context, db *sql.DB) {
	var item CollectionProduction

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}

	query := `INSERT INTO collectionsProductions (id, collection_id, production_id, comment) 
	          VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, item.ID, item.CollectionID, item.ProductionID, item.Comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": ""})
}

// компонента CreateCollection. Получаем все произведения из всех коллекций.
func GetCollectionsProductions(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, collection_id, production_id, comment FROM collectionsProductions")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var collectionsProductions []CollectionProduction
	for rows.Next() {
		var col CollectionProduction
		if err := rows.Scan(&col.ID, &col.CollectionID, &col.ProductionID, &col.Comment); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		collectionsProductions = append(collectionsProductions, col)
	}
	c.JSON(http.StatusOK, gin.H{"collectionsProductions": collectionsProductions})
}

// компонента MainPage. Возращаем id случайной подборкиs
func RandomCollection(c *gin.Context, db *sql.DB) {
	query := `
        SELECT id 
        FROM collections 
        ORDER BY RANDOM() 
        LIMIT 1
    `

	row := db.QueryRow(query)

	var collectionID string

	err := row.Scan(&collectionID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": collectionID,
	})
}

// аналогична админке
func AddCollection(c *gin.Context, db *sql.DB) {
	var col Collection
	if err := c.ShouldBindJSON(&col); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
	query := `INSERT INTO collections (id, author, topic, message) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, col.ID, col.Author, col.Topic, col.Message)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": ""})
}

// аналогична админке
func GetCollections(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT id, author, topic, message FROM collections")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer rows.Close()

	var collections []Collection
	for rows.Next() {
		var col Collection
		if err := rows.Scan(&col.ID, &col.Author, &col.Topic, &col.Message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}
		collections = append(collections, col)
	}
	c.JSON(http.StatusOK, gin.H{"collections": collections})
}

// аналогична админке
func GetCollection(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var col Collection
	query := `SELECT id, author, topic, message FROM collections WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&col.ID, &col.Author, &col.Topic, &col.Message)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": ""})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
		return
	}
	c.JSON(http.StatusOK, col)
}
