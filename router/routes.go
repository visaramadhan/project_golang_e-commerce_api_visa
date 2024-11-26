package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/address"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/banner"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/cart"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/category"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/orders"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/product"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/users"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/wishlist"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/config"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/manager"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/middleware"
)

func SetupRouter(router *gin.Engine) error {

	router.Use(middleware.LogRequestMiddleware(logrus.New()))

	infraManager := manager.NewInfraManager(config.Cfg)
	serviceManager := manager.NewRepoManager(infraManager)
	repoManager := manager.NewServiceManager(serviceManager)

	userHandler := users.NewHandler(repoManager.UserService())
	bannerHandler := banner.NewBannerHandler(repoManager.BannerService())
	categoryHandler := category.NewCategoryHandler(repoManager.CategoryService())
	productHandler := product.NewProductHandler(repoManager.ProductService())
	bestSelling := product.NewProductHandler(repoManager.BestSellingProductsService())
	wishlistHandler := wishlist.NewWishlistHandler(repoManager.WishlistService())
	cartHandler := cart.NewCartHandler(repoManager.CartService())
	orderHandler := orders.NewOrderHandler(repoManager.OrderService())
	addresHandler := address.NewAddressHandler(repoManager.AddresService())

	v1 := router.Group("/api/v1")
	{
		eCommerce := v1.Group("/e-commerce")
		{
			auth := eCommerce.Group("/auth")
			{
				auth.POST("/register", userHandler.Register)
						auth.POST("/login", userHandler.Login)
			}
			homePage := eCommerce.Group("Home-page")
			{
				homePage.GET("/banners", bannerHandler.GetBanners)
				homePage.GET("/category", categoryHandler.GetCategories)
				homePage.GET("/products", productHandler.GetAllProducts)
				homePage.GET("/best-selling-products", bestSelling.GetBestSellingProducts)
				homePage.PUT("/product/:id/promo", productHandler.UpdatePromoProduct)
				homePage.GET("/recommendations", productHandler.GetRecommendedProducts)
				homePage.POST("/wishlist", middleware.AuthMiddleware, wishlistHandler.AddToWishlist)
				homePage.GET("/wishlist-list", middleware.AuthMiddleware, wishlistHandler.GetAllWishlist)
				homePage.DELETE("/wishlist/:product_id", middleware.AuthMiddleware, wishlistHandler.DeleteWishlist)
			}

			homePageAllProduct := eCommerce.Group("all-product")
			{
				homePageAllProduct.GET("/", productHandler.GetAllProductsByName)
				homePageAllProduct.POST("/cart", cartHandler.AddToCart)
			}

			productByCategory := eCommerce.Group("product-by-category")
			{
				productByCategory.GET("/", productHandler.GetProductsByCategory)
			}

			checkoutFlow := eCommerce.Group("checkout")
			{
				checkoutFlow.GET("product-detail", productHandler.GetProductDetail)
				checkoutFlow.GET("/list", cartHandler.ListCart)
				checkoutFlow.PUT("/update", middleware.AuthMiddleware, cartHandler.UpdateCart)
				checkoutFlow.DELETE("/delete", middleware.AuthMiddleware, cartHandler.DeleteCart)
				checkoutFlow.POST("/add-orders", middleware.AuthMiddleware, orderHandler.CreateOrder)
			}

			accountAddres := eCommerce.Group("account", middleware.AuthMiddleware)
			{
				accountAddres.POST("/create", addresHandler.CreateAddress)
				accountAddres.GET("/addresses", addresHandler.GetAddresses)
				accountAddres.PUT("/addresses/:address_id", addresHandler.UpdateAddress)
				accountAddres.DELETE("/addresses/:address_id", addresHandler.DeleteAddress)
			}
		}
	}

	return router.Run()

}
