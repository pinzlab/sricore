package sri

import (
	"encoding/xml"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const (
	// Length of the access key without the validator digit
	accessKeyLength = 48

	// Date in ddmmaaaa format
	dateFormat = "02012006"
)

// AccessKey define las propiedades necesarias para generar una clave de acceso
// única para un comprobante electrónico. Esta clave está compuesta por varios elementos
// que cumplen con los requisitos establecidos por el SRI (Servicio de Rentas Internas).
type AccessKey struct {
	// Date es la fecha de emisión del comprobante en formato "ddmmaaaa".
	Date time.Time

	// VoucherType es el tipo de comprobante. Debe ser uno de los valores constantes predefinidos.
	VoucherType VoucherType

	// RUC es el número de RUC del contribuyente. Este valor debe tener 13 dígitos.
	RUC string

	// Env es el tipo de ambiente. Debe ser uno de los valores constantes predefinidos.
	Env EnvType

	// Establishment representa la identificación única del establecimiento.
	// Este valor es específico del establecimiento donde se genera el comprobante.
	// Ejemplo: "001" para el primer establecimiento del contribuyente.
	Establishment string

	// EmissionPoint representa la identificación única del punto de emisión.
	// Este valor se refiere al punto desde donde se emite el comprobante electrónico.
	// Ejemplo: "001" para el primer punto de emisión del establecimiento del contribuyente.
	EmissionPoint string

	// Sequential es el número secuencial del comprobante. Este valor debe ser una cadena
	// de 9 caracteres, con ceros a la izquierda si es necesario.
	// Ejemplo: "000000001"
	Sequential string

	// Code es el código numérico generado automáticamente. Este valor debe tener 8 dígitos.
	// Ejemplo: "12345678"
	Code string
}

// string genera el valor base para la clave de acceso concatenando
// los campos relevantes como fecha, tipo de comprobante, RUC, ambiente,
// serie, número secuencial, código y tipo de emisión. También valida
// la longitud y el formato del valor generado para asegurar que cumpla
// con las condiciones requeridas (48 dígitos, todos numéricos).
//
// Retorna:
//   - El valor base de la clave de acceso (una cadena numérica de 48 caracteres) si es válido.
//   - Un error si el valor no tiene la longitud o formato correctos.
func (ak *AccessKey) string() (string, error) {

	// Concatenate all the relevant fields to form the base value (48 digits)
	value := ak.Date.Format(dateFormat) + // Date in ddmmaaaa format
		string(ak.VoucherType) +
		ak.RUC +
		string(ak.Env) +
		ak.Establishment +
		ak.EmissionPoint +
		ak.Sequential +
		ak.Code +
		string(IssueNormal)

	// Check that the value is numeric
	if !regexp.MustCompile(`^\d{48}$`).MatchString(value) {
		return "", errors.New("invalid_ak_format")
	}

	return value, nil
}

// GetNumber genera el número completo del comprobante concatenando el
// establecimiento, punto de emisión y número secuencial. Este formato se
// usa para identificar de forma única el comprobante dentro del
// establecimiento y punto de emisión específicos.
func (ak *AccessKey) GetNumber() string {
	return fmt.Sprintf("%s-%s-%s", ak.Establishment, ak.EmissionPoint, ak.Sequential)
}

// GetSerie genera la serie del comprobante concatenando el establecimiento
// y el punto de emisión. Esto ayuda a identificar la serie de donde se
// origina el comprobante con base en el establecimiento y punto de emisión.
func (ak *AccessKey) GetSerie() string {
	return ak.Establishment + ak.EmissionPoint
}

// Generate calcula la clave de acceso completa generando primero el valor base
// y luego aplicando el algoritmo módulo 11 para determinar el dígito verificador.
// Esta función garantiza que la clave de acceso generada sea válida y completa.
//
// Retorna:
//   - La clave de acceso completa (valor base + dígito verificador) si es exitosa.
//   - Un error si hay un problema al generar el valor base o calcular el dígito verificador.
func (ak *AccessKey) Generate() (string, error) {

	// Apply the modulo 11 algorithm with a weight starting from 7
	weight := 7
	summation := 0
	value, err := ak.string()

	if err != nil {
		return "", err
	}

	// Iterate over each character in the data string
	for index := range accessKeyLength {
		// Convert character to integer and multiply by weight
		value, _ := strconv.Atoi(string(value[index]))

		summation += value * weight

		// Decrease weight, and reset to 7 if it drops below 2
		weight--
		if weight < 2 {
			weight = 7
		}
	}

	// Calculate the result using modulo 11
	result := 11 - (summation % 11)

	// Apply specific rules for result 11 and 10
	if result == 11 {
		result = 0
	} else if result == 10 {
		result = 1
	}

	// Return the concatenated original data with the calculated validator digit
	return value + strconv.Itoa(result), nil
}

// FromString convierte una cadena de clave de acceso en un objeto AccessKey.
// Valida el formato de la clave de acceso proporcionada y extrae sus componentes individuales.
// La clave de acceso debe ser una cadena numérica con exactamente 49 dígitos.
//
// Parámetros:
// - accessKey: Una cadena que representa la clave de acceso a ser interpretada.
//
// Retorna:
// - error: Un error, si alguna validación o análisis falla (por ejemplo, formato inválido o problemas de análisis).
func (ak *AccessKey) FromString(accessKey string) error {

	// Check that the value is numeric
	if !regexp.MustCompile(`^\d{49}$`).MatchString(accessKey) {
		return errors.New("invalid_ak_format")
	}

	// Extract the 48-character base string and the validator digit
	baseKey := accessKey[:accessKeyLength]

	// Extract individual components from the base key
	dateStr := baseKey[:8]                      // ddmmaaaa (8 characters)
	ak.VoucherType = VoucherType(baseKey[8:10]) // 2 characters
	ak.RUC = baseKey[10:23]                     // 13 characters
	ak.Env = EnvType(baseKey[23:24])            // 1 character
	ak.Establishment = baseKey[24:27]           // 3 characters
	ak.EmissionPoint = baseKey[27:30]           // 3 characters
	ak.Sequential = baseKey[30:39]              // 9 characters
	ak.Code = baseKey[39:47]                    // 8 characters

	// Parse the date string into a time.Time object
	var err error
	ak.Date, err = time.Parse(dateFormat, dateStr)
	if err != nil {
		return errors.New("invalid_ak_date")
	}

	return nil
}

// UnmarshalXML implementa el deserializado personalizado para AccessKey.
func (ak *AccessKey) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var akString string
	if err := d.DecodeElement(&akString, &start); err != nil {
		return err
	}

	// Parse the AccessKey string into an AccessKey object
	err := ak.FromString(akString)
	if err != nil {
		return err
	}

	return nil
}

// MarshalXML implementa el serializado personalizado para AccessKey.
func (ak AccessKey) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	akString, err := ak.Generate()
	if err != nil {
		return err
	}
	return e.EncodeElement(akString, start)
}
