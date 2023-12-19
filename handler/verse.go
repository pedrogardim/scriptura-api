package handler

import (
	"encoding/json"
	"log"
	"scriptura/scriptura-api/db"
	"scriptura/scriptura-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetVerse(c *fiber.Ctx) error {
	db := db.DB
	var verses []models.Verse

	db.Table("verse").
		Select("*, bible.text, book.name as book_name").
		Joins("left join book on book.id = book_id").
		Joins("left join bible_en_kjv bible on bible.verse_id = verse.id").
		Where("book.id::varchar ilike $1 OR book.code ilike $1 OR book.short_name ilike $1", c.Params("book")).
		Limit(10).
		Scan(&verses)

	jsonData, err := json.Marshal(verses)

	if err != nil {
		log.Fatal(err)
	}

	return c.SendString(string(jsonData))
}
