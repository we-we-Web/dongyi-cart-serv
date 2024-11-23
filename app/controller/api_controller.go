package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/we-we-Web/dongyi-cart-serv/app/usecases"
)

type CartController struct {
	cartUseCase usecases.CartUseCase
}

func NewCartController(cartUseCase usecases.CartUseCase) *CartController {
	return &CartController{cartUseCase}
}

func (h *CartController) SaveCart(c *gin.Context) {
	type Body struct {
		ID string `json:"id"` // owner id
	}
	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart, err := h.cartUseCase.Save(body.ID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func (h *CartController) GetCart(c *gin.Context) {
	type Body struct {
		ID string `json:"id"` // cart id
	}
	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart, err := h.cartUseCase.GetByID(body.ID)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}
