package repository

import (
	"fmt"
	"time"

	"github.com/we-we-Web/dongyi-cart-serv/app/domain"
	"github.com/we-we-Web/dongyi-cart-serv/app/entity"
	"gorm.io/gorm"
)

type CartRepository interface {
	Save(cartID, userID string, t *time.Time) (*domain.Cart, error)
	GetByID(id string) (*domain.Cart, error)
	// UpdByID(field string, cart *domain.Cart) (*domain.Cart, error)
	DeleteByID(cartID string) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) Save(cartID, userID string, t *time.Time) (*domain.Cart, error) {
	cart := domain.NewCart(cartID, userID, t)
	if err := r.db.Create(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (r *cartRepository) GetByID(cartID string) (*domain.Cart, error) {
	var cart *domain.Cart
	if err := r.db.Where("id = ?", cartID).Order("id").First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

// func (r *cartRepository) UpdByID(field string, user *domain.Cart) (*domain.Cart, error) {
// 	userEntity, err := parseToEntity(user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	v := reflect.ValueOf(userEntity).Elem()
// 	f := v.FieldByName(field)
// 	if !f.IsValid() {
// 		return nil, errors.New("specified field does not exist in user entity")
// 	}

// 	if err := r.db.Model(userEntity).Update(field, f.Interface()).Error; err != nil {
// 		return nil, err
// 	}
// 	return r.GetByID(user.ID)
// }

func (r *cartRepository) DeleteByID(cartID string) error {
	result := r.db.Where("id = ?", cartID).Delete(&entity.Cart{})
	if result.Error != nil {
		return fmt.Errorf("error occur when deleting the cart: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("cart %s was not found", cartID)
	}

	return nil
}
