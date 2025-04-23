package sri

// VoucherType es un tipo personalizado que representa los distintos tipos de comprobantes
// válidos en el sistema de facturación electrónica del SRI (Servicio de Rentas Internas).
type VoucherType string

const (
	// Invoice representa el comprobante tipo factura, identificado con el código "01".
	// Utilizado para documentar la venta de bienes o prestación de servicios.
	Invoice VoucherType = "01"

	// Purchase representa el comprobante de compras, identificado con el código "03".
	// Utilizado para registrar adquisiciones de bienes y servicios realizadas por el contribuyente.
	Purchase VoucherType = "03"

	// CreditNote representa la nota de crédito, identificada con el código "04".
	// Utilizada para corregir o anular facturas previamente emitidas.
	CreditNote VoucherType = "04"

	// DebitNote representa la nota de débito, identificada con el código "05".
	// Utilizada para modificar o incrementar valores en facturas ya emitidas.
	DebitNote VoucherType = "05"

	// Delivery representa la guía de remisión, identificada con el código "06".
	// Utilizada para respaldar el traslado o transporte de bienes.
	Delivery VoucherType = "06"

	// Retention representa el comprobante de retención, identificado con el código "07".
	// Utilizado para documentar la retención de impuestos aplicadas a terceros.
	Retention VoucherType = "07"
)
