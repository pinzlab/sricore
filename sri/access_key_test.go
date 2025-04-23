package sri

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccessKeyString_Valid(t *testing.T) {
	// Create an AccessKey object with valid data
	ak := &AccessKey{
		Date:          time.Date(2020, time.February, 20, 0, 0, 0, 0, time.UTC),
		VoucherType:   Invoice,
		RUC:           "1791251237001",
		Env:           EnvProd,
		Establishment: "001",
		EmissionPoint: "001",
		Sequential:    "005814917",
		Code:          "12345678",
	}

	// Call the string() method to generate the base access key
	result, err := ak.string()

	// Verify that there is no error
	require.NoError(t, err)

	// Verify that the result has a length of 48 characters
	assert.Equal(t, 48, len(result))

	// Verify that the generated value is numeric
	assert.Regexp(t, `^\d{48}$`, result)

	// Verify that the generated value matches the expected value

	assert.Equal(t, "200220200117912512370012001001005814917123456781", result)
}

func TestAccessKeyString_InvalidLength(t *testing.T) {
	// Create an AccessKey object with valid data but modify one field to make the length incorrect
	ak := &AccessKey{
		Date:          time.Date(2020, time.February, 20, 0, 0, 0, 0, time.UTC),
		VoucherType:   Invoice,
		RUC:           "1791251237001",
		Env:           EnvProd,
		Establishment: "001",
		EmissionPoint: "001",
		Sequential:    "005814917",
		Code:          "123456789", // Change the code to 9 digits to make the total length incorrect
	}

	// Call the string() method and expect an error
	_, err := ak.string()

	// Verify that there is an invalid length error
	assert.Error(t, err)
	assert.Equal(t, "invalid_ak_format", err.Error())
}

func TestAccessKeyString_InvalidFormat(t *testing.T) {
	// Create an AccessKey object with a series containing non-numeric characters
	ak := &AccessKey{
		Date:          time.Date(2020, time.February, 20, 0, 0, 0, 0, time.UTC),
		VoucherType:   Invoice,
		RUC:           "1791251237001",
		Env:           EnvProd,
		Establishment: "00A", // Change the series to contain letters
		EmissionPoint: "001",
		Sequential:    "005814917",
		Code:          "12345678",
	}

	// Call the string() method and expect an error
	_, err := ak.string()

	// Verify that there is a format error
	assert.Error(t, err)
	assert.Equal(t, "invalid_ak_format", err.Error())
}

func TestGenerateAccessKey_Valid(t *testing.T) {
	// Create an AccessKey object with valid data
	ak := &AccessKey{
		Date:          time.Date(2020, time.February, 20, 0, 0, 0, 0, time.UTC),
		VoucherType:   Invoice,
		RUC:           "1234567890001",
		Env:           EnvProd,
		Establishment: "001",
		EmissionPoint: "001",
		Sequential:    "000000001",
		Code:          "12345678",
	}

	// Call the Generate() method to generate the full access key
	result, err := ak.Generate()

	// Verify that there is no error
	require.NoError(t, err)

	// Verify that the result has a length of 49 (48 characters + 1 validator digit)
	assert.Equal(t, 49, len(result))

	// Verify that the generated value is numeric
	assert.Regexp(t, `^\d{49}$`, result)
}

