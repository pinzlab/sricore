package sri

// EnvType define los posibles entornos para invocar los servicios web
// publicados por la Administración Tributaria (SRI).
type EnvType string

const (
	// EnvTest representa el entorno de pruebas, identificado por el código "1".
	// Se utiliza para fines de prueba y validación.
	EnvTest EnvType = "1"

	// EnvProd representa el entorno de producción, identificado por el código "2".
	// Se utiliza para transacciones reales en el entorno de producción.
	EnvProd EnvType = "2"
)
