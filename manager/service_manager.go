package manager

import (
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/address"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/banner"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/cart"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/category"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/orders"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/product"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/users"
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/api/wishlist"
)

type ServiceManager interface {
	UserService() users.Service
	BannerService() banner.BannerService
	CategoryService() category.CategoryService
	ProductService() product.ProductService
	BestSellingProductsService() product.ProductService
	WishlistService() wishlist.WishlistService
	CartService() cart.CartService
	OrderService() orders.OrderService
	AddresService() address.AddressService
}

// serviceManager adalah implementasi dari interface ServiceManager.
type serviceManager struct {
    repoManager RepoManager
}

// NewServiceManager membuat instance baru dari serviceManager.
func NewServiceManager(repoManager RepoManager) ServiceManager {
    return &serviceManager{
        repoManager: repoManager,
    }
}

func (m *serviceManager) AddresService() address.AddressService {
	return address.NewAddressService(m.repoManager.AddresRepo())
}
func (m *serviceManager) OrderService() orders.OrderService {
	return orders.NewOrderService(m.repoManager.OrderRepo())
}
func (m *serviceManager) UserService() users.Service {
	return users.NewService(m.repoManager.UserRepo())
}

func (m *serviceManager) BannerService() banner.BannerService {
	return banner.NewBannerService(m.repoManager.BannerRepo())
}

func (m *serviceManager) CategoryService() category.CategoryService {
	return category.NewService(m.repoManager.CategoryRepo())
}

func (m *serviceManager) ProductService() product.ProductService {
	return product.NewProductService(m.repoManager.ProductRepo())
}

func (m *serviceManager) BestSellingProductsService() product.ProductService {
	return product.NewProductService(m.repoManager.BestSellingProductsRepo())
}

func (m *serviceManager) WishlistService() wishlist.WishlistService {
	return wishlist.NewWishlistService(m.repoManager.WishlistRepo())
}

func (m *serviceManager) CartService() cart.CartService {
	return cart.NewCartService(
		m.repoManager.CartRepo(),
		m.ProductService(), // Tambahkan ProductService di sini
	)
}
