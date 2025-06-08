package main

import (
	"github.com/gin-gonic/gin"

	"sui.cot-exemption/app/api"
	"sui.cot-exemption/app/config"
	"sui.cot-exemption/app/docs"
)

func main() {
	vpnClient := server.NewVPNClient()
	err := vpnClient.Login()
	if err != nil {
		panic(err)
	}
	err = vpnClient.Connect()
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.Title = config.Conf.GetServiceName()
	docs.SwaggerInfo.Version = config.Conf.GetServiceVersion()
	docs.SwaggerInfo.Description = config.Conf.GetServiceDescription()

	ctx := gin.Context{}

	server, err := server.BuildServer(&ctx)
	if err != nil {
		panic(err)
	}

	server.Run(config.Conf.GetServiceAddress())
}
