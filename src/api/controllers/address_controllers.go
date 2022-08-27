package controllers

import (
	"net/http"

	"github.com/ggaj/address/src/application/usecase_address"
	"github.com/gin-gonic/gin"
)

type IAddressController interface {
	CreateAddress(g *gin.Context)
}

type AddressController struct {
	CreateAddressByPostalcodeUsecase usecase_address.ICreateAddressByPostalcodeUsecase
}

func Constructor(ca usecase_address.ICreateAddressByPostalcodeUsecase) IAddressController {
	return &AddressController{
		ca,
	}
}

func (ad *AddressController) CreateAddress(g *gin.Context) {

	var addressInput usecase_address.AddressInput
	if err := g.ShouldBindJSON(&addressInput); err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	address, err := ad.CreateAddressByPostalcodeUsecase.Handle(&addressInput)
	if err != nil {
		g.String(http.StatusBadRequest, err.Error())
		return
	}

	g.JSON(http.StatusCreated, address)
}
