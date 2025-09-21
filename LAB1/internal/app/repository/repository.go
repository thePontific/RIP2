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

// repository.go
type Order struct {
	ID            int
	Title         string
	Distance      float32
	StarType      string
	Magnitude     float32
	Description   string
	Mass          float32
	Temperature   int
	DiscoveryDate string
	ImageName     string // Новое поле для названия изображения
}

// CartItem представляет элемент в корзине
type CartItem struct {
	OrderID   int
	Comment   string
	IsPrimary bool
}

// Cart представляет корзину/заявку
type Cart struct {
	ID    int
	Items []CartItem
	//Total  float32
	//Status string
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
			ImageName:     "1",
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
			ImageName:     "2",
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
			ImageName:     "3",
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
			ImageName:     "4",
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
			ImageName:     "5",
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
			ImageName:     "6",
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

func (r *Repository) GetCart(cartID int) (Cart, error) {
	// Заглушка - возвращаем тестовую корзину
	cart := Cart{
		ID: cartID,
		Items: []CartItem{
			{OrderID: 1, Comment: "Срочно"},
			{OrderID: 3, Comment: "Резерв"},
			{OrderID: 2, Comment: "Резерв"},
		},
	}
	return cart, nil
}

// Считаем просто количество элементов в корзине
func (r *Repository) GetCartItemsCount(cartID int) (int, error) {
	cart, err := r.GetCart(cartID)
	if err != nil {
		return 0, err
	}

	// Просто возвращаем количество элементов
	return len(cart.Items), nil
}
