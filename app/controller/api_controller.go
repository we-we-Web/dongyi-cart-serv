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
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func (h *CartController) DeleteCart(c *gin.Context) {
	type Body struct {
		ID string `json:"id"` // cart id
	}
	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.cartUseCase.DeleteByID(body.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (h *CartController) UpdCartItem(c *gin.Context) {
	type Body struct {
		ID        string `json:"id"`      // cart id
		Product   string `json:"product"` // product id
		Delta     int    `json:"delta"`
		Remaining int    `json:"remaining"`
	}
	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart, err := h.cartUseCase.UpdProductItem(body.ID, body.Product, body.Delta, body.Remaining)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}
