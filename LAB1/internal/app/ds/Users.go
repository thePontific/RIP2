package ds

type Users struct {
	ID          int    `gorm:"primaryKey"`
	Login       string `gorm:"type:varchar(25);unique;not null"`
	Password    string `gorm:"type:varchar(100);not null"`
	IsModerator bool   `gorm:"default:false"`
}
