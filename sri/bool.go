package sri

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// Bool es un tipo personalizado para manejar el marshalling y unmarshalling de booleanos
// con valores específicos ("SI" para true, "NO" para false) según lo requerido por el sistema.
type Bool bool

// UnmarshalXML implementa un unmarshalling personalizado para Bool.
func (b *Bool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var boolStr string
	if err := d.DecodeElement(&boolStr, &start); err != nil {
		return err
	}

	switch boolStr {
	case "SI":
		*b = true
	case "NO":
		*b = false
	default:
		return fmt.Errorf(InvalidBoolFormatMsg, boolStr, start.Name.Local)
	}

	return nil
}

// MarshalXML implementa un marshalling personalizado para Bool.
func (b Bool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var boolStr string

	if b {
		boolStr = "SI"
	} else {
		boolStr = "NO"
	}

	return e.EncodeElement(boolStr, start)
}

// UnmarshalJSON implementa un unmarshalling personalizado para Bool en formato JSON.
func (b *Bool) UnmarshalJSON(data []byte) error {
	var boolStr string
	if err := json.Unmarshal(data, &boolStr); err != nil {
		return err
	}

	switch boolStr {
	case "SI":
		*b = true
	case "NO":
		*b = false
	default:
		return fmt.Errorf(InvalidBoolFormatMsg, boolStr, "Bool")
	}

	return nil
}

// MarshalJSON implementa un marshalling personalizado para Bool en formato JSON.
func (b Bool) MarshalJSON() ([]byte, error) {
	var boolStr string

	if b {
		boolStr = "SI"
	} else {
		boolStr = "NO"
	}

	return json.Marshal(boolStr)
}
