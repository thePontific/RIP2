package repository

import (
	"fmt"
	"strings"
)

type Repository struct {
}

func NewRepository() (*Repository, error) {
	return &Repository{}, nil
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
			DiscoveryDate: "1940 год",
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
			DiscoveryDate: "2008 год",
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
			DiscoveryDate: "1885 год",
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
			DiscoveryDate: "1784 год",
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
			DiscoveryDate: "1922 год",
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
			DiscoveryDate: "1923 год",
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
