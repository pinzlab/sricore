package sri

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalXML(t *testing.T) {
	// Test case with valid "SI" value
	xmlData := []byte(`<Bool>SI</Bool>`)
	var b Bool
	err := xml.Unmarshal(xmlData, &b)

	// Ensure there is no error and the value is correctly set to true
	assert.NoError(t, err)
	assert.True(t, bool(b))

	// Test case with valid "NO" value
	xmlData = []byte(`<Bool>NO</Bool>`)
	err = xml.Unmarshal(xmlData, &b)

	// Ensure there is no error and the value is correctly set to false
	assert.NoError(t, err)
	assert.False(t, bool(b))

	// Test case with invalid value
	xmlData = []byte(`<Bool>INVALID</Bool>`)
	err = xml.Unmarshal(xmlData, &b)

	// Ensure there is an error for invalid value
	assert.Error(t, err)
}

func TestMarshalXML(t *testing.T) {
	// Test case with true value
	b := Bool(true)

	// Marshal the Bool value into XML
	xmlData, err := xml.Marshal(b)

	// Ensure there is no error during marshaling
	assert.NoError(t, err)

	// Check if the marshaled data matches the expected format ("SI" for true)
	expectedXML := `<Bool>SI</Bool>`
	assert.Equal(t, expectedXML, string(xmlData))

	// Test case with false value
	b = Bool(false)

	// Marshal the Bool value into XML
	xmlData, err = xml.Marshal(b)

	// Ensure there is no error during marshaling
	assert.NoError(t, err)

	// Check if the marshaled data matches the expected format ("NO" for false)
	expectedXML = `<Bool>NO</Bool>`
	assert.Equal(t, expectedXML, string(xmlData))
}

func TestBoolValues(t *testing.T) {
	// Test that the Bool type holds the correct values
	bTrue := Bool(true)
	bFalse := Bool(false)

	// Check if the value is true
	assert.True(t, bool(bTrue))
	// Check if the value is false
	assert.False(t, bool(bFalse))
}

func TestUnmarshalJSON(t *testing.T) {
	// Test case with valid "SI" value
	jsonData := []byte(`"SI"`)
	var b Bool
	err := json.Unmarshal(jsonData, &b)

	// Ensure there is no error and the value is correctly set to true
	assert.NoError(t, err)
	assert.True(t, bool(b))

	// Test case with valid "NO" value
	jsonData = []byte(`"NO"`)
	err = json.Unmarshal(jsonData, &b)

	// Ensure there is no error and the value is correctly set to false
	assert.NoError(t, err)
	assert.False(t, bool(b))

	// Test case with invalid value
	jsonData = []byte(`"INVALID"`)
	err = json.Unmarshal(jsonData, &b)

	// Ensure there is an error for invalid value
	assert.Error(t, err)
}

func TestMarshalJSON(t *testing.T) {
	// Test case with true value
	b := Bool(true)

	// Marshal the Bool value into JSON
	jsonData, err := json.Marshal(b)

	// Ensure there is no error during marshaling
	assert.NoError(t, err)

	// Check if the marshaled data matches the expected format ("SI" for true)
	expectedJSON := `"SI"`
	assert.Equal(t, expectedJSON, string(jsonData))

	// Test case with false value
	b = Bool(false)

	// Marshal the Bool value into JSON
	jsonData, err = json.Marshal(b)

	// Ensure there is no error during marshaling
	assert.NoError(t, err)

	// Check if the marshaled data matches the expected format ("NO" for false)
	expectedJSON = `"NO"`
	assert.Equal(t, expectedJSON, string(jsonData))
}
