package id

import "errors"

var (
	// Format errors
	errInvalidDNIFormat = errors.New("La cédula debe tener exactamente 10 dígitos")
	errInvalidRUCFormat = errors.New("El RUC debe tener exactamente 13 dígitos")

	// Province errors
	errProvinceOutOfRange = errors.New("El código de provincia debe estar entre 01 y 24")

	// Contributor errors
	errInvalidNaturalContributor = errors.New("El código para persona natural debe estar entre 0 y 5")
	errInvalidPrivateContributor = errors.New("El para sociedades privadas debe ser igual a 9")
	errInvalidPublicContributor  = errors.New("El para instituciones públicas debe ser igual a 6")
	errUnknownContributorType    = errors.New("El código de contribuyente no es valido")

	// Establishment errors
	errEstabOutOfRange = errors.New("El código de establecimiento debe ser mayor a 1")

	// checksum errors
	errInvalidChecksum = errors.New("El dígito verificador es inválido")
)
