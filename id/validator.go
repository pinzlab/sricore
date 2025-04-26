package id

import "strconv"

// IsDNI verifica si un número de cédula ecuatoriana es válido.
//
// Comprueba el formato, el código de provincia, el tipo de contribuyente
// y el dígito verificador usando el algoritmo Módulo 10.
func IsDNI(value string) error {
	if err := isDigit(dni, value); err != nil {
		return err
	}

	if err := checkProvince(value[:2]); err != nil {
		return err
	}

	if err := checkContributor(contributorNatural, value[2:3]); err != nil {
		return err
	}

	if err := checkModule10(value); err != nil {
		return err
	}

	return nil
}

// IsNaturalRUC verifica si un RUC correspondiente a una persona natural es válido.
//
// Realiza validaciones sobre el formato, el código de provincia,
// el tipo de contribuyente, el código del establecimiento y el dígito verificador
// usando el algoritmo Módulo 11 para personas naturales.
func IsNaturalRUC(value string) error {
	if err := isDigit(ruc, value); err != nil {
		return err
	}

	if err := checkProvince(value[:2]); err != nil {
		return err
	}

	if err := checkContributor(contributorNatural, value[2:3]); err != nil {
		return err
	}

	if err := checkEstablishment(value[10:]); err != nil {
		return err
	}

	if err := checkModule11(contributorNatural, value); err != nil {
		return err
	}

	return nil
}

// IsPrivateRUC verifica si un RUC correspondiente a una empresa privada es válido.
//
// Valida el formato, el código de provincia, el tipo de contribuyente,
// el código del establecimiento y el dígito verificador usando el algoritmo Módulo 11
// específico para privados.
func IsPrivateRUC(value string) error {
	if err := isDigit(ruc, value); err != nil {
		return err
	}

	if err := checkProvince(value[:2]); err != nil {
		return err
	}

	if err := checkContributor(contributorPrivate, value[2:3]); err != nil {
		return err
	}

	if err := checkEstablishment(value[10:]); err != nil {
		return err
	}

	if err := checkModule11(contributorPrivate, value); err != nil {
		return err
	}

	return nil
}

// IsPublicRUC verifica si un RUC correspondiente a una entidad pública es válido.
//
// Realiza validaciones sobre el formato, el código de provincia, el tipo de contribuyente,
// el establecimiento y el dígito verificador usando el algoritmo Módulo 11 para entidades públicas.
func IsPublicRUC(value string) error {
	if err := isDigit(ruc, value); err != nil {
		return err
	}

	if err := checkProvince(value[:2]); err != nil {
		return err
	}

	if err := checkContributor(contributorPublic, value[2:3]); err != nil {
		return err
	}

	if err := checkEstablishment(value[10:]); err != nil {
		return err
	}

	if err := checkModule11(contributorPublic, value); err != nil {
		return err
	}

	return nil
}

// IsRUC verifica si un RUC ecuatoriano es válido, determinando automáticamente
// el tipo de contribuyente en base al tercer dígito del número.
//
// Aplica validaciones comunes a todos los tipos de RUC, como el formato,
// la provincia, el código del establecimiento y el dígito verificador correspondiente.
func IsRUC(value string) error {
	if err := isDigit(ruc, value); err != nil {
		return err
	}

	if err := checkProvince(value[:2]); err != nil {
		return err
	}

	var contributor contributorType
	contributorCode, _ := strconv.Atoi(value[2:3])

	switch {
	case contributorCode >= 0 && contributorCode <= 5:
		contributor = contributorNatural
	case contributorCode == 6:
		contributor = contributorPublic
	case contributorCode == 9:
		contributor = contributorPrivate
	default:
		return errUnknownContributorType
	}

	if err := checkContributor(contributor, value[2:3]); err != nil {
		return err
	}

	if err := checkModule11(contributor, value); err != nil {
		return err
	}

	if err := checkEstablishment(value[10:]); err != nil {
		return err
	}

	return nil
}
