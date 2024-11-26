package manager

import (
	"github.com/visaramadhan/project_golang_e-commerce_api_visa/config"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type InfraManager interface {
	Conn() *gorm.DB
}

type infraManager struct {
	cfg *config.Config
}

func (i *infraManager) Conn() *gorm.DB {
	return config.DB
}

func NewInfraManager(configParam *config.Config) InfraManager {
	infra := &infraManager{
		cfg: configParam,
	}

	return infra
}
