package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"LAB1/internal/app/ds"
)

// ====== Каталог услуг (звёзд) ======
func (h *Handler) GetStars(ctx *gin.Context) {
	var stars []ds.Star
	var err error

	searchQuery := ctx.Query("starname")
	if searchQuery == "" {
		stars, err = h.Repository.GetStars()
	} else {
		stars, err = h.Repository.SearchStarByTitle(searchQuery)
	}
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	userID := 1 // временно фиксируем пользователя

	// пробуем получить черновую корзину
	cart, err := h.Repository.GetDraftCartByCreatorID(userID)
	cartItemsCount := 0
	if err == nil {
		cartItemsCount, _ = h.Repository.CountCartItems(cart.ID)
	}

	ctx.HTML(http.StatusOK, "stars_catalog.html", gin.H{
		"time":           time.Now().Format("15:04:05"),
		"stars":          stars,
		"starname":       searchQuery,
		"cartID":         cart.ID,
		"cartItemsCount": cartItemsCount,
		"minioService":   h.MinioService,
	})

}

// ====== Детали услуги (звезды) ======
func (h *Handler) GetStarDetails(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	star, err := h.Repository.GetStar(id)
	if err != nil {
		h.errorHandler(ctx, http.StatusNotFound, err)
		return
	}

	imageURL := h.MinioService.GetImageURL(star.ImageName)

	ctx.HTML(http.StatusOK, "star_details.html", gin.H{
		"Star":     star,
		"imageURL": imageURL,
	})
}

// ====== Детали заявки ======
func (h *Handler) GetCartDetails(ctx *gin.Context) {
	userID := 1 // временно фиксируем пользователя

	// Пытаемся получить корзину по ID из URL
	cartIDStr := ctx.Param("id")
	var cart ds.Cart
	var err error

	if cartIDStr != "" {
		cartID, err := strconv.Atoi(cartIDStr)
		if err == nil {
			cart, err = h.Repository.GetCartByID(cartID)
		}
	}

	// Если корзины нет или ID не указан — ищем черновик
	if err != nil {
		cart, err = h.Repository.GetDraftCartByCreatorID(userID)
	}

	// Если черновика всё ещё нет — создаём новый
	if err != nil || cart.ID == 0 {
		cart = ds.Cart{
			CreatorID:  userID,
			Status:     ds.StatusDraft,
			DateCreate: time.Now(),
		}
		if err := h.Repository.CreateCart(&cart); err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	// Считаем количество элементов
	cartItemsCount, _ := h.Repository.CountCartItems(cart.ID)

	// Собираем детали элементов заявки
	var cartItemsWithDetails []gin.H
	for _, item := range cart.Items {
		star, err := h.Repository.GetStar(item.StarID)
		if err != nil {
			logrus.Warnf("Звезда с ID %d не найдена: %v", item.StarID, err)
			continue
		}
		cartItemsWithDetails = append(cartItemsWithDetails, gin.H{
			"Star":     star,
			"Comment":  item.Comment,
			"Quantity": item.Quantity,
		})
	}

	ctx.HTML(http.StatusOK, "starscart_calc_speed.html", gin.H{
		"cart":           cart,
		"cartItems":      cartItemsWithDetails,
		"cartItemsCount": cartItemsCount,
		"minioService":   h.MinioService,
	})
}

// Удаление услуги
func (h *Handler) DeleteStar(ctx *gin.Context) {
	starIDStr := ctx.PostForm("star_id")
	starID, err := strconv.Atoi(starIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID"})
		return
	}

	err = h.Repository.DeleteStar(starID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// После удаления — редирект на страницу каталога
	ctx.Redirect(http.StatusFound, "/Andromeda")
}

//ДОБАВЛЕНИЕ

func (h *Handler) AddStarToCart(ctx *gin.Context) {
	userID := 1 // временно фиксируем пользователя

	starIDStr := ctx.PostForm("star_id")
	quantityStr := ctx.PostForm("quantity")
	comment := ctx.PostForm("comment")

	starID, err := strconv.Atoi(starIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID звезды"})
		return
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity < 1 {
		quantity = 1
	}

	// Получаем черновую корзину пользователя
	cart, err := h.Repository.GetDraftCartByCreatorID(userID)
	if err != nil {
		// если нет черновой корзины — создаём её
		cart = ds.Cart{
			CreatorID:  userID,
			Status:     ds.StatusDraft,
			DateCreate: time.Now(),
		}
		if err := h.Repository.CreateCart(&cart); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// GORM автоматически заполняет cart.ID после Create
	}

	// Проверяем, что cart.ID реально установлен
	if cart.ID == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось получить ID корзины"})
		return
	}

	// Добавляем элемент в корзину
	item := ds.CartItem{
		CartID:   cart.ID,
		StarID:   starID,
		Quantity: quantity,
		Comment:  comment,
	}

	if err := h.Repository.AddCartItem(&item); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Redirect(http.StatusFound, "/Andromeda")
}

// ПОТОМ ЗАКИНУТЬ В КАРД ФАЙЛИК
func (h *Handler) DeleteCart(ctx *gin.Context) {
	cartIDStr := ctx.PostForm("cart_id")
	cartID, _ := strconv.Atoi(cartIDStr)

	err := h.Repository.RawDeleteCartByID(cartID)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusFound, "/Andromeda")
}
