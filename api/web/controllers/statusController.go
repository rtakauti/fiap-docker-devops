package controllers

import (
	"fmt"

	"../../datamodels"
	"../../services"
)

type StatusController struct {
	// Our MovieService, it's an interface which
	// is binded from the main application.
	Service services.StatusServiceTop
}

func (c *StatusController) Get() (results []datamodels.Status) {
	fmt.Printf("%s", c)
	return c.Service.GetAll()
}
