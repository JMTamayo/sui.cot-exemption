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

// Server represents the engine of the HTTP server, which is responsible for handling HTTP requests and responses.
type Server struct{}

// BuildServer builds the HTTP server.
//
// Arguments:
//   - ctx: The context of the request.
//
// Returns:
//   - The HTTP server as a gin.Engine instance.
//   - An error if the server fails to build.
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
