package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/we-we-Web/dongyi-cart-serv/app/domain"
	"github.com/we-we-Web/dongyi-cart-serv/app/repository"
	"github.com/we-we-Web/dongyi-cart-serv/app/usecases"
)

func TestCartUseCase_Save(t *testing.T) {
	mockRepo := new(repository.MockCartRepository)
	uc := usecases.NewCartUseCase(mockRepo)

	cartID := "test_cart"
	expectedCart := &domain.Cart{ID: cartID}
	mockRepo.On("Save", cartID, mock.AnythingOfType("time.Time")).Return(expectedCart, nil)

	cart, err := uc.Save(cartID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCart, cart)
	mockRepo.AssertCalled(t, "Save", cartID, mock.AnythingOfType("time.Time"))
}

func TestCartUseCase_GetByID(t *testing.T) {
	mockRepo := new(repository.MockCartRepository)
	uc := usecases.NewCartUseCase(mockRepo)

	cartID := "test_cart"
	expectedCart := &domain.Cart{ID: cartID}
	mockRepo.On("GetByID", cartID).Return(expectedCart, nil)

	cart, err := uc.GetByID(cartID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCart, cart)
	mockRepo.AssertCalled(t, "GetByID", cartID)
}

func TestCartUseCase_DeleteByID(t *testing.T) {
	mockRepo := new(repository.MockCartRepository)
	uc := usecases.NewCartUseCase(mockRepo)

	cartID := "test_cart"
	mockRepo.On("DeleteByID", cartID).Return(nil)

	err := uc.DeleteByID(cartID)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "DeleteByID", cartID)
}

func TestCartUseCase_UpdProductItem_AddNewItem(t *testing.T) {
	mockRepo := new(repository.MockCartRepository)
	uc := usecases.NewCartUseCase(mockRepo)

	cartID := "test_cart"
	productID := "product_1"
	quantity := 3
	initialCart := &domain.Cart{ID: cartID, Products: []domain.CartItem{}}
	updatedCart := &domain.Cart{ID: cartID, Products: []domain.CartItem{{Product: productID, Quantity: quantity}}}

	mockRepo.On("GetByID", cartID).Return(initialCart, nil)
	mockRepo.On("UpdByID", "Products", updatedCart).Return(updatedCart, nil)

	cart, err := uc.UpdProductItem(cartID, productID, quantity)

	assert.NoError(t, err)
	assert.Equal(t, updatedCart, cart)
	mockRepo.AssertCalled(t, "GetByID", cartID)
	mockRepo.AssertCalled(t, "UpdByID", "Products", updatedCart)
}

func TestCartUseCase_UpdProductItem_UpdateQuantity(t *testing.T) {
	mockRepo := new(repository.MockCartRepository)
	uc := usecases.NewCartUseCase(mockRepo)

	cartID := "test_cart"
	productID := "product_1"
	newQuantity := 5
	initialCart := &domain.Cart{ID: cartID, Products: []domain.CartItem{{Product: productID, Quantity: 2}}}
	updatedCart := &domain.Cart{ID: cartID, Products: []domain.CartItem{{Product: productID, Quantity: newQuantity}}}

	mockRepo.On("GetByID", cartID).Return(initialCart, nil)
	mockRepo.On("UpdByID", "Products", updatedCart).Return(updatedCart, nil)

	cart, err := uc.UpdProductItem(cartID, productID, newQuantity)

	assert.NoError(t, err)
	assert.Equal(t, updatedCart, cart)
	mockRepo.AssertCalled(t, "GetByID", cartID)
	mockRepo.AssertCalled(t, "UpdByID", "Products", updatedCart)
}

func TestCartUseCase_UpdProductItem_RemoveItem(t *testing.T) {
	mockRepo := new(repository.MockCartRepository)
	uc := usecases.NewCartUseCase(mockRepo)

	cartID := "test_cart"
	productID := "product_1"
	initialCart := &domain.Cart{ID: cartID, Products: []domain.CartItem{{Product: productID, Quantity: 2}}}
	updatedCart := &domain.Cart{ID: cartID, Products: []domain.CartItem{}}

	mockRepo.On("GetByID", cartID).Return(initialCart, nil)
	mockRepo.On("UpdByID", "Products", updatedCart).Return(updatedCart, nil)

	cart, err := uc.UpdProductItem(cartID, productID, 0)

	assert.NoError(t, err)
	assert.Equal(t, updatedCart, cart)
	mockRepo.AssertCalled(t, "GetByID", cartID)
	mockRepo.AssertCalled(t, "UpdByID", "Products", updatedCart)
}
