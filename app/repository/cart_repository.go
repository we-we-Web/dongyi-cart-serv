package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/we-we-Web/dongyi-cart-serv/app/domain"
	"github.com/we-we-Web/dongyi-cart-serv/app/entity"
	"gorm.io/gorm"
)

type CartRepository interface {
	Save(cartID string, t *time.Time) (*domain.Cart, error)
	GetByID(cartID string) (*domain.Cart, error)
	// UpdByID(field string, cart *domain.Cart) (*domain.Cart, error)
	DeleteByID(cartID string) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) Save(cartID string, t *time.Time) (*domain.Cart, error) {
	cartModel := domain.NewCart(cartID, t)
	cartEntity, err := parseToEntity(cartModel)
	if err != nil {
		return nil, err
	}
	if err := r.db.Create(cartEntity).Error; err != nil {
		return nil, err
	}
	return cartModel, nil
}

func (r *cartRepository) GetByID(cartID string) (*domain.Cart, error) {
	var cartEntity *entity.Cart
	if err := r.db.Where("id = ?", cartID).Order("id").First(&cartEntity).Error; err != nil {
		return nil, err
	}
	cartModel, err := parseToModel(cartEntity)
	if err != nil {
		return nil, err
	}
	return cartModel, nil
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

func parseToEntity(cart *domain.Cart) (*entity.Cart, error) {
	productsStr, err := strSerialize(cart.Products)
	if err != nil {
		return nil, err
	}
	cartEntity := &entity.Cart{
		ID:       cart.ID,
		Products: productsStr,
		CreateAt: cart.CreateAt,
		UpdateAt: cart.UpdateAt,
	}
	return cartEntity, nil
}

func parseToModel(cart *entity.Cart) (*domain.Cart, error) {
	productsData, err := strUnserialize(cart.Products)
	if err != nil {
		return nil, err
	}
	cartModel := &domain.Cart{
		ID:       cart.ID,
		Products: productsData,
		CreateAt: cart.CreateAt,
		UpdateAt: cart.UpdateAt,
	}
	return cartModel, nil
}

func strSerialize(sa []domain.CartItem) (string, error) {
	s, err := json.Marshal(sa)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func strUnserialize(s string) ([]domain.CartItem, error) {
	var ca []domain.CartItem
	err := json.Unmarshal([]byte(s), &ca)
	return ca, err
}
