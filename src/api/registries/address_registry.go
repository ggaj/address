package registries

import (
	"github.com/ggaj/address/src/api/controllers"
	"github.com/ggaj/address/src/application/usecase_address"
	"github.com/ggaj/address/src/data/repositories"
	"github.com/ggaj/address/src/service"
)

func (r *Registry) Constructor() controllers.IAddressController {
	return controllers.Constructor(
		r.CreateAddressByPostalcodeUsecase(),
	)
}

func (r *Registry) NewAddressRepository() repositories.IAddressRepository {
	return repositories.Constructor(r.db)
}

func (r *Registry) CreateAddressByPostalcodeUsecase() usecase_address.ICreateAddressByPostalcodeUsecase {
	return usecase_address.Constructor(r.NewAddressRepository(), r.AddressService())
}

func (r *Registry) AddressService() service.IAddressService {
	return service.Constructor()
}
