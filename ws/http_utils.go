package ws

import (
	"encoding/json"
	"io"
	"log"
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
		log.Printf("GET request error: %v", err)
		return result, ErrHTTPRequest
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code from SRI: %d", resp.StatusCode)
		return result, ErrHTTPStatus
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return result, ErrReadBody
	}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Failed to unmarshal JSON: %v", err)
		return result, ErrJSONUnmarshal
	}

	return result, nil
}
