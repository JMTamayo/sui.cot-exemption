package models

import "time"

type CotExemptionReport struct {
	CompanyId        string `json:"id_empresa"`
	Niu              string `json:"niu"`
	CompanyName      string `json:"nombre_prestador"`
	CompanyShortName string `json:"sigla"`
	FilingCode       string `json:"radicado"`
}

type TviSuperserviciosCotExemptionResponse struct {
	ExemptionReport []CotExemptionReport `json:"listado_exentos"`
}

type CotExemptionResponse struct {
	VerificationTimestamp   time.Time                             `json:"verification_timestamp"`
	SuiCotExemptionResponse TviSuperserviciosCotExemptionResponse `json:"sui_cot_exemption_response"`
}
