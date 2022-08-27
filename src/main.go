package main

import (
	"github.com/ggaj/address/src/api/config"
	"github.com/ggaj/address/src/api/registries"
	"github.com/ggaj/address/src/api/routers"
	"github.com/ggaj/address/src/data/database"
)

func main() {

	config.ReadConfig()

	db := database.Connect()
	defer db.Close()

	r := registries.Constructor(db)

	routers.RouteInit(r.NewAppController())
}
