package handlers
import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/telman03/go-url-service/cache"



)
func GetAnalytics(c *fiber.Ctx) error {
    shortCode := c.Params("shortcode")

    ctx := context.Background()
    redisKey := "clicks:" + shortCode

    clicks, err := cache.RedisClient.Get(ctx, redisKey).Result()
    if err != nil {
        clicks = "0" /
    }

    return c.JSON(fiber.Map{"shortcode": shortCode, "clicks": clicks})
}