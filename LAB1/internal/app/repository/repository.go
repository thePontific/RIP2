package repository

import (
	"fmt"
	"strings"
)

type Repository struct {
	currentCartID int // для заглушки корзины
}

func NewRepository() (*Repository, error) {
	return &Repository{
		currentCartID: 1, // заглушка - всегда ID = 1
	}, nil
}

type Order struct {
	ID            int
	Title         string
	Distance      float32
	StarType      string  // тип звезды
	Magnitude     float32 // звездная величина
	Description   string  // описание
	Mass          float32 // масса (в солнечных массах)
	Temperature   int     // температура поверхности
	DiscoveryDate string  // дата открытия (добавьте это поле)
}

// CartItem представляет элемент в корзине
type CartItem struct {
	OrderID   int
	Quantity  int
	Comment   string
	IsPrimary bool
}

// Cart представляет корзину/заявку
type Cart struct {
	ID     int
	Items  []CartItem
	Total  float32 // итоговая стоимость (заглушка)
	Status string
}

func (r *Repository) GetOrder(id int) (Order, error) {
	orders, err := r.GetOrders()
	if err != nil {
		return Order{}, err
	}

	for _, order := range orders {
		if order.ID == id {
			return order, nil
		}
	}
	return Order{}, fmt.Errorf("заказ не найден")
}

func (r *Repository) GetOrders() ([]Order, error) {
	orders := []Order{
		{
			ID:            1,
			Title:         "Майалл II",
			Distance:      8150,
			StarType:      "Переменная звезда",
			Magnitude:     -5.8,
			Description:   "Яркая переменная звезда в скоплении",
			Mass:          12.5,
			Temperature:   3500,
			DiscoveryDate: "В 1940 году",
		},
		{
			ID:            2,
			Title:         "M31N 2008",
			Distance:      18000,
			StarType:      "Новая звезда",
			Magnitude:     -7.2,
			Description:   "Повторная новая в галактике Андромеды",
			Mass:          1.4,
			Temperature:   28000,
			DiscoveryDate: "В 2008 году",
		},
		{
			ID:            3,
			Title:         "S Андромеды",
			Distance:      37500,
			StarType:      "Сверхновая",
			Magnitude:     -18.5,
			Description:   "Историческая сверхновая 1885 года",
			Mass:          8.7,
			Temperature:   10000,
			DiscoveryDate: "В 1885 году",
		},
		{
			ID:            4,
			Title:         "NGC 206",
			Distance:      49000,
			StarType:      "Звездное скопление",
			Magnitude:     -9.1,
			Description:   "Крупнейшее звездное облако в Андромеде",
			Mass:          15000,
			Temperature:   4500,
			DiscoveryDate: "В 1784 году",
		},
		{
			ID:            5,
			Title:         "AE Андромеды",
			Distance:      19500,
			StarType:      "Двойная система",
			Magnitude:     -6.3,
			Description:   "Затменная двойная звезда",
			Mass:          6.8,
			Temperature:   12000,
			DiscoveryDate: "В 1922 году",
		},
		{
			ID:            6,
			Title:         "M31-V1",
			Distance:      58700,
			StarType:      "Цефеида",
			Magnitude:     -4.2,
			Description:   "Первая обнаруженная цефеида в Андромеде",
			Mass:          5.2,
			Temperature:   6000,
			DiscoveryDate: "В 1923 году",
		},
	}

	if len(orders) == 0 {
		return nil, fmt.Errorf("массив пустой")
	}

	return orders, nil
}

func (r *Repository) GetOrdersByTitle(title string) ([]Order, error) {
	orders, err := r.GetOrders()
	if err != nil {
		return []Order{}, err
	}

	var result []Order
	for _, order := range orders {
		if strings.Contains(strings.ToLower(order.Title), strings.ToLower(title)) {
			result = append(result, order)
		}
	}

	return result, nil
}

// Добавляем методы для работы с корзиной
func (r *Repository) GetCart(cartID int) (Cart, error) {
	// Заглушка - возвращаем тестовую корзину
	cart := Cart{
		ID:     cartID,
		Total:  15420.50, // заглушка для вычислений
		Status: "В обработке",
		Items: []CartItem{
			{OrderID: 1, Quantity: 2, Comment: "Срочно", IsPrimary: true},
			{OrderID: 3, Quantity: 1, Comment: "Резерв", IsPrimary: false},
		},
	}
	return cart, nil
}

func (r *Repository) AddToCart(cartID, orderID int, quantity int, comment string, isPrimary bool) error {
	// Заглушка - в реальной реализации здесь будет логика добавления в БД
	return nil
}

func (r *Repository) GetCartItemsCount(cartID int) (int, error) {
	cart, err := r.GetCart(cartID)
	if err != nil {
		return 0, err
	}

	// Считаем общее количество товаров
	totalItems := 0
	for _, item := range cart.Items {
		totalItems += item.Quantity
	}

	return totalItems, nil
}
