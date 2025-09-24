package repository

import (
	"LAB1/internal/app/ds"
	"fmt"
)

// "Удаление" услуги — помечаем is_delete = true
func (r *Repository) DeleteStar(starID int) error {
	err := r.db.Model(&ds.Star{}).Where("id = ?", starID).UpdateColumn("is_delete", true).Error
	if err != nil {
		return fmt.Errorf("ошибка при удалении услуги с id %d: %w", starID, err)
	}
	return nil
}

// Получаем все услуги, только не удалённые
func (r *Repository) GetStars() ([]ds.Star, error) {
	var stars []ds.Star
	err := r.db.Where("is_delete = false").Find(&stars).Error
	if err != nil {
		return nil, err
	}
	return stars, nil
}

// Получение услуги по ID
func (r *Repository) GetStar(id int) (ds.Star, error) {
	var star ds.Star
	err := r.db.Where("id = ? AND is_delete = false", id).First(&star).Error
	if err != nil {
		return ds.Star{}, err
	}
	return star, nil
}

// Поиск услуги по названию, только не удалённые
func (r *Repository) SearchStarByTitle(title string) ([]ds.Star, error) {
	var stars []ds.Star
	err := r.db.Where("title ILIKE ? AND is_delete = false", "%"+title+"%").Find(&stars).Error
	if err != nil {
		return nil, err
	}
	return stars, nil
}