func TestFromString_Success(t *testing.T) {
	tests := []struct {
		accessKey string
		expected  AccessKey
	}{
		{
			accessKey: "2002202001179125123700120010010058149171234567817",
			expected: AccessKey{
				Date:          time.Date(2020, time.February, 20, 0, 0, 0, 0, time.UTC),
				VoucherType:   Invoice,
				RUC:           "1791251237001",
				Env:           EnvProd,
				Establishment: "001",
				EmissionPoint: "001",
				Sequential:    "005814917",
				Code:          "12345678",
			},
		},

		{
			accessKey: "2002202001179125123700120010010058149181234567812",
			expected: AccessKey{
				Date:          time.Date(2020, time.February, 20, 0, 0, 0, 0, time.UTC),
				VoucherType:   Invoice,
				RUC:           "1791251237001",
				Env:           EnvProd,
				Establishment: "001",
				EmissionPoint: "001",
				Sequential:    "005814918",
				Code:          "12345678",
			},
		},

		{
			accessKey: "2002202001179125123700120010010058149121234567811",
			expected: AccessKey{
				Date:          time.Date(2020, time.February, 20, 0, 0, 0, 0, time.UTC),
				VoucherType:   Invoice,
				RUC:           "1791251237001",
				Env:           EnvProd,
				Establishment: "001",
				EmissionPoint: "001",
				Sequential:    "005814912",
				Code:          "12345678",
			},
		},

		{
			accessKey: "2002202001171404598400120010010001837471234567812",
			expected: AccessKey{
				Date:          time.Date(2020, time.February, 20, 0, 0, 0, 0, time.UTC),
				VoucherType:   Invoice,
				RUC:           "1714045984001",
				Env:           EnvProd,
				Establishment: "001",
				EmissionPoint: "001",
				Sequential:    "000183747",
				Code:          "12345678",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.accessKey, func(t *testing.T) {
			var ak AccessKey
			// Call the FromString method
			err := ak.FromString(test.accessKey)

			// Assert that no error occurred
			assert.NoError(t, err, "Failed to parse access key %s", test.accessKey)

			// Assert that the parsed AccessKey struct matches the expected one
			assert.Equal(t, test.expected, ak, "Parsed AccessKey does not match the expected value for access key %s", test.accessKey)
		})
	}
}

func TestFromString_Failure(t *testing.T) {
	tests := []struct {
		accessKey string
		errMsg    string
	}{
		{
			// Extra digit (50 digits instead of 49)
			accessKey: "20022020011791251237001200100100581491712345678171",
			errMsg:    "invalid_ak_format",
		},
		{
			// Missing digit (48 digits instead of 49)
			accessKey: "200220200117912512370012001001005814918123456781",
			errMsg:    "invalid_ak_format",
		},
		{
			// Non-numeric character ('A' is invalid)
			accessKey: "2002202001179125123A00120010010058149121234567811",
			errMsg:    "invalid_ak_format",
		},
		{
			// Invalid date (Invalid date: 30nd of February)
			accessKey: "3002202001171404598400120010010001837471234567812",
			errMsg:    "invalid_ak_date",
		},
	}

	for _, test := range tests {
		t.Run(test.accessKey, func(t *testing.T) {
			ak := &AccessKey{}

			// Call the FromString method
			err := ak.FromString(test.accessKey)

			// Assert that the error message contains the expected error
			if assert.Error(t, err) {
				assert.Contains(t, err.Error(), test.errMsg, "Expected error to contain: "+test.errMsg)
			}
		})
	}
}

func TestAccessKeyUnmarshalXML(t *testing.T) {

	// Unmarshal the AccessKey from the provided XML string
	xmlData := []byte("<AccessKey>2002202001179125123700120010010058149171234567817</AccessKey>")
	var ak AccessKey
	err := xml.Unmarshal(xmlData, &ak)

	// Assert no error during unmarshalling
	assert.NoError(t, err)

	// Assert that the parsed AccessKey struct matches the expected one
	expected := AccessKey{
		Date:          time.Date(2020, time.February, 20, 0, 0, 0, 0, time.UTC),
		VoucherType:   Invoice,
		RUC:           "1791251237001",
		Env:           EnvProd,
		Establishment: "001",
		EmissionPoint: "001",
		Sequential:    "005814917",
		Code:          "12345678",
	}
	assert.Equal(t, expected, ak)

}

func TestAccessKeyMarshalXML(t *testing.T) {

	ak := AccessKey{
		Date:          time.Date(2020, time.February, 20, 0, 0, 0, 0, time.UTC),
		VoucherType:   Invoice,
		RUC:           "1791251237001",
		Env:           EnvProd,
		Establishment: "001",
		EmissionPoint: "001",
		Sequential:    "005814917",
		Code:          "12345678",
	}

	// Marshal the AccessKey into XML
	xmlData, err := xml.Marshal(ak)

	// Assert no error during marshalling
	assert.NoError(t, err)

	// Expected XML output
	expectedXML := "<AccessKey>2002202001179125123700120010010058149171234567817</AccessKey>"
	assert.Equal(t, expectedXML, string(xmlData))
}
