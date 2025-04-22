package ws

import "errors"

var (
	ErrHTTPRequest   = errors.New("No se pudo realizar la solicitud HTTP al SRI")
	ErrHTTPStatus    = errors.New("El SRI respondió con un estado no exitoso")
	ErrReadBody      = errors.New("No se pudo leer el cuerpo de la respuesta")
	ErrJSONUnmarshal = errors.New("No se pudo procesar la respuesta JSON")
)
