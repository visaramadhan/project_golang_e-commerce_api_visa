package product

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/visaramadhan/project_golang_e-commerce_api_visa/dto"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	response, err := h.service.GetAllProducts(limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) GetBestSellingProducts(c *gin.Context) {
	month, err := strconv.Atoi(c.DefaultQuery("month", "11"))
	if err != nil || month <= 0 || month > 12 {
		month = 11
	}

	year, err := strconv.Atoi(c.DefaultQuery("year", "2024"))
	if err != nil || year <= 0 {
		year = 2024
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	fmt.Printf("Fetching best selling products for month: %d, year: %d, page: %d, pageSize: %d\n", month, year, page, pageSize)

	products, err := h.service.GetBestSellingProducts(month, year, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Found %d best selling products\n", len(products))

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func (h *ProductHandler) UpdatePromoProduct(c *gin.Context) {
	var dto dto.PromoProductDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productID := c.Param("id")

	updatedProduct, err := h.service.UpdatePromoProduct(productID, dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update promo product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Promo product updated successfully",
		"data":    updatedProduct,
	})
}

func (h *ProductHandler) GetRecommendedProducts(c *gin.Context) {
	recommendedProducts, err := h.service.GetRecommendedProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get recommended products"})
		return
	}

	fmt.Printf("Recommended Products Response: %+v\n", recommendedProducts)
	c.JSON(http.StatusOK, gin.H{
		"products": recommendedProducts,
	})
}

func (h *ProductHandler) GetAllProductsByName(c *gin.Context) {
	name := c.Query("name")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	products, totalPages, err := h.service.GetProducts(name, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products":     products,
		"total_pages":  totalPages,
		"current_page": page,
	})
}

func (h *ProductHandler) GetProductsByCategory(c *gin.Context) {
	categoryID := c.DefaultQuery("category_id", "")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	products, err := h.service.GetProductsByCategory(categoryID, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch products by category",
			"message": err.Error(),
		})
		return
	}

	totalData := len(products)
	totalPages := (totalData + limit - 1) / limit

	c.JSON(http.StatusOK, gin.H{
		"products":     products,
		"total_data":   totalData,
		"total_pages":  totalPages,
		"current_page": page,
	})
	log.Printf("Fetched products: %+v", products)
}

func (h *ProductHandler) GetProductDetail(c *gin.Context) {
	productID := c.Param("id")
	product, err := h.service.GetProductsByID(productID)
	if err != nil {
		fmt.Println("error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}
