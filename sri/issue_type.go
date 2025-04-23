package sri

// IssueType define los posibles tipos de emisión para la facturación electrónica
// al generar claves de acceso y facturas electrónicas aprobadas por el SRI.
type IssueType string

const (
	// IssueNormal representa un tipo de emisión normal, utilizado para facturas estándar.
	// Este se identifica con el valor "1".
	IssueNormal IssueType = "1"
)
