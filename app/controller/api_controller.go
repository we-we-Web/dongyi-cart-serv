package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/we-we-Web/dongyi-cart-serv/app/controller/dto"
	"github.com/we-we-Web/dongyi-cart-serv/app/usecases"
)

type CartController struct {
	cartUseCase usecases.CartUseCase
}

func NewCartController(cartUseCase usecases.CartUseCase) *CartController {
	return &CartController{cartUseCase}
}

func (h *CartController) SaveCart(c *gin.Context) {
	var body dto.AccessCartRequest
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
	var body dto.AccessCartRequest
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
	var body dto.AccessCartRequest
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
	var body dto.UpdProductRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart, err := h.cartUseCase.UpdProductItem(body.ID, body.Product, body.Size, body.Delta, body.Remaining)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func (h *CartController) ClearCart(c *gin.Context) {
	var body dto.AccessCartRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.cartUseCase.ClearCart(body.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "clear successfully"})
}

func (h *CartController) RemoveItem(c *gin.Context) {
	var body dto.RemoveItemRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.cartUseCase.RemoveItem(body.ID, body.Product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "remove item successfully"})
}
