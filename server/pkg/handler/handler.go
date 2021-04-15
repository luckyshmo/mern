package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/luckyshmo/api-example/pkg/service"

	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity) //JWT Auth
	{
		users := api.Group("/user")
		{
			users.GET("/", h.getUserList)
			users.GET("/:id", h.getUser)
		}

		keep := api.Group("/keep")
		{
			keep.GET("/", h.keep)
		}
	}

	return router
}
