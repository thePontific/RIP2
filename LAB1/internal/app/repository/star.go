package repository

import (
	"fmt"

	"LAB1/internal/app/ds"
)

func (r *Repository) GetStars() ([]ds.Star, error) {
	var stars []ds.Star
	err := r.db.Find(&stars).Error
	if err != nil {
		return nil, err
	}
	if len(stars) == 0 {
		return nil, fmt.Errorf("список звезд пуст")
	}
	return stars, nil
}

func (r *Repository) GetStar(id int) (ds.Star, error) {
	var star ds.Star
	err := r.db.First(&star, id).Error
	if err != nil {
		return ds.Star{}, err
	}
	return star, nil
}

func (r *Repository) SearchStarByTitle(title string) ([]ds.Star, error) {
	var stars []ds.Star
	err := r.db.Where("title ILIKE ?", "%"+title+"%").Find(&stars).Error
	if err != nil {
		return nil, err
	}
	return stars, nil
}

func (r *Repository) GetCartByID(cartID int) (ds.Cart, error) {
	var cart ds.Cart
	err := r.db.Preload("Items").First(&cart, cartID).Error
	if err != nil {
		return ds.Cart{}, err
	}
	return cart, nil
}

func (r *Repository) CountCartItems(cartID int) (int, error) {
	var count int64
	err := r.db.Model(&ds.CartItem{}).Where("cart_id = ?", cartID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
