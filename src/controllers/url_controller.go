package controllers

import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/hex"
	"log"
	"os"
	"urlshorter/conf"
	"urlshorter/src/models"
	"urlshorter/src/res"
	"urlshorter/src/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

type UrlController struct{}

func NewUrlController() *UrlController {
	return &UrlController{}
}

func (u *UrlController) IndexShortenUrl(c *fiber.Ctx) error {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
    shortUrl := c.Params("shortUrl")

    if shortUrl == "" {
        // Operasi untuk membuat URL pendek baru

        var req models.Url

        if err := c.BodyParser(&req); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(res.ResponseCode{
                Status:  fiber.StatusBadRequest,
                Message: err.Error(),
                Data:    nil,
            })
        }

        if err := utils.Validate(req); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(res.ResponseCode{
                Status:  fiber.StatusBadRequest,
                Message: err.Error(),
                Data:    nil,
            })
        }

        encode := base32.StdEncoding.EncodeToString([]byte(req.OriginalUrl))
        hash := sha256.Sum256([]byte(encode))
        hashsort := hex.EncodeToString(hash[:])[:6]

        // Assign hash pendek ke request model
        req.ShortUrl = hashsort
        err := confiq.Database().Transaction(func(tx *gorm.DB) error {
            if err := tx.Create(&req).Error; err != nil {
                return err
            }
            return nil
        })

        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(res.ResponseCode{
                Status:  fiber.StatusInternalServerError,
                Message: err.Error(),
                Data:    nil,
            })
        }

		baseURL := os.Getenv("BASE_URL")
		shortUrlWithDomain := baseURL + "/" + hashsort
        return c.JSON(res.ResponseCode{
            Status:  fiber.StatusOK,
            Message: "Success",
            Data:    shortUrlWithDomain,
        })
    } else {
        // Operasi untuk redirect ke URL asli

        var url models.Url
        if err := confiq.Database().Where("short_url = ?", shortUrl).First(&url).Error; err != nil {
            return c.Status(fiber.StatusNotFound).JSON(res.ResponseCode{
                Status:  fiber.StatusNotFound,
                Message: "URL not found",
                Data:    nil,
            })
        }

        return c.Redirect(url.OriginalUrl, fiber.StatusFound)
    }
}

