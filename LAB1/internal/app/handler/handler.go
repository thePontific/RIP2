package handler

import (
	"LAB1/internal/app/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{
		Repository: r,
	}
}

func (h *Handler) GetOrder(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Error(err)
	}

	order, err := h.Repository.GetOrder(id)
	if err != nil {
		logrus.Error(err)
	}

	ctx.HTML(http.StatusOK, "order.html", gin.H{
		"order": order,
	})
}

func (h *Handler) GetOrders(ctx *gin.Context) {
	var orders []repository.Order
	var err error

	searchQuery := ctx.Query("query")
	if searchQuery == "" {
		orders, err = h.Repository.GetOrders()
		if err != nil {
			logrus.Error(err)
		}
	} else {
		orders, err = h.Repository.GetOrdersByTitle(searchQuery)
		if err != nil {
			logrus.Error(err)
		}
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"time":   time.Now().Format("15:04:05"),
		"orders": orders,
		"query":  searchQuery,
	})
}

func (h *Handler) GetCart(ctx *gin.Context) {
	cartIDStr := ctx.Param("id")
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		logrus.Error("Неверный ID корзины:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID корзины"})
		return
	}

	cart, err := h.Repository.GetCart(cartID)
	if err != nil {
		logrus.Error("Ошибка получения корзины:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения корзины"})
		return
	}

	// Получаем информацию о товарах в корзине
	var cartItemsWithDetails []gin.H
	for _, item := range cart.Items {
		order, err := h.Repository.GetOrder(item.OrderID)
		if err != nil {
			logrus.Errorf("Ошибка получения заказа %d: %v", item.OrderID, err)
			continue
		}

		cartItemsWithDetails = append(cartItemsWithDetails, gin.H{
			"Order":     order,
			"Quantity":  item.Quantity,
			"Comment":   item.Comment,
			"IsPrimary": item.IsPrimary,
		})
	}

	ctx.HTML(http.StatusOK, "cart.html", gin.H{
		"cart":      cart,
		"cartItems": cartItemsWithDetails,
	})
}

func (h *Handler) GetCartItemsCount(ctx *gin.Context) {
	cartID := 1
	count, err := h.Repository.GetCartItemsCount(cartID)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения корзины"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"count": count})
}
