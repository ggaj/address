package usecase_address_test

import (
	"errors"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/ggaj/address/src/application/usecase_address"
	"github.com/ggaj/address/src/domain/entities"
	"github.com/ggaj/address/src/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockAddressRepository struct {
	mock.Mock
}

func (mock *MockAddressRepository) Insert(address *entities.Address) (*entities.Address, error) {
	args := mock.Called()
	return args.Get(0).(*entities.Address), args.Error(1)
}

type MockAddressService struct {
	mock.Mock
}

func (mock *MockAddressService) SearchByPostalcode(postalcode string) (*service.AddressResponseDto, error) {
	args := mock.Called()
	return args.Get(0).(*service.AddressResponseDto), args.Error(1)
}

func TestHandle(t *testing.T) {
	mockRepo := new(MockAddressRepository)
	mockService := new(MockAddressService)

	addressResponseDto := service.AddressResponseDto{}
	faker.FakeData(&addressResponseDto)

	address, err := entities.Constructor(
		addressResponseDto.PostalCode,
		addressResponseDto.Street,
		addressResponseDto.Neighborhood,
		addressResponseDto.City,
		addressResponseDto.State,
	)

	addressInput := usecase_address.AddressInput{
		PostalCode: addressResponseDto.PostalCode,
	}

	mockService.On("SearchByPostalcode").Return(&addressResponseDto, nil)
	mockRepo.On("Insert").Return(address, nil)

	createAddressByPostalCode := usecase_address.Constructor(mockRepo, mockService)
	addressOutput, err := createAddressByPostalCode.Handle(&addressInput)

	mockRepo.AssertExpectations(t)
	mockService.AssertExpectations(t)

	require.NotNil(t, addressOutput)
	require.Nil(t, err)

}

func TestSearchByPostalCodeFailed(t *testing.T) {

	postalCode := "69000-000"
	mockRepo := new(MockAddressRepository)
	mockService := new(MockAddressService)

	addressResponseDto := service.AddressResponseDto{}
	addressResponseDto.PostalCode = postalCode

	addressInput := usecase_address.AddressInput{
		PostalCode: addressResponseDto.PostalCode,
	}

	mockService.On("SearchByPostalcode").Return(&addressResponseDto, errors.New("400 is Bad request"))

	createAddressByPostalCode := usecase_address.Constructor(mockRepo, mockService)
	addressOutput, err := createAddressByPostalCode.Handle(&addressInput)

	mockRepo.AssertExpectations(t)

	require.Nil(t, addressOutput)
	require.NotNil(t, err)
	assert.Equal(t, "400 is Bad request", err.Error())

}

func TestAddressFailed(t *testing.T) {

	postalCode := "69085-288"
	mockRepo := new(MockAddressRepository)
	mockService := new(MockAddressService)

	addressResponseDto := service.AddressResponseDto{}
	addressResponseDto.PostalCode = postalCode

	addressInput := usecase_address.AddressInput{
		PostalCode: addressResponseDto.PostalCode,
	}

	mockService.On("SearchByPostalcode").Return(&addressResponseDto, nil)

	createAddressByPostalCode := usecase_address.Constructor(mockRepo, mockService)
	addressOutput, err := createAddressByPostalCode.Handle(&addressInput)

	mockRepo.AssertExpectations(t)

	require.Nil(t, addressOutput)
	require.NotNil(t, err)

}

func TestInvalidInsert(t *testing.T) {
	mockRepo := new(MockAddressRepository)
	mockService := new(MockAddressService)

	addressResponseDto := service.AddressResponseDto{}
	faker.FakeData(&addressResponseDto)

	address, err := entities.Constructor(
		addressResponseDto.PostalCode,
		addressResponseDto.Street,
		addressResponseDto.Neighborhood,
		addressResponseDto.City,
		addressResponseDto.State,
	)

	addressInput := usecase_address.AddressInput{
		PostalCode: addressResponseDto.PostalCode,
	}

	mockService.On("SearchByPostalcode").Return(&addressResponseDto, nil)
	mockRepo.On("Insert").Return(address, errors.New(""))

	createAddressByPostalCode := usecase_address.Constructor(mockRepo, mockService)
	addressOutput, err := createAddressByPostalCode.Handle(&addressInput)

	mockRepo.AssertExpectations(t)
	mockService.AssertExpectations(t)

	require.Nil(t, addressOutput)
	require.NotNil(t, err)

}
