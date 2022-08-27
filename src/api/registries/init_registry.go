package registries

import (
	"github.com/ggaj/address/src/api/controllers"
	"github.com/jinzhu/gorm"
)

type IRegistry interface {
	NewAppController() controllers.AppController
}

type Registry struct {
	db *gorm.DB
}

func Constructor(db *gorm.DB) IRegistry {
	return &Registry{db}
}

func (r *Registry) NewAppController() controllers.AppController {
	return controllers.AppController{
		AddressController: r.Constructor(),
	}
}
