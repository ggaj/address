package service_test

import (
	"testing"

	"github.com/ggaj/address/src/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSearchByPostalCode(t *testing.T) {

	postalCode := "69085-288"

	addressService := service.Constructor()
	searchByPostalcode, err := addressService.SearchByPostalcode(postalCode)

	require.NotNil(t, searchByPostalcode)
	require.Nil(t, err)

	assert.Equal(t, postalCode, searchByPostalcode.PostalCode)
	assert.NotNil(t, searchByPostalcode.Street)
	assert.NotNil(t, searchByPostalcode.City)
	assert.NotNil(t, searchByPostalcode.State)

}
func TestInvalidUrl(t *testing.T) {

	addressService := service.Constructor()
	_, err := addressService.SearchByPostalcode(";asdsadasd@$%#@")

	require.NotNil(t, err)
	require.Equal(t, "Invalid url", err.Error())

}

func TestStatusCode400(t *testing.T) {

	addressService := service.Constructor()
	_, err := addressService.SearchByPostalcode(";")

	require.NotNil(t, err)
	require.Equal(t, "400 Bad Request", err.Error())

}
