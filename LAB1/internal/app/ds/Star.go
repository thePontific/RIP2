package ds

type Star struct {
	ID            int     `gorm:"primaryKey"`
	Title         string  `gorm:"type:varchar(100);not null"` // Название
	Distance      float32 // Расстояние до звезды
	StarType      string  // Тип звезды
	Magnitude     float32 // Светимость
	Description   string  // Описание звезды
	Mass          float32 // Масса
	Temperature   int     // Температура
	DiscoveryDate string  `gorm:"type:varchar(50)"` // Дата открытия
	ImageName     string  // Имя изображения для Minio
}
