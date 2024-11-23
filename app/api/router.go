package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(corsMiddleware())

	// handler := delivery.NewChatHandler(chatUseCase)

	// router.POST("/api/chat-add", handler.SaveChat)
	// router.POST("/api/chat-get", handler.GetChat)
	// router.PATCH("/api/msg-add", handler.SendMessage)
	// router.PATCH("/api/name-modify", handler.ChangeChatName)
	// router.PATCH("/api/member-add", handler.AddNewMember)
	// router.PATCH("/api/member-remove", handler.RemoveMember)

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
