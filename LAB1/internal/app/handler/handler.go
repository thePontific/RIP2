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

// Создает новый обработчик с доступом к репозиториям и Minio-сервису
func NewHandler(r *repository.Repository, ms *service.MinioService) *Handler {
	return &Handler{
		Repository:   r,
		MinioService: ms,
	}
}

// ====== Каталог звезд ======
// Возвращает страницу каталога звезд.
// Если задан query-параметр, ищет звезды по названию, иначе возвращает все звезды.
// Также получает количество элементов в корзине для отображения в интерфейсе.
func (h *Handler) GetStars(ctx *gin.Context) {
	var stars []repository.Star
	var err error

	searchQuery := ctx.Query("query")
	if searchQuery == "" {
		stars, err = h.Repository.ListStars() // Получаем все звезды
		if err != nil {
			logrus.Error("Ошибка получения списка звезд:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения списка звезд"})
			return
		}
	} else {
		stars, err = h.Repository.SearchStarByTitle(searchQuery) // Ищем звезды по названию
		if err != nil {
			logrus.Error("Ошибка поиска звезд по названию:", err)
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Звезды не найдены"})
			return
		}
	}

	// Здесь можно динамически получать ID корзины, сейчас для примера = 1
	cartID := 1
	cartItemsCount, _ := h.Repository.CountCartItems(cartID)

	// Отправляем данные в HTML-шаблон
	ctx.HTML(http.StatusOK, "stars_catalog.html", gin.H{
		"time":           time.Now().Format("15:04:05"),
		"stars":          stars,
		"query":          searchQuery,
		"cartID":         cartID,
		"cartItemsCount": cartItemsCount,
		"minioService":   h.MinioService, // Используется для получения URL картинок
	})
}

// ====== Детали звезды ======
// Возвращает страницу с деталями конкретной звезды по ID.
// Загружает объект звезды и формирует URL картинки через MinioService.
func (h *Handler) GetStarDetails(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.Error("Неверный ID звезды:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID звезды"})
		return
	}

	star, err := h.Repository.FindStarByID(id) // Ищем звезду по ID
	if err != nil {
		logrus.Error("Звезда не найдена:", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Звезда не найдена"})
		return
	}

	imageURL := h.MinioService.GetImageURL(star.ImageName) // Получаем URL картинки

	// Отправляем данные в HTML-шаблон
	ctx.HTML(http.StatusOK, "star_details.html", gin.H{
		"Star":     star,
		"imageURL": imageURL,
	})
}

// ====== Корзина / расчет заявки ======
// Возвращает страницу корзины с деталями всех добавленных звезд.
// Загружает все элементы корзины, детали каждой звезды и количество элементов.
func (h *Handler) GetCartDetails(ctx *gin.Context) {
	cartIDStr := ctx.Param("id")
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		logrus.Error("Неверный ID корзины:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID корзины"})
		return
	}

	cart, err := h.Repository.GetCartByID(cartID) // Получаем корзину по ID
	if err != nil {
		logrus.Error("Ошибка получения корзины:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения корзины"})
		return
	}

	cartItemsCount, _ := h.Repository.CountCartItems(cartID) // Считаем количество элементов в корзине

	// Получаем все ID звезд из корзины
	var starIDs []int
	for _, item := range cart.Items {
		starIDs = append(starIDs, item.StarID)
	}

	// Получаем сразу все звезды, которые есть в корзине (оптимизация)
	starsMap := make(map[int]repository.Star)
	for _, starID := range starIDs {
		star, err := h.Repository.FindStarByID(starID)
		if err != nil {
			logrus.Warnf("Звезда с ID %d не найдена: %v", starID, err)
			continue
		}
		starsMap[starID] = star
	}

	// Формируем массив с деталями каждого элемента корзины
	var cartItemsWithDetails []gin.H
	for _, item := range cart.Items {
		star, ok := starsMap[item.StarID]
		if !ok {
			continue // если звезда не найдена, пропускаем
		}

		cartItemsWithDetails = append(cartItemsWithDetails, gin.H{
			"Star":      star,
			"Comment":   item.Comment,
			"IsPrimary": item.IsPrimary,
		})
	}

	// Отправляем данные в HTML-шаблон
	ctx.HTML(http.StatusOK, "cart_calc_speed.html", gin.H{
		"cart":           cart,
		"cartItems":      cartItemsWithDetails,
		"cartItemsCount": cartItemsCount,
		"minioService":   h.MinioService,
	})
}
