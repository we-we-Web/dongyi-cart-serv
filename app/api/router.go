package api

import (
	"github.com/gin-gonic/gin"
	"github.com/we-we-Web/dongyi-cart-serv/app/controller"
	"github.com/we-we-Web/dongyi-cart-serv/app/usecases"
)

func NewRouter(cartUseCase usecases.CartUseCase) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(corsMiddleware())

	handler := controller.NewCartController(cartUseCase)

	router.POST("/api/cart-create", handler.SaveCart)
	router.POST("/api/cart-get", handler.GetCart)
	router.DELETE("/api/cart-del", handler.DeleteCart)
	router.PATCH("/api/item-upd", handler.UpdCartItem)
	router.PATCH("/api/cart-clear", handler.ClearCart)

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
