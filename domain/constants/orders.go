package constants

const (
	OrdersPath         = "/orders"
	OrderWithIDPath    = "/orders/:id"
	DefaultOrderStatus = "creado"
	CancelOrderStatus  = "cancelado"
)

var InvalidCancelStatus = []string{"en_ruta", "entregado"}
