package sri

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidBoolXML  = errors.New("valor no válido para Bool en XML")
	ErrInvalidBoolJSON = errors.New("valor no válido para Bool en JSON")
)

// WrapInvalidBoolXML agrega contexto con el valor original recibido.
func WrapInvalidBoolXML(value string) error {
	return fmt.Errorf("%w: %q", ErrInvalidBoolXML, value)
}

// WrapInvalidBoolJSON agrega contexto con el valor original recibido.
func WrapInvalidBoolJSON(value string) error {
	return fmt.Errorf("%w: %q", ErrInvalidBoolJSON, value)
}
