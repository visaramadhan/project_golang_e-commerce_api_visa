package category

type CategoryService interface {
	GetAllCategories() ([]Category, error)
}

type categoryService struct {
	repository CategoryRepository
}

func NewService(repository CategoryRepository) CategoryService {
	return &categoryService{repository: repository}
}

func (s *categoryService) GetAllCategories() ([]Category, error) {
	return s.repository.FindAll()
}
