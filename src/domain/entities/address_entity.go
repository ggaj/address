package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Address struct {
	ID           string    `json:"id" valid:"uuid" gorm:"type:uuid;primary_key"`
	PostalCode   string    `json:"name" valid:"notnull"`
	Street       string    `json:"street" valid:"notnull"`
	Neighborhood string    `json:"neighborhood" valid:"notnull"`
	City         string    `json:"city" valid:"notnull"`
	State        string    `json:"state" valid:"notnull"`
	CreatedAt    time.Time `json:"created_at" valid:"-"`
	UpdatedAt    time.Time `json:"updated_at" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func Constructor(
	postalCode string,
	street string,
	neighborhood string,
	city string,
	state string,
) (*Address, error) {

	address := Address{
		ID:           uuid.NewV4().String(),
		PostalCode:   postalCode,
		Street:       street,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := address.validate()
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func (a *Address) validate() error {

	_, err := govalidator.ValidateStruct(a)
	if err != nil {
		return err
	}

	return nil
}
