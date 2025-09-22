package repository

import (
	"fmt"
	"strings"
)

// Репозиторий хранит данные о звёздах и корзинах (заглушка)
type Repository struct {
	currentCartID int // для заглушки корзины, всегда = 1
}

// Создает новый репозиторий
func NewRepository() (*Repository, error) {
	return &Repository{
		currentCartID: 1, // заглушка - всегда ID = 1
	}, nil
}

// ====== Структуры данных ======

// Структура звезды
type Star struct {
	ID            int
	Title         string
	Distance      float32 // Расстояние до звезды
	StarType      string  // Тип звезды
	Magnitude     float32 // Светимость
	Description   string  // Описание звезды
	Mass          float32 // Масса
	Temperature   int     // Температура
	DiscoveryDate string  // Дата открытия
	ImageName     string  // Имя изображения для Minio
}

// Элемент корзины, ссылается на заказ
type CartItem struct {
	StarID    int
	Comment   string // Комментарий пользователя
	IsPrimary bool   // Основной элемент (если нужно)
}

// Структура корзины
type Cart struct {
	ID    int
	Items []CartItem // Элементы корзины
}

// ====== Методы репозитория ======

// Найти звезду по ID
func (r *Repository) FindStarByID(id int) (Star, error) {
	stars, err := r.ListStars() // Получаем все звезды
	if err != nil {
		return Star{}, fmt.Errorf("не удалось получить список звезд: %v", err)
	}

	// Ищем звезду с нужным ID
	for _, star := range stars {
		if star.ID == id {
			return star, nil
		}
	}

	return Star{}, fmt.Errorf("звезда с ID %d не найдена", id)
}

// Получить список всех звезд
func (r *Repository) ListStars() ([]Star, error) {
	stars := []Star{
		{ID: 1, Title: "Майалл II", Distance: 8150, StarType: "Переменная звезда", Magnitude: -5.8, Description: "Яркая переменная звезда в скоплении", Mass: 12.5, Temperature: 3500, DiscoveryDate: "В 1940 году", ImageName: "1"},
		{ID: 2, Title: "M31N 2008", Distance: 18000, StarType: "Новая звезда", Magnitude: -7.2, Description: "Повторная новая в галактике Андромеды", Mass: 1.4, Temperature: 28000, DiscoveryDate: "В 2008 году", ImageName: "2"},
		{ID: 3, Title: "S Андромеды", Distance: 37500, StarType: "Сверхновая", Magnitude: -18.5, Description: "Историческая сверхновая 1885 года", Mass: 8.7, Temperature: 10000, DiscoveryDate: "В 1885 году", ImageName: "3"},
		{ID: 4, Title: "NGC 206", Distance: 49000, StarType: "Звездное скопление", Magnitude: -9.1, Description: "Крупнейшее звездное облако в Андромеде", Mass: 15000, Temperature: 4500, DiscoveryDate: "В 1784 году", ImageName: "4"},
		{ID: 5, Title: "AE Андромеды", Distance: 19500, StarType: "Двойная система", Magnitude: -6.3, Description: "Затменная двойная звезда", Mass: 6.8, Temperature: 12000, DiscoveryDate: "В 1922 году", ImageName: "5"},
		{ID: 6, Title: "M31-V1", Distance: 58700, StarType: "Цефеида", Magnitude: -4.2, Description: "Первая обнаруженная цефеида в Андромеде", Mass: 5.2, Temperature: 6000, DiscoveryDate: "В 1923 году", ImageName: "6"},
	}

	if len(stars) == 0 {
		return nil, fmt.Errorf("список звезд пуст")
	}

	return stars, nil
}

// Поиск звезд по названию (регистронезависимо)
func (r *Repository) SearchStarByTitle(title string) ([]Star, error) {
	stars, err := r.ListStars()
	if err != nil {
		return []Star{}, fmt.Errorf("не удалось получить список звезд: %v", err)
	}

	var result []Star
	for _, star := range stars {
		if strings.Contains(strings.ToLower(star.Title), strings.ToLower(title)) {
			result = append(result, star)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("звезды с названием, содержащим '%s', не найдены", title)
	}

	return result, nil
}

// Получить корзину по ID
func (r *Repository) GetCartByID(cartID int) (Cart, error) {
	// Заглушка: возвращает корзину с тремя элементами
	cart := Cart{
		ID: cartID,
		Items: []CartItem{
			{StarID: 1, Comment: "Срочно"},
			{StarID: 3, Comment: "Резерв"},
			{StarID: 2, Comment: "Резерв"},
		},
	}
	return cart, nil
}

// Подсчет количества элементов в корзине
func (r *Repository) CountCartItems(cartID int) (int, error) {
	cart, err := r.GetCartByID(cartID)
	if err != nil {
		return 0, fmt.Errorf("не удалось получить корзину с ID %d: %v", cartID, err)
	}

	return len(cart.Items), nil
}
