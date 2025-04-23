package sri

import "errors"

// Errores base (tipo error)
var (
	ErrInvalidAccessKeyFormat = errors.New("Formato inválido de clave de acceso")
	ErrInvalidAccessKeyDate   = errors.New("Fecha inválida en clave de acceso")
	ErrInvalidAccessKeyDigit  = errors.New("Error al calcular el dígito verificador de la clave de acceso")
	ErrInvalidVoucherDate     = errors.New("Fecha inválida en formato SRI (esperado 02/01/2006)")
)

// Mensajes con formato (tipo string)
const (
	InvalidBoolFormatMsg = "el valor %q no es válido para %q"
)
