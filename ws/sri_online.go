package ws

import (
	"fmt"
	"net/http"
)

const (
	sriOnline      string = "https://srienlinea.sri.gob.ec"
	sriContributor string = "/sri-catastro-sujeto-servicio-internet/rest"
)

// SRIOnline es un cliente para interactuar con los servicios del SRI de Ecuador.
type SRIOnline struct {
	client *http.Client
}

// NewSRIOnline crea una nueva instancia de SRIOnline con la URI base y un cliente HTTP.
func NewSRIOnline() *SRIOnline {
	return &SRIOnline{
		client: &http.Client{},
	}
}

// contributorURL construye una URL completa para un endpoint del servicio de contribuyentes.
//
// Ejemplo:
//
//	s.contributorURL("/Establecimiento/consultarPorNumeroRuc?numeroRuc=%s", "1790016919001")
func (s *SRIOnline) contributorURL(endpoint string, args ...any) string {
	return fmt.Sprintf(sriOnline+sriContributor+endpoint, args...)
}
