package clientesdto

type NuevaActualizacionDto struct {
	IdHistorial               int      `json:"id_historial"`
	Uid                       string   `json:"uid"`
	Fecha                     string   `json:"fecha"`
	FechaVencimiento          string   `json:"fecha_vencimiento"`
	Puntos                    int      `json:"puntos"`
	BalancePuntos             int      `json:"balance_puntos"`
	IdEstadoPuntos            int      `json:"id_estado_puntos"`
	Tipo                      string   `json:"tipo"`
	IdPremio                  int      `json:"id_premio"`
	UidCms                    string   `json:"uid_cms"`
	AccionRed                 string   `json:"accion_red"`
	Valor                     float64  `json:"valor"`
	IdComercio                int      `json:"id_comercio"`
	IdSucursal                int      `json:"id_sucursal"`
	IpCajero                  string   `json:"ip_cajero"`
	Factura                   string   `json:"factura"`
	Nota                      string   `json:"nota"`
	IdFranquicia              int      `json:"id_franquicia"`
	IdTipoTransaccion         int      `json:"id_tipo_transaccion"`
	CodigoBin                 string   `json:"codigo_bin"`
	Frecuencia                float64  `json:"frecuencia"`
	TicketPromedio            float64  `json:"ticket_promedio"`
	GastoTotal                float64  `json:"gasto_total"`
	CantidadPremiosReclamados int      `json:"premios_reclamados_cantidad"`
	Transacciones             int      `json:"transacciones"`
	Banderazos                int      `json:"banderazos"`
	Cargas                    int      `json:"cargas"`
	Redenciones               int      `json:"redenciones"`
	Visitas                   int      `json:"visitas"`
	UltimaVisita              string   `json:"ultima_visita"`
	Etiquetas                 int      `json:"etiquetas"`
	Productos                 []string `json:"codigo_item"`
	Categorias                []int    `json:"categorias"`
	PremiosReclamados         []string `json:"premios_reclamdos"`
	IdPremiosReclamados       []int    `json:"id_premios_reclamados"`
	IdSucursales              []int    `json:"id_sucursales"`
}
