package banner

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BannerHandler struct {
	service BannerService
}

func NewBannerHandler(service BannerService) *BannerHandler {
	return &BannerHandler{service: service}
}

func (h *BannerHandler) GetBanners(c *gin.Context) {
	banners, err := h.service.GetBanners()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch banners"})
		return
	}

	var response []gin.H
	for _, banner := range banners {
		response = append(response, gin.H{
			"photo":     banner.Photo,
			"title":     banner.Title,
			"subtitle":  banner.Subtitle,
			"path_page": banner.PathPage,
		})
	}

	c.JSON(http.StatusOK, response)
}
