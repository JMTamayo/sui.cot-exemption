package domain

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"sui.cot-exemption/app/models"
	"sui.cot-exemption/app/services"
)

// CotExemptsDomain represents the domain for the cot exemption API.
type CotExemptsDomain struct {
	tviSuperserviciosService *services.TviSuperserviciosService
}

// NewCotExemptsDomain creates a new cot exemption domain.
//
// Arguments:
//   - None.
//
// Returns:
//   - The cot exemption domain.
func NewCotExemptsDomain() *CotExemptsDomain {
	return &CotExemptsDomain{}
}

// VerifyCotExemption verifies the COT exemption for a given niu.
// @Summary Verify COT exemption for a given niu
// @Tags SUI
// @Accept json
// @Produce json
// @Param niu query string true "niu"
// @Success 200 {object} models.CotExemptionResponse
// @Failure 400 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router /sui/verify-cot-exemption [get]
func (d *CotExemptsDomain) VerifyCotExemption(c *gin.Context) {
	niu := c.Query("niu")

	report, err := d.tviSuperserviciosService.VerifyCotExempts(&niu)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	cotExemptionResponse := models.CotExemptionResponse{
		VerificationTimestamp:   time.Now(),
		SuiCotExemptionResponse: *report,
	}

	c.JSON(http.StatusOK, cotExemptionResponse)
}
