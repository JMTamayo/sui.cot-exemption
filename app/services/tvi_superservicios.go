package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"sui.cot-exemption/app/config"
	"sui.cot-exemption/app/models"
)

type TviSuperserviciosService struct {
	client *http.Client
}

func NewTviSuperserviciosService(host string) *TviSuperserviciosService {
	return &TviSuperserviciosService{
		client: &http.Client{},
	}
}

func (s *TviSuperserviciosService) VerifyCotExempts(niu *string) (*models.TviSuperserviciosCotExemptionResponse, *models.HttpError) {
	url := fmt.Sprintf("%s?NIU=%s", config.Conf.GetTviSuperserviciosVerificarExentosHost(), *niu)

	resp, err := http.Get(url)
	if err != nil {
		return nil, models.NewHttpError(http.StatusInternalServerError, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, models.NewHttpError(resp.StatusCode, "Error al verificar la exención de cotización")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, models.NewHttpError(http.StatusInternalServerError, err.Error())
	}

	var cotExemptionResponse models.TviSuperserviciosCotExemptionResponse
	err = json.Unmarshal(body, &cotExemptionResponse)
	if err != nil {
		return nil, models.NewHttpError(http.StatusInternalServerError, err.Error())
	}

	return &cotExemptionResponse, nil
}
