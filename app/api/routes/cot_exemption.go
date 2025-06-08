package routes

import (
	"github.com/gin-gonic/gin"

	"sui.cot-exemption/app/domain"
)

type CotExemptsRouter struct {
	CotExemptsDomain *domain.CotExemptsDomain
}

func BuildCotExemptsRouter() *CotExemptsRouter {
	return &CotExemptsRouter{
		CotExemptsDomain: domain.NewCotExemptsDomain(),
	}
}

func (r *CotExemptsRouter) GetRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/verify-cot-exemption", r.CotExemptsDomain.VerifyCotExemption)
}
