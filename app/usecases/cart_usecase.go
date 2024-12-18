package usecases

import (
	"fmt"
	"time"

	"github.com/we-we-Web/dongyi-cart-serv/app/domain"
	"github.com/we-we-Web/dongyi-cart-serv/app/repository"
)

type CartUseCase interface {
	Save(userID string) (*domain.Cart, error)
	GetByID(cartID string) (*domain.Cart, error)
	DeleteByID(cartID string) error
	UpdProductItem(cartID, productID, size string, delta int, remaining int) (*domain.Cart, error)
	ClearCart(cartID string) error
	RemoveItem(cartID, productID string) error
}

type cartUseCase struct {
	repo repository.CartRepository
}

func NewCartUseCase(repo repository.CartRepository) CartUseCase {
	return &cartUseCase{repo}
}

func (uc *cartUseCase) Save(cartID string) (*domain.Cart, error) {
	t := time.Now()
	cart, err := uc.repo.Save(cartID, t)
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
	for _, item := range cart.Products {
		if item.Spec == nil {
			cart.Products = []domain.CartItem{}
			field := "Products"
			if _, err := uc.repo.UpdByID(field, cart); err != nil {
				return nil, err
			}
		}
	}
	return cart, nil
}

func (uc *cartUseCase) DeleteByID(cartID string) error {
	return uc.repo.DeleteByID(cartID)
}

func (uc *cartUseCase) UpdProductItem(cartID, productID, size string, delta int, remaining int) (*domain.Cart, error) {
	cart, err := uc.repo.GetByID(cartID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch cart %s: %w", cartID, err)
	}

	if cart == nil {
		return nil, fmt.Errorf("cart %s not found", cartID)
	}

	found := false
	for i, item := range cart.Products {

		if item.Product == productID {
			if item.Spec[size]+delta > remaining {
				return nil, fmt.Errorf("quantity exceeds remaining stock")
			}
			cart.Products[i].Spec[size] += delta
			found = true
			if cart.Products[i].Spec[size] <= 0 {
				cart.Products = removeProductItem(cart.Products, i)
			}
			break
		}
	}

	if !found && delta > 0 && delta <= remaining {
		cart.Products = appendNewProductItem(cart.Products, productID, size, delta)
	}

	field := "Products"
	cart, err = uc.repo.UpdByID(field, cart)
	if err != nil {
		return nil, fmt.Errorf("failed to update cart item in cart %s: %w", cart.ID, err)
	}

	return cart, nil
}

func (uc *cartUseCase) ClearCart(cartID string) error {
	cart, err := uc.repo.GetByID(cartID)
	if err != nil {
		return err
	}
	cart.Products = []domain.CartItem{}
	field := "Products"
	_, err = uc.repo.UpdByID(field, cart)
	if err != nil {
		return err
	}
	return nil
}

func (uc *cartUseCase) RemoveItem(cartID, productID string) error {
	cart, err := uc.GetByID(cartID)
	if err != nil {
		return err
	}
	pos := -1
	for idx, item := range cart.Products {
		if item.Product == productID {
			pos = idx
			break
		}
	}
	if pos != -1 {
		cart.Products = removeProductItem(cart.Products, pos)
	}
	field := "Products"
	_, err = uc.repo.UpdByID(field, cart)
	if err != nil {
		return err
	}
	return nil
}

func appendNewProductItem(products []domain.CartItem, productID, size string, quantity int) []domain.CartItem {
	newProducts := append(products, domain.CartItem{
		Product: productID,
		Spec: map[string]int{
			size: quantity,
		},
	})
	return newProducts
}

func removeProductItem(products []domain.CartItem, idx int) []domain.CartItem {
	newProducts := append(products[:idx], products[idx+1:]...)
	return newProducts
}
