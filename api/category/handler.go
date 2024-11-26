package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service CategoryService
}

func NewCategoryHandler(service CategoryService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetCategories(c *gin.Context) {
	categories, err := h.service.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})
}
