package controllers

import (
	_ "net/http"
	_ "time"

	"github.com/EltIsma/YandexLavka/cmd/internal/services"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}
func tollboothLimiter() gin.HandlerFunc {
	limiter := tollbooth.NewLimiter(2, nil)

	return func(c *gin.Context) {
	 httpError := tollbooth.LimitByRequest(limiter, c.Writer, c.Request)
	 if httpError != nil {
	  c.AbortWithStatusJSON(httpError.StatusCode, gin.H{"error": httpError.Message})
	  return
	 }
	 c.Next()
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(tollboothLimiter())

	courier := router.Group("/couriers")
	{
		courier.POST("/", h.createCouriers)
		courier.GET("/", h.getAllCouriers)
		courier.GET("/:id", h.getCourierById)
		courier.GET("/meta-info/:id", h.getCourierRatingSalary)
	}

	orders := router.Group("/orders")
	{
		orders.POST("/", h.createOrders)
		orders.GET("/", h.getAllOrders)
		orders.GET("/:id", h.getOrderById)
		orders.POST("/complete", h.ordersComplete)
	}

	return router
}
