package main

import (
	"ModelProvider/app"
	"github.com/goadesign/goa"
)

// ModelProviderController implements the ModelProvider resource.
type ModelProviderController struct {
	*goa.Controller
}

// NewModelProviderController creates a ModelProvider controller.
func NewModelProviderController(service *goa.Service) *ModelProviderController {
	return &ModelProviderController{Controller: service.NewController("ModelProviderController")}
}

// AskData runs the askData action.
func (c *ModelProviderController) AskData(ctx *app.AskDataModelProviderContext) error {

	return ctx.OK([]byte("Request id"))
}

// Upload runs the upload action.
func (c *ModelProviderController) Upload(ctx *app.UploadModelProviderContext) error {


	return ctx.OK([]byte("Transaction ID"))
}

// UploadFinal runs the uploadFinal action.
func (c *ModelProviderController) UploadFinal(ctx *app.UploadFinalModelProviderContext) error {

	return ctx.OK([]byte("Transaction Hash of ETH"))
}
