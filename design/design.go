package design                                     // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Model Provider Client", func() {
	Title("Model Provider Client")
	Description("Model Provider Client")
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

	Action("create", func() {
		Description("create smart contract")
		Routing(POST("/create/:ETH_key/:smart_contract"))
		Params(func() {
			Param("smart_contract", String, "smart contract")					// 智能合约
			Param("ETH_key", String, "ETH private key for transaction")		// 以太坊交易秘钥，以后会隐藏
		})
		Response(OK,  "plain/text")	// 返回智能合约的地址
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotImplemented, ErrorMedia)
	})


	Action("upload", func() {
		Description("upload model")
		Routing(POST("/upload/:model_hash/:ETH_key/:HE_key/:RSA_key/:contract_hash"))
		Params(func() {
			Param("model_hash", String, "encrypted model IPFS address")	// 同态加密之后的模型地址
			Param("ETH_key", String, "ETH private key for transaction")	// 以太坊交易秘钥，以后会隐藏
			Param("HE_key", String, "Homomorphic Encryption Key")			// 同态加密公钥
			Param("RSA_key", String, "RSA public key")					// RSA 公钥
			Param("contract_hash", String, "smart contract hash")			// 智能合约的地址，之后使用智能合约地址代替模型地址
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotImplemented, ErrorMedia)
	})

	Action("askData", func() {
		Description("ask for data")
		Routing(POST("/askData/:data_hash/:ETH_key/:contract_hash"))
		Params(func() {
			// 智能合约地址和被请求的数据的地址可以成为数据请求的唯一标识
			Param("ETH_key", String, "ETH private key for transaction")	// 以太坊交易秘钥，以后会隐藏
			Param("data_hash", String, "data hash")						// 被请求的数据的ipfs地址
			Param("contract_hash", String, "smart contract hash")			// 智能合约的地址
		})
		Response(OK,  "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotImplemented, ErrorMedia)
	})

	//TODO
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
		Response(NotImplemented, ErrorMedia)
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