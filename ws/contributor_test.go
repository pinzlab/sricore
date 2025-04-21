package ws

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// RUC de la empresa electrica riobamba S.A.
const eersaRuc = "0690000512001"

// service es una instancia global del cliente SRIOnline utilizada.
var service *SRIOnline

func init() {
	service = NewSRIOnline()
}

func TestCheckRUC(t *testing.T) {

	url := service.contributorURL("/ConsolidadoContribuyente/existePorNumeroRuc?numeroRuc=%s", eersaRuc)

	res, err := service.client.Get(url)

	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

}
func TestGetContributors(t *testing.T) {

	url := service.contributorURL("/ConsolidadoContribuyente/obtenerPorNumerosRuc?&ruc=%s", eersaRuc)

	res, err := service.client.Get(url)

	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

}

func TestGetEstablishments(t *testing.T) {

	url := service.contributorURL("/Establecimiento/consultarPorNumeroRuc?numeroRuc=%s", eersaRuc)

	res, err := service.client.Get(url)

	assert.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

}
