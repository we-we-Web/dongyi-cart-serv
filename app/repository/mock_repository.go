package repository

import (
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/we-we-Web/dongyi-cart-serv/app/domain"
)

type MockCartRepository struct {
	mock.Mock
}

func (m *MockCartRepository) Save(cartID string, t time.Time) (*domain.Cart, error) {
	args := m.Called(cartID, t)
	return args.Get(0).(*domain.Cart), args.Error(1)
}

func (m *MockCartRepository) GetByID(cartID string) (*domain.Cart, error) {
	args := m.Called(cartID)
	return args.Get(0).(*domain.Cart), args.Error(1)
}

func (m *MockCartRepository) DeleteByID(cartID string) error {
	args := m.Called(cartID)
	return args.Error(0)
}

func (m *MockCartRepository) UpdByID(field string, cart *domain.Cart) (*domain.Cart, error) {
	args := m.Called(field, cart)
	return args.Get(0).(*domain.Cart), args.Error(1)
}
