package repository

import "LAB1/internal/app/ds"

// ====== Получить пользователя по ID ======
user, err := r.GetUserByID(userID)
if err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        logrus.Warnf("Пользователь с ID %d не найден, используется заглушка", userID)
        // возможно создать временного пользователя или обработать иначе
    } else {
        return err
    }
}

func (r *Repository) GetDraftCartByCreatorID(creatorID int) (ds.Cart, error) {
	var cart ds.Cart
	err := r.db.Preload("Items").
		Where("creator_id = ? AND status = ?", creatorID, ds.StatusDraft).
		First(&cart).Error
	if err != nil {
		return ds.Cart{}, err
	}
	return cart, nil
}
