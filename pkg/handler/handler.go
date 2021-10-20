// Слой работы с Http

package handler

import (
  "github.com/gin-gonic/gin"
  "github.com/piklin/test_ozon_fintech/pkg/service"
)

type Handler struct {
  services *service.Service
}

func NewHandler(services *service.Service) *Handler {
  return &Handler {
    services: services,
  }
}

func (h *Handler) InitRoutes() *gin.Engine {
  router := gin.New()

  router.GET("/:short_url", h.getFullURL)
  router.POST("/", h.createShortURL)

  return router
}
