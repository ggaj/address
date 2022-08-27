package entities_test

import (
	"testing"

	"github.com/bxcodec/faker"
	"github.com/ggaj/address/src/domain/entities"
	"github.com/stretchr/testify/require"
)

func TestAddressValid(t *testing.T) {
	address := entities.Address{}
	faker.FakeData(&address)

	addressEntity, err := entities.Constructor(
		address.PostalCode,
		address.Street,
		address.Neighborhood,
		address.City,
		address.State,
	)

	require.NotNil(t, addressEntity)
	require.Equal(t, address.PostalCode, addressEntity.PostalCode)
	require.Nil(t, err)
}

func TestAddressInvalid(t *testing.T) {
	address := entities.Address{}

	addressEntity, err := entities.Constructor(
		address.PostalCode,
		address.Street,
		address.Neighborhood,
		address.City,
		address.State,
	)

	require.Nil(t, addressEntity)
	require.NotNil(t, err)
}
