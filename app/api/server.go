package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"sui.cot-exemption/app/api/routes"
	"sui.cot-exemption/app/config"
	"sui.cot-exemption/app/models"
)

type Server struct{}

func BuildServer(ctx *gin.Context) (*gin.Engine, *models.Error) {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.Conf.GetAllowedOrigins()
	corsConfig.AllowMethods = []string{"GET", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	corsConfig.AllowCredentials = true
	router.Use(cors.New(corsConfig))

	cotExemptsRouter := routes.BuildCotExemptsRouter()

	sui := router.Group("/sui")
	{
		cotExemptsRouter.GetRoutes(sui)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router, nil
}
