package usecases

import (
	"time"

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

func (uc *cartUseCase) Save(cartID string) (*domain.Cart, error) {
	t := time.Now()
	cart, err := uc.repo.Save(cartID, &t)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (uc *cartUseCase) GetByID(cartID string) (*domain.Cart, error) {
	cart, err := uc.repo.GetByID(cartID)
	if err != nil {
		return nil, err
	}
	return cart, nil
}
