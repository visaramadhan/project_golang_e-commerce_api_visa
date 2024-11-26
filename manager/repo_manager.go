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

type RepoManager interface {
	UserRepo() users.Repository
	BannerRepo() banner.BannerRepository
	CategoryRepo() category.CategoryRepository
	ProductRepo() product.ProductRepository
	BestSellingProductsRepo() product.ProductRepository
	WishlistRepo() wishlist.WishlistRepository
	CartRepo() cart.CartRepository
	OrderRepo() orders.OrderRepository
	AddresRepo() address.AddressRepository
}

type repoManager struct {
	infraManager InfraManager
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infraManager: infra,
	}
}

func (m *repoManager) AddresRepo() address.AddressRepository {
	return address.NewAddressRepository(m.infraManager.Conn())
}

func (m *repoManager) OrderRepo() orders.OrderRepository {
	return orders.NewOrderRepository(m.infraManager.Conn())
}

func (m *repoManager) UserRepo() users.Repository {
	return users.NewRepository(m.infraManager.Conn())
}

func (m *repoManager) BannerRepo() banner.BannerRepository {
	return banner.NewBannerRepository(m.infraManager.Conn())
}

func (m *repoManager) CategoryRepo() category.CategoryRepository {
	return category.NewRepository(m.infraManager.Conn())
}

func (m *repoManager) ProductRepo() product.ProductRepository {
	return product.NewProductRepository(m.infraManager.Conn())
}

func (m *repoManager) BestSellingProductsRepo() product.ProductRepository {
	return product.NewProductRepository(m.infraManager.Conn())
}

func (m *repoManager) WishlistRepo() wishlist.WishlistRepository {
	return wishlist.NewWishlistRepository(m.infraManager.Conn())
}

func (m *repoManager) CartRepo() cart.CartRepository {
	return cart.NewCartRepository(m.infraManager.Conn())
}
