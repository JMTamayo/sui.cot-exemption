package domain

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"sui.cot-exemption/app/models"
	"sui.cot-exemption/app/services"
)

type CotExemptsDomain struct {
	tviSuperserviciosService *services.TviSuperserviciosService
}

func NewCotExemptsDomain() *CotExemptsDomain {
	return &CotExemptsDomain{}
}

// Verify COT exemption
// @Summary Verify COT exemption for a given niu
// @Tags SUI
// @Accept json
// @Produce json
// @Param niu query string true "niu"
// @Success 200 {object} models.CotExemptionResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /sui/verify-cot-exemption [get]
func (d *CotExemptsDomain) VerifyCotExemption(c *gin.Context) {
	niu := c.Query("niu")

	report, err := d.tviSuperserviciosService.VerifyCotExempts(&niu)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Body})
		return
	}

	cotExemptionResponse := models.CotExemptionResponse{
		VerificationTimestamp:   time.Now(),
		SuiCotExemptionResponse: *report,
	}

	c.JSON(http.StatusOK, cotExemptionResponse)
}
