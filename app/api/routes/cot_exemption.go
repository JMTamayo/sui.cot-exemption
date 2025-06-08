package routes

import (
	"github.com/gin-gonic/gin"

	"sui.cot-exemption/app/domain"
)

// CotExemptsRouter represents the router for the cot exemption API.
type CotExemptsRouter struct {
	CotExemptsDomain *domain.CotExemptsDomain
}

// BuildCotExemptsRouter builds the cot exemption router.
//
// Arguments:
//   - None.
//
// Returns:
//   - The cot exemption router.
func BuildCotExemptsRouter() *CotExemptsRouter {
	return &CotExemptsRouter{
		CotExemptsDomain: domain.NewCotExemptsDomain(),
	}
}

// GetRoutes gets the routes for the cot exemption API.
//
// Arguments:
//   - routerGroup: The router group to add the routes to.
//
// Returns:
//   - None.
func (r *CotExemptsRouter) GetRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/verify-cot-exemption", r.CotExemptsDomain.VerifyCotExemption)
}
