package usecase_address

import (
	"github.com/ggaj/address/src/data/repositories"
	"github.com/ggaj/address/src/domain/entities"
	"github.com/ggaj/address/src/service"
)

type AddressInput struct {
	PostalCode string `json:"postalcode"`
}

type AddressOutput struct {
	PostalCode   string `json:"postalcode"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

type ICreateAddressByPostalcodeUsecase interface {
	Handle(input *AddressInput) (*AddressOutput, error)
}

type CreateAddressByPostalcodeUsecase struct {
	AddressRepository repositories.IAddressRepository
	AddressService    service.IAddressService
}

func Constructor(ar repositories.IAddressRepository, as service.IAddressService) ICreateAddressByPostalcodeUsecase {
	return &CreateAddressByPostalcodeUsecase{
		AddressRepository: ar,
		AddressService:    as,
	}
}

func (c *CreateAddressByPostalcodeUsecase) Handle(input *AddressInput) (*AddressOutput, error) {

	addressResponse, err := c.AddressService.SearchByPostalcode(input.PostalCode)
	if err != nil {
		return nil, err
	}

	address, err := entities.Constructor(
		addressResponse.PostalCode,
		addressResponse.Street,
		addressResponse.Neighborhood,
		addressResponse.City,
		addressResponse.State,
	)
	if err != nil {
		return nil, err
	}

	address, err = c.AddressRepository.Insert(address)
	if err != nil {
		return nil, err
	}

	return &AddressOutput{
		address.PostalCode,
		address.Street,
		address.Neighborhood,
		address.City,
		address.State,
	}, nil
}
