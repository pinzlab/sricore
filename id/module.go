package id

import (
	"strconv"
	"strings"
)

// calcModule10 calculates the verification digit using the Modulo 10 algorithm.
// This is used for validating Ecuadorian cédulas and RUCs of natural persons.
func calcModule10(value string) int {
	var checksum int
	digits := strings.Split(value[:9], "")

	for index, item := range digits {
		digit, _ := strconv.Atoi(item)

		// Multiply digits in odd positions by 2, and subtract 9 if result is greater than 9
		if index%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		checksum += digit
	}

	residue := checksum % 10
	if residue == 0 {
		return 0
	}
	return 10 - residue
}

// calcModule11 is a general-purpose function to calculate a verification digit
// using the Modulo 11 algorithm with custom coefficients.
// It is used as the base function for validating private and public RUCs.
func calcModule11(digits []string, coefficients []int) int {
	var checksum int
	for i, coef := range coefficients {
		digit, _ := strconv.Atoi(digits[i])
		checksum += digit * coef
	}

	residue := checksum % 11
	if residue == 0 {
		return 0
	}
	return 11 - residue
}

// calcPublicModule11 applies the Modulo 11 algorithm with specific coefficients
// for validating public sector RUCs in Ecuador.
func calcPublicModule11(value string) int {
	digits := strings.Split(value, "")
	coefficients := []int{3, 2, 7, 6, 5, 4, 3, 2}
	return calcModule11(digits, coefficients)
}

// calcPrivateModule11 applies the Modulo 11 algorithm with specific coefficients
// for validating private company RUCs in Ecuador.
func calcPrivateModule11(value string) int {
	digits := strings.Split(value, "")
	coefficients := []int{4, 3, 2, 7, 6, 5, 4, 3, 2}
	return calcModule11(digits, coefficients)
}
