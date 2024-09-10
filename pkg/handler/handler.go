package handler

import (
	_ "github.com/asstrahanec/go-clickhouse/docs"
	"github.com/asstrahanec/go-clickhouse/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		events := api.Group("/event")
		{
			events.POST("/", h.createEvent)
			events.GET("/", h.getEvents)

		}
	}
	return router
}
