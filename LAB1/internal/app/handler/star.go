package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"LAB1/internal/app/ds"
)

// ====== Каталог звёзд ======
func (h *Handler) GetStars(ctx *gin.Context) {
	var stars []ds.Star
	var err error

	searchQuery := ctx.Query("starname")
	if searchQuery == "" {
		stars, err = h.Repository.GetStars()
		if err != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, err)
			return
		}
	} else {
		stars, err = h.Repository.SearchStarByTitle(searchQuery)
		if err != nil {
			h.errorHandler(ctx, http.StatusNotFound, err)
			return
		}
	}

	cartID := 1 // временно фиксируем ID корзины
	cartItemsCount, _ := h.Repository.CountCartItems(cartID)

	ctx.HTML(http.StatusOK, "stars_catalog.html", gin.H{
		"time":           time.Now().Format("15:04:05"),
		"stars":          stars,
		"starname":       searchQuery,
		"cartID":         cartID,
		"cartItemsCount": cartItemsCount,
		"minioService":   h.MinioService,
	})
}

// ====== Детали звезды ======
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

// ====== Корзина ======
func (h *Handler) GetStarscartDetails(ctx *gin.Context) {
	cartIDStr := ctx.Param("id")
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	cart, err := h.Repository.GetCartByID(cartID)
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	cartItemsCount, _ := h.Repository.CountCartItems(cartID)

	var cartItemsWithDetails []gin.H
	for _, item := range cart.Items {
		star, err := h.Repository.GetStar(item.StarID)
		if err != nil {
			logrus.Warnf("Звезда с ID %d не найдена: %v", item.StarID, err)
			continue
		}
		cartItemsWithDetails = append(cartItemsWithDetails, gin.H{
			"Star":      star,
			"Comment":   item.Comment,
			"IsPrimary": item.IsPrimary,
		})
	}

	ctx.HTML(http.StatusOK, "starscart_calc_speed.html", gin.H{
		"cart":           cart,
		"cartItems":      cartItemsWithDetails,
		"cartItemsCount": cartItemsCount,
		"minioService":   h.MinioService,
	})
}
