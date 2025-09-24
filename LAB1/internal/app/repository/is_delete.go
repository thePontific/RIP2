package repository

import (
	"LAB1/internal/app/ds"
)

// ResetDeletedStars сбрасывает все is_delete в false для таблицы Star
func (r *Repository) ResetDeletedStars() error {
	return r.db.Model(&ds.Star{}).Where("is_delete = ?", true).Update("is_delete", false).Error
}
