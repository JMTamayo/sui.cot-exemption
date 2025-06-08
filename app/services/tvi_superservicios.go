package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"sui.cot-exemption/app/config"
	"sui.cot-exemption/app/models"
)

// TviSuperserviciosService represents the service for the TVI Superservicios.
type TviSuperserviciosService struct {
	client *http.Client
}

// NewTviSuperserviciosService creates a new TVI Superservicios service.
//
// Arguments:
//   - host: The host of the TVI Superservicios service.
//
// Returns:
//   - The TVI Superservicios service.
func NewTviSuperserviciosService(host string) *TviSuperserviciosService {
	return &TviSuperserviciosService{
		client: &http.Client{},
	}
}

// VerifyCotExempts verifies the COT exemption for a given niu.
//
// Arguments:
//   - niu: The niu to verify the COT exemption for.
//
// Returns:
//   - The COT exemption response.
func (s *TviSuperserviciosService) VerifyCotExempts(niu *string) (*models.TviSuperserviciosCotExemptionResponse, *models.HTTPError) {
	url := fmt.Sprintf("%s?NIU=%s", config.Conf.GetTviSuperserviciosVerificarExentosHost(), *niu)

	resp, err := http.Get(url)
	if err != nil {
		return nil, models.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, models.NewHTTPError(resp.StatusCode, "Error al verificar la exención de cotización")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, models.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var cotExemptionResponse models.TviSuperserviciosCotExemptionResponse
	err = json.Unmarshal(body, &cotExemptionResponse)
	if err != nil {
		return nil, models.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return &cotExemptionResponse, nil
}
