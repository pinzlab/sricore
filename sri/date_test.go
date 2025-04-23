package sri

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateUnmarshalXML(t *testing.T) {
	// Test case with valid date format
	xmlData := []byte(`<Date>20/02/2020</Date>`)
	var date Date
	err := xml.Unmarshal(xmlData, &date)

	// Ensure there is no error and the date is parsed correctly
	assert.NoError(t, err)
	expectedDate := time.Date(2020, time.February, 20, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, expectedDate, date.Time)

	// Test case with invalid date format
	xmlData = []byte(`<Date>2020-02-20</Date>`)
	err = xml.Unmarshal(xmlData, &date)

	// Ensure there is an error for an invalid date
	assert.Error(t, err)
}

func TestDateMarshalXML(t *testing.T) {
	// Test case with a valid Date
	date := Date{Time: time.Date(2020, time.February, 20, 0, 0, 0, 0, time.UTC)}

	// Marshal the date into XML
	xmlData, err := xml.Marshal(date)

	// Ensure there is no error during marshaling
	assert.NoError(t, err)

	// Check if the marshaled data matches the expected format
	expectedXML := `<Date>20/02/2020</Date>`
	assert.Equal(t, expectedXML, string(xmlData))
}

func TestNow(t *testing.T) {
	// Get the current date using Now()
	currentDate := Now()

	// Ensure the current date is not zero (meaning it was properly initialized)
	assert.NotNil(t, currentDate)
	assert.True(t, currentDate.Time.Before(time.Now().Add(time.Second)))
}
