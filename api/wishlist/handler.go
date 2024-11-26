package wishlist

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	service WishlistService
}

func NewWishlistHandler(service WishlistService) *WishlistHandler {
	return &WishlistHandler{service: service}
}

func (h *WishlistHandler) AddToWishlist(c *gin.Context) {
	// var wishlist Wishlist
	var request struct {
		ProductID string `json:"product_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// wishlist.ID = uuid.NewString()
	// request.ProductID = uuid.NewString()
	err := h.service.AddProductToWishlist(userID.(string), request.ProductID)
	if err != nil {
		fmt.Println("error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product to wishlist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success added product to wishlist"})
}

func (h *WishlistHandler) GetAllWishlist(c *gin.Context) {
	userID := c.GetString("user_id")

	wishlists, err := h.service.GetAllWishlist(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":   userID,
		"wishlists": wishlists,
	})
}

func (h *WishlistHandler) DeleteWishlist(c *gin.Context) {
	userID := c.GetString("user_id")
	productID := c.Param("product_id")

	err := h.service.DeleteWishlist(userID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Wishlist item deleted successfully",
	})
}
