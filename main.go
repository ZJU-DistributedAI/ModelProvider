//go:generate goagen bootstrap -d ModelProvider/design

package main

import (
	"ModelProvider/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("DataClient")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "ComputingProvider" controller
	c := NewComputingProviderController(service)
	app.MountComputingProviderController(service, c)
	// Mount "swagger" controller
	c2 := NewSwaggerController(service)
	app.MountSwaggerController(service, c2)
	// Mount "swagger-ui" controller
	c3 := NewSwaggerUIController(service)
	app.MountSwaggerUIController(service, c3)

	// Start service
	if err := service.ListenAndServe(":2626"); err != nil {
		service.LogError("startup", "err", err)
	}

}
