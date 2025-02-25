package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/telman03/go-url-service/cache"
	"github.com/telman03/go-url-service/database"
	"github.com/telman03/go-url-service/models"
	"github.com/telman03/go-url-service/utils"
)

const cacheExpiration = 24 * time.Hour

func ShortenURL(c *fiber.Ctx) error {
	type Request struct {
		URL string `json:"url"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	shortCode := utils.GenerateShortCode(6)

	url := models.URL{
		OriginalURL: body.URL,
		ShortCode:   shortCode,
	}


	if err := database.DB.Create(&url).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save URL"})
	}


	err := cache.RedisClient.Set(context.Background(), shortCode, body.URL, 24*time.Hour).Err()
	if err != nil {
		log.Println("⚠️ Failed to cache URL in Redis:", err)
	}

	return c.JSON(fiber.Map{"short_url": "http://localhost:8080/" + shortCode})
}



func RedirectURL(c *fiber.Ctx) error {
	shortCode := c.Params("shortcode")

	var url models.URL
	if err := database.DB.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "URL not found"})
	}


	ctx := context.Background()
	redisKey := "clicks:" + shortCode

	err := cache.RedisClient.Incr(ctx, redisKey).Err()
	if err != nil {
		log.Println("❌ Redis Increment Error:", err)
	}

	
	clicks, err := cache.RedisClient.Get(ctx, redisKey).Result()
	if err != nil {
		log.Println("❌ Redis Get Error:", err)
	}

	log.Println("✅ Click Count Updated for", shortCode, "->", clicks)

	return c.Redirect(url.OriginalURL, fiber.StatusMovedPermanently)
}