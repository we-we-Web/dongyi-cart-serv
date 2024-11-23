package usecases

import (
	"github.com/we-we-Web/dongyi-cart-serv/app/domain"
	"github.com/we-we-Web/dongyi-cart-serv/app/repository"
)

type CartUseCase interface {
	Save(userID string) (*domain.Cart, error)
	GetByID(cartID string) (*domain.Cart, error)
}

type cartUseCase struct {
	repo repository.CartRepository
}

func NewCartUseCase(repo repository.CartRepository) CartUseCase {
	return &cartUseCase{repo}
}

func (uc *cartUseCase) Save(userID string) (*domain.Cart, error) {
	return nil, nil
}

func (uc *cartUseCase) GetByID(cartID string) (*domain.Cart, error) {
	return nil, nil
}
