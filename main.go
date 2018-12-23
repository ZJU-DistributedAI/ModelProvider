//go:generate goagen bootstrap -d ModelProvider/design

package main

import (
	"ModelProvider/app"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"os"
)

func main() {
	//设置全局参数
	set_all_env_parameter()

	// Create service
	service := goa.New("DataClient")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "ModelProvider" controller
	c := NewModelProviderController(service)
	app.MountModelProviderController(service, c)
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

func set_all_env_parameter(){
	err := os.Setenv("IPFS_HOST","http://47.52.231.230:8899/storage")

	if err != nil {
		fmt.Print(err)
	}
}
