package ws

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// get realiza una solicitud HTTP GET a la URL indicada y deserializa la respuesta JSON
// en el tipo especificado.
//
// Parameters:
//
//	client: Cliente HTTP para realizar la solicitud.
//	url: URL para la solicitud GET.
//
// Returns:
//   - Un valor deserializado del tipo indicado (T).
//   - Un error si ocurre algún problema durante la solicitud o deserialización.
func get[T any](client *http.Client, url string) (T, error) {
	var result T

	resp, err := client.Get(url)
	if err != nil {
		return result, fmt.Errorf("error realizando la solicitud al SRI: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("respuesta no exitosa del SRI, código de estado: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("error leyendo el cuerpo de la respuesta: %v", err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return result, fmt.Errorf("error procesando la respuesta JSON: %v", err)
	}

	return result, nil
}
