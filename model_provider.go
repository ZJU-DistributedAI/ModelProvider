package main

import (
	"ModelProvider/app"
	"fmt"
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
	// ModelProviderController_AskData: start_implement

	// Put your logic her
	return ctx.NotImplemented(goa.ErrInternal("Not implemented"))
	// ModelProviderController_AskData: end_implement
}

// Create runs the create action.
func (c *ModelProviderController) Create(ctx *app.CreateModelProviderContext) error {
	// ModelProviderController_Create: start_implement

	// Put your logic here

	return ctx.NotImplemented(goa.ErrInternal("Not implemented"))
	// ModelProviderController_Create: end_implement
}

// Upload runs the upload action.
func (c *ModelProviderController) Upload(ctx *app.UploadModelProviderContext) error {

	//compress all files to zip, zip file's name is model's name
	err := compress(ctx.ModelHash,ctx.RSAKey,ctx.HEKey,ctx.ETHKey)
	if err != nil{
		fmt.Print("err:",err)
		return ctx.BadRequest(err)
	}

	//upload zip
	zipName := get_file_name(ctx.ModelHash) + ".zip"
	err = uploadModel(zipName)
	if err != nil{
		fmt.Println(err)
	}

	delete_file(zipName)

	return ctx.OK([]byte(ctx.ETHKey))
	//return ctx.NotImplemented(goa.ErrInternal("Not implemented"))
	// ModelProviderController_Upload: end_implement
}

// UploadFinal runs the uploadFinal action.
func (c *ModelProviderController) UploadFinal(ctx *app.UploadFinalModelProviderContext) error {
	// ModelProviderController_UploadFinal: start_implement

	// Put your logic here

	return ctx.NotImplemented(goa.ErrInternal("Not implemented"))
	// ModelProviderController_UploadFinal: end_implement
}
