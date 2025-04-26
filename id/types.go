package id

// docType represents the type of identification document.
// It is defined as an 8-bit integer.
type docType int8

const (
	// dni represents a standard national ID (Cédula) in Ecuador.
	dni docType = 10

	// ruc represents the Unique Taxpayer Registry (RUC) number in Ecuador.
	ruc docType = 13
)

// contributorType defines the type of entity associated with a RUC.
// It is used to differentiate the rules that apply for validation based on entity type.
type contributorType int

const (
	// contributorNatural represents a natural person (individual).
	contributorNatural contributorType = iota

	// contributorPublic represents a public institution.
	contributorPublic

	// contributorPrivate represents a private company or organization.
	contributorPrivate
)
