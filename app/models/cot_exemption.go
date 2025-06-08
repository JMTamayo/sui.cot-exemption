package models

import "time"

// CotExemptionReport represents the report of the cot exemption.
type CotExemptionReport struct {
	CompanyID        string `json:"id_empresa"`
	Niu              string `json:"niu"`
	CompanyName      string `json:"nombre_prestador"`
	CompanyShortName string `json:"sigla"`
	FilingCode       string `json:"radicado"`
}

// TviSuperserviciosCotExemptionResponse represents the response of the TVI Superservicios cot exemption service.
type TviSuperserviciosCotExemptionResponse struct {
	ExemptionReport []CotExemptionReport `json:"listado_exentos"`
}

// CotExemptionResponse represents the response of the cot exemption API.
type CotExemptionResponse struct {
	VerificationTimestamp   time.Time                             `json:"verification_timestamp"`
	SuiCotExemptionResponse TviSuperserviciosCotExemptionResponse `json:"report"`
}
