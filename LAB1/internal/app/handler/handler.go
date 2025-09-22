package handler

import (
	"LAB1/internal/app/repository"
	"LAB1/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Repository   *repository.Repository
	MinioService *service.MinioService
}

func NewHandler(r *repository.Repository, ms *service.MinioService) *Handler {
	return &Handler{
		Repository:   r,
		MinioService: ms,
	}
}

// ====== Каталог звезд ======
func (h *Handler) GetStars(ctx *gin.Context) {
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

	cartItemsCount, _ := h.Repository.GetCartItemsCount(1)

	ctx.HTML(http.StatusOK, "stars_catalog.html", gin.H{
		"time":           time.Now().Format("15:04:05"),
		"orders":         orders,
		"query":          searchQuery,
		"cartID":         1,
		"cartItemsCount": cartItemsCount,
		"minioService":   h.MinioService,
	})
}

// ====== Детали звезды ======
func (h *Handler) GetStarDetails(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID звезды"})
		return
	}

	order, err := h.Repository.GetOrder(id)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Звезда не найдена"})
		return
	}

	imageURL := h.MinioService.GetImageURL(order.ImageName)

	ctx.HTML(http.StatusOK, "star_details.html", gin.H{
		"order":    order,
		"imageURL": imageURL,
	})
}

// ====== Корзина / расчет заявки ======
func (h *Handler) GetCartDetails(ctx *gin.Context) {
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

	cartItemsCount, _ := h.Repository.GetCartItemsCount(cartID)

	var cartItemsWithDetails []gin.H
	for _, item := range cart.Items {
		order, err := h.Repository.GetOrder(item.OrderID)
		if err != nil {
			logrus.Errorf("Ошибка получения заказа %d: %v", item.OrderID, err)
			continue
		}

		cartItemsWithDetails = append(cartItemsWithDetails, gin.H{
			"Order":     order,
			"Comment":   item.Comment,
			"IsPrimary": item.IsPrimary,
		})
	}

	ctx.HTML(http.StatusOK, "cart_calc_speed.html", gin.H{
		"cart":           cart,
		"cartItems":      cartItemsWithDetails,
		"cartItemsCount": cartItemsCount,
		"minioService":   h.MinioService,
	})
}
