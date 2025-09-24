package ds

// ====== Звезда ======
type Star struct {
	ID            int     `gorm:"primaryKey"`
	Title         string  // Название
	Distance      float32 // Расстояние до звезды
	StarType      string  // Тип звезды
	Magnitude     float32 // Светимость
	Description   string  // Описание звезды
	Mass          float32 // Масса
	Temperature   int     // Температура
	DiscoveryDate string  // Дата открытия
	ImageName     string  // Имя изображения для Minio
}

// ====== Элемент корзины ======
type CartItem struct {
	ID        int    `gorm:"primaryKey"`
	CartID    int    // Внешний ключ на корзину
	StarID    int    // Внешний ключ на звезду
	Comment   string // Комментарий
	IsPrimary bool   // Основной элемент
}

// ====== Корзина ======
type Cart struct {
	ID    int        `gorm:"primaryKey"`
	Items []CartItem `gorm:"foreignKey:CartID"`
}
