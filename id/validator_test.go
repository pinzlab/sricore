package id

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testObject struct {
	value string
	err   error
}

func TestIsDNI(t *testing.T) {
	tests := []testObject{
		{value: "0601234560", err: nil},                          // Valid DNI
		{value: "0601234578", err: nil},                          // Valid DNI
		{value: "0601234586", err: nil},                          // Valid DNI
		{value: "A601234560", err: errInvalidDNIFormat},          // Invalid format
		{value: "0601234561", err: errInvalidChecksum},           // Invalid checksum
		{value: "0001234560", err: errProvinceOutOfRange},        // Invalid province code
		{value: "2501234560", err: errProvinceOutOfRange},        // Invalid province code
		{value: "0661234560", err: errInvalidNaturalContributor}, // Invalid contributor code

	}

	for _, test := range tests {
		t.Run(test.value, func(t *testing.T) {
			result := IsDNI(test.value)
			assert.Equal(t, test.err, result, "For value %s", test.value)
		})
	}
}

func TestIsNaturalRUC(t *testing.T) {
	tests := []testObject{
		{value: "0601234560001", err: nil},                          // Valid RUC
		{value: "0601234578001", err: nil},                          // Valid RUC
		{value: "0601234586001", err: nil},                          // Valid RUC
		{value: "A601234560001", err: errInvalidRUCFormat},          // Invalid format
		{value: "0601234561001", err: errInvalidChecksum},           // Invalid checksum
		{value: "0001234560001", err: errProvinceOutOfRange},        // Invalid province code
		{value: "2501234560001", err: errProvinceOutOfRange},        // Invalid province code
		{value: "0661234560001", err: errInvalidNaturalContributor}, // Invalid contributor code
		{value: "0601234560000", err: errEstabOutOfRange},           // Invalid establishment code

	}

	for _, test := range tests {
		t.Run(test.value, func(t *testing.T) {
			result := IsNaturalRUC(test.value)
			assert.Equal(t, test.err, result, "For value %s", test.value)
		})
	}
}

func TestIsPublicRUC(t *testing.T) {
	tests := []testObject{
		{value: "0661234560001", err: nil},                         // Valid RUC
		{value: "0661234640001", err: nil},                         // Valid RUC
		{value: "0661234720001", err: nil},                         // Valid RUC
		{value: "A661234560001", err: errInvalidRUCFormat},         // Invalid format
		{value: "0661234570001", err: errInvalidChecksum},          // Invalid checksum
		{value: "0061234560001", err: errProvinceOutOfRange},       // Invalid province code
		{value: "2561234560001", err: errProvinceOutOfRange},       // Invalid province code
		{value: "0671234560001", err: errInvalidPublicContributor}, // Invalid contributor code
		{value: "0661234560000", err: errEstabOutOfRange},          // Invalid establishment code

	}

	for _, test := range tests {
		t.Run(test.value, func(t *testing.T) {
			result := IsPublicRUC(test.value)
			assert.Equal(t, test.err, result, "For value %s", test.value)
		})
	}
}

func TestIsPrivateRUC(t *testing.T) {
	tests := []testObject{
		{value: "0691234568001", err: nil},                          // Valid RUC
		{value: "0691234665001", err: nil},                          // Valid RUC
		{value: "0691234762001", err: nil},                          // Valid RUC
		{value: "A691234568001", err: errInvalidRUCFormat},          // Invalid format
		{value: "0691234567001", err: errInvalidChecksum},           // Invalid checksum
		{value: "0091234568001", err: errProvinceOutOfRange},        // Invalid province code
		{value: "2591234568001", err: errProvinceOutOfRange},        // Invalid province code
		{value: "0681234568001", err: errInvalidPrivateContributor}, // Invalid contributor code
		{value: "0691234568000", err: errEstabOutOfRange},           // Invalid establishment code

	}

	for _, test := range tests {
		t.Run(test.value, func(t *testing.T) {
			result := IsPrivateRUC(test.value)
			assert.Equal(t, test.err, result, "For value %s", test.value)
		})
	}
}

func TestIsRuc(t *testing.T) {
	tests := []string{
		"0601234560001", // Valid natural RUC
		"0601234578001", // Valid natural RUC
		"0661234560001", // Valid public RUC
		"0661234640001", // Valid public RUC
		"0691234568001", // Valid private RUC
		"0691234665001", // Valid private RUC
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			result := IsRUC(test)
			assert.NoError(t, result)
		})
	}
}
