package sri

import (
	"encoding/xml"
	"fmt"
	"time"
)

// VoucherDateFormat es el formato de fecha ("02/01/2006") usado por el SRI Ecuador.
const VoucherDateFormat = "02/01/2006"

// Dates es un tipo que extiende time.Time para usar el formato de fecha del SRI.
type Date struct {
	time.Time
}

// UnmarshalXML deserializa Dates usando el formato del SRI.
func (date *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var dateStr string
	if err := d.DecodeElement(&dateStr, &start); err != nil {
		return err
	}

	// Parse the date string into a time object
	parsedTime, err := time.Parse(VoucherDateFormat, dateStr)
	if err != nil {
		return fmt.Errorf("failed to parse time: %v", err)
	}

	date.Time = parsedTime
	return nil
}

// MarshalXML serializa Dates usando el formato del SRI.
func (date Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	dateStr := date.Format(VoucherDateFormat)

	return e.EncodeElement(dateStr, start)
}

// Now returns the current local time as a Date object.
//
// It wraps the current time in the custom Date type, which uses a specific
// date format ("02/01/2006") for marshalling and unmarshalling.
//
// Returns:
//
//	Date: The current local time.
func Now() Date {
	currentTime := time.Now()
	return Date{Time: currentTime}
}
