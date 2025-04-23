package sri

import "errors"

// TaxType representa el tipo de impuesto en el sistema del SRI (Servicio de Rentas Internas) de Ecuador.
type TaxType string

// Constantes que representan los diferentes tipos de impuestos soportados por el SRI.
const (
	IVA    TaxType = "2" // IVA - Impuesto al Valor Agregado
	ICE    TaxType = "3" // ICE - Impuesto a los Consumos Especiales
	IRBPNR TaxType = "5" // IRBPNR - Impuesto a la Renta de No Residentes
)

const (
	Iva0              = "0"  // Código de IVA para 0%
	Iva5              = "5"  // Código de IVA para 5%
	Iva12             = "2"  // Código de IVA para 12%
	Iva13             = "10" // Código de IVA para 13%
	Iva14             = "3"  // Código de IVA para 14%
	Iva15             = "4"  // Código de IVA para 15%
	IvaNoTaxObject    = "6"  // Código de IVA para "No Objeto de Impuesto"
	IvaExempt         = "7"  // Código de IVA para "Exento de IVA"
	IvaDifferentiated = "8"  // Código de IVA para "IVA diferenciado"
)

// ivaPercents mapea los códigos de IVA con sus porcentajes correspondientes o una cadena vacía para códigos sin porcentaje.
var ivaPercents = map[string]string{
	Iva0:              "0",
	Iva12:             "12",
	Iva14:             "14",
	Iva15:             "15",
	Iva5:              "5",
	Iva13:             "13",
	IvaNoTaxObject:    "",
	IvaExempt:         "",
	IvaDifferentiated: "",
}

// GetIvaPercent obtiene el porcentaje de IVA asociado con el código de IVA proporcionado.
// Retorna el valor del porcentaje si el código es válido, o un error si el código es inválido.
func GetIvaPercent(code string) (*string, error) {
	fee, ok := ivaPercents[code]
	if !ok {
		return nil, errors.New("invalid_iva_code")
	}
	return &fee, nil
}
