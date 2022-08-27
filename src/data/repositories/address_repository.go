package repositories

import (
	"github.com/ggaj/address/src/domain/entities"
	"github.com/jinzhu/gorm"
)

type IAddressRepository interface {
	Insert(address *entities.Address) (*entities.Address, error)
}

type AddressRepository struct {
	db *gorm.DB
}

func Constructor(db *gorm.DB) IAddressRepository {
	return &AddressRepository{db}
}

func (ar *AddressRepository) Insert(address *entities.Address) (*entities.Address, error) {

	err := ar.db.Create(address).Error
	if err != nil {
		return nil, err
	}

	return address, nil
}
