package design                                     // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("DataClient", func() {
	Title("Data client to add or delete data")
	Description("Add or delete data")
	Scheme("http")
	Host("localhost:2626")
})

/*
********************************************************
(3)  Model Provider Client
********************************************************
 */

var _ = Resource("ModelProvider", func() {
	BasePath("/model")

	Action("upload", func() {
		// 模型上传之后，在智能合约中应该生成一个transaction id
		Description("upload model")
		Routing(POST("/upload/:hash/:ETH_key/:HE_key/:RSA_key"))
		Params(func() {
			Param("hash", String, "encrypted model IPFS address")
			Param("ETH_key", String, "ETH private key for transaction") // use it to send transaction
			Param("HE_key", String, "Homomorphic Encryption Key")
			Param("RSA_key", String, "RSA public key")
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("askData", func() {
		// 模型方请求之后，在智能合约中应该生成一个请求ID [ request_id ]， 每一个 [ request_id ] 都会对应一个数据和一个运算方
		Description("ask for data")
		Routing(POST("/askData/:hash/:ETH_key/:transID"))
		Params(func() {
			Param("hash", String, "data hash")
			Param("ETH_key", String, "ETH private key for transaction")
			Param("transID", Integer, "ask data for transaction[ID]")
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("uploadFinal", func() {
		Description("upload final argument and final distance for transaction ID")
		Routing(POST("/uploadFinal/:hash/:rsa_key/:trans_id/:ETH_key"))
		Params(func() {
			Param("hash", String, "final arguments hash")
			Param("rsa_key", String, "rsa private key for transaction id")
			Param("trans_id", Integer, "transaction_id")
			Param("ETH_key", String, "ETH private key for transaction")
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

})


var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
})

var _ = Resource("swagger-ui", func() {

	Files("/swagger-ui/*filepath", "swagger-ui/")
})