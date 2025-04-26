package id

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// isDigit checks whether the given value is a numeric string of the expected length
// for the specified document type (dni or ruc).
func isDigit(doc docType, value string) error {
	exp := fmt.Sprintf(`^\d{%d}$`, doc)

	if !regexp.MustCompile(exp).MatchString(value) {
		if doc == dni {
			return errInvalidDNIFormat
		}

		return errInvalidRUCFormat
	}

	return nil

}

// checkProvince validates that the province code (first two digits of the ID)
// is within the valid range (1 to 24), corresponding to Ecuadorian provinces.
func checkProvince(code string) error {
	value, _ := strconv.Atoi(code)

	if value < 1 || value > 24 {
		return errProvinceOutOfRange
	}

	return nil
}

// checkContributor validates the third digit of the RUC to ensure it matches
// the expected pattern for the given contributor type.
func checkContributor(contributor contributorType, value string) error {
	digito, _ := strconv.Atoi(value)

	switch contributor {
	case contributorNatural:
		if digito < 0 || digito > 5 {
			return errInvalidNaturalContributor
		}
	case contributorPrivate:
		if digito != 9 {
			return errInvalidPrivateContributor
		}
	case contributorPublic:
		if digito != 6 {
			return errInvalidPublicContributor
		}
	default:
		return errUnknownContributorType
	}

	return nil
}

// checkEstablishment verifies that the establishment code (typically the last three digits of the RUC)
// is greater than zero, as required by Ecuadorian regulations.
func checkEstablishment(code string) error {
	value, _ := strconv.Atoi(code)

	if value < 1 {
		return errEstabOutOfRange
	}

	return nil
}

// checkModule10 validates the checksum of a value using the Modulo 10 algorithm,
// commonly used for cédulas and natural person RUCs in Ecuador.
func checkModule10(value string) error {
	result := calcModule10(value[:9])

	if !strings.HasSuffix(value, strconv.Itoa(result)) {
		return errInvalidChecksum
	}

	return nil
}

// checkModule11 validates the checksum of a RUC value using the Modulo 11 algorithm,
// which varies depending on the contributor type.
func checkModule11(contributor contributorType, value string) error {
	var result int
	var digit string

	switch contributor {
	case contributorNatural:
		digit = value[9:10]
		result = calcModule10(value[:9])

	case contributorPrivate:
		digit = value[9:10]
		result = calcPrivateModule11(value[:9])

	case contributorPublic:
		digit = value[8:9]
		result = calcPublicModule11(value[:8])

	default:
		return errUnknownContributorType
	}

	if digit != strconv.Itoa(result) {
		return errInvalidChecksum
	}

	return nil
}
