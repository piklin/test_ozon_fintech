package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piklin/test_ozon_fintech/models"
)

func (h *Handler) createShortURL(c *gin.Context) {
	var input models.URLRequest

	if error := c.BindJSON(&input); error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid input")
		return
	}

	shortURL, error := h.services.ShortURL.Create(input)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, models.ShortURLResponce{
		ShortURL: shortURL,
	})
}

func (h *Handler) getFullURL(c *gin.Context) {
	shortURL := c.Param("short_url")
	if shortURL == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid input")
		return
	}

	fullURL, error := h.services.ShortURL.GetFullURL(shortURL)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal server error")
		return
	} else if fullURL == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "URL does not exist")
		return
	}

	c.JSON(http.StatusOK, models.FullURLResponce{
		FullURL: fullURL,
	})
}
