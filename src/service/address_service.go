package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type AddressResponseDto struct {
	PostalCode   string `json:"cep"`
	Street       string `json:"logradouro"`
	Complement   string `json:"complemento"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
}

type IAddressService interface {
	SearchByPostalcode(postalcode string) (*AddressResponseDto, error)
}

type AddressService struct{}

func Constructor() IAddressService {
	return &AddressService{}
}

func (as *AddressService) SearchByPostalcode(postalcode string) (*AddressResponseDto, error) {

	response, err := http.Get("https://viacep.com.br/ws/" + postalcode + "/json/")
	if err != nil {
		return nil, errors.New("Invalid url")
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	body, _ := ioutil.ReadAll(response.Body)

	address := AddressResponseDto{}
	json.Unmarshal(body, &address)

	return &address, nil
}
