package ws

import (
	"github.com/pinzlab/sricore/sri"
)

// ContributorDates contiene las fechas clave relacionadas con el estado del contribuyente.
type ContributorDates struct {
	// StartDate: Fecha en la que el contribuyente inició sus actividades económicas.
	StartDate string `json:"fechaInicioActividades"`

	// CancellationDate: Fecha en la que el contribuyente cesó sus actividades o fue cancelado.
	CancellationDate string `json:"fechaCese"`

	// RestartDate: Fecha en la que el contribuyente reinició sus actividades luego de un cese.
	RestartDate string `json:"fechaReinicioActividades"`

	// UpdateDate: Fecha de la última actualización de los datos del contribuyente.
	UpdateDate string `json:"fechaActualizacion"`
}

// Representative representa a un representante legal del contribuyente.
type Representative struct {
	// Dni: Número de identificación del representante legal.
	Dni string `json:"identificacion"`

	// Name: Nombre completo del representante legal.
	Name string `json:"nombre"`
}

// Contributor contiene la información general y fiscal de un contribuyente.
type Contributor struct {
	// Ruc: Número único de Registro Único de Contribuyente (RUC) asignado al contribuyente.
	Ruc string `json:"numeroRuc"`

	// BusinessName: Razón social o nombre comercial del contribuyente.
	BusinessName string `json:"razonSocial"`

	// Status: Estado actual del contribuyente en la base de datos del SRI.
	Status string `json:"estadoContribuyenteRuc"`

	// EconomicActivity: Descripción de la actividad económica principal del contribuyente.
	EconomicActivity string `json:"actividadEconomicaPrincipal"`

	// Type: Tipo de contribuyente según el SRI.
	Type string `json:"tipoContribuyente"`

	// Regime: Régimen fiscal bajo el cual está registrado el contribuyente.
	// Ejemplos: "Régimen general", "Régimen simplificado", etc.
	Regime string `json:"regimen"`

	// Category: Categoría del contribuyente (si aplica).
	Category *string `json:"categoria"`

	// MustKeepAccounting: Indica si el contribuyente está obligado a llevar contabilidad.
	MustKeepAccounting sri.Bool `json:"obligadoLlevarContabilidad"`

	// WithholdingAgent: Indica si el contribuyente actúa como agente de retención.
	WithholdingAgent sri.Bool `json:"agenteRetencion"`

	// SpecialTaxpayer: Indica si el contribuyente tiene algún régimen especial.
	SpecialTaxpayer sri.Bool `json:"contribuyenteEspecial"`

	// TaxpayerDates: Estructura que contiene las fechas relevantes del contribuyente (inicio, cese, reinicio y última actualización).
	TaxpayerDates ContributorDates `json:"informacionFechasContribuyente"`

	// Representatives: Lista de los representantes legales del contribuyente.
	Representatives []*Representative `json:"representantesLegales"`

	// CancellationReason: Motivo por el cual el contribuyente fue cancelado o suspendido (si aplica).
	CancellationReason *string `json:"motivoCancelacionSuspension"`

	// IsGhostTaxpayer: Indica si el contribuyente está marcado como "contribuyente fantasma" (ficticio).
	IsGhostTaxpayer sri.Bool `json:"contribuyenteFantasma"`

	// NonexistentTransactions: Indica si el contribuyente tiene transacciones inexistentes (fraudulentas).
	NonexistentTransactions sri.Bool `json:"transaccionesInexistente"`
}

// Establishment contiene la información de un establecimiento asociado a un contribuyente.
type Establishment struct {
	// TradeName: Nombre comercial o nombre fantasía del establecimiento.
	TradeName *string `json:"nombreFantasiaComercial"`

	// Type: Tipo de establecimiento según la clasificación del SRI.
	Type string `json:"tipoEstablecimiento"`

	// Address: Dirección completa del establecimiento.
	Address string `json:"direccionCompleta"`

	// Status: Estado actual del establecimiento.
	Status string `json:"estado"`

	// Number: Número único que identifica al establecimiento.
	Number string `json:"numeroEstablecimiento"`

	// IsMain: Indica si el establecimiento es la oficina matriz del contribuyente.
	IsMain sri.Bool `json:"matriz"`
}

// CheckRUC verifica si el RUC existe en la base de datos del SRI.
//
// Este endpoint en parte del API oficial del SRI, pero no están documentados públicamente.
// Su uso puede estar sujeto a cambios o restricciones sin previo aviso.
func (s *SRIOnline) CheckRUC(ruc string) (bool, error) {
	url := s.contributorURL("/ConsolidadoContribuyente/existePorNumeroRuc?numeroRuc=%s", ruc)

	return get[bool](s.client, url)
}

// GetContributors obtiene la información de un contribuyente por su número de RUC.
//
// Este endpoint en parte del API oficial del SRI, pero no están documentados públicamente.
// Su uso puede estar sujeto a cambios o restricciones sin previo aviso.
func (s *SRIOnline) GetContributors(ruc string) ([]*Contributor, error) {
	url := s.contributorURL("/ConsolidadoContribuyente/obtenerPorNumerosRuc?&ruc=%s", ruc)

	return get[[]*Contributor](s.client, url)
}

// GetEstablishments obtiene la información de los establecimientos asociados a un RUC.
//
// Este endpoint en parte del API oficial del SRI, pero no están documentados públicamente.
// Su uso puede estar sujeto a cambios o restricciones sin previo aviso.
func (s *SRIOnline) GetEstablishments(ruc string) ([]*Establishment, error) {
	url := s.contributorURL("/Establecimiento/consultarPorNumeroRuc?numeroRuc=%s", ruc)

	return get[[]*Establishment](s.client, url)
}
