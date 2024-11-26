package banner

type BannerService interface {
	GetBanners() ([]Banner, error)
}

type bannerService struct {
	repo BannerRepository
}

func NewBannerService(repo BannerRepository) BannerService {
	return &bannerService{repo: repo}
}

func (s *bannerService) GetBanners() ([]Banner, error) {
	return s.repo.GetAllBanners()
}
