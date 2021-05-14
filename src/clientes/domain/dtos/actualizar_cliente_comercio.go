package clientesdto

type CarguePuntosDto struct {
	Uid                        string                `json:"uid"`
	IdComercio                 int                   `json:"id_comercio"`
	TipoDocumento              string                `json:"tipo_documento"`
	FechaRegistroComercio      string                `json:"fecha_registro_comercio"`
	EstadoComercio             string                `json:"estado_comercio"`
	EnvioMail                  int                   `json:"envio_mail"`
	EnvioSms                   int                   `json:"envio_sms"`
	EnvioPushComercio          int                   `json:"envio_push_comercio"`
	PuntosActivos              int                   `json:"puntos_activos"`
	PuntosVencidos             int                   `json:"puntos_vencidos"`
	PuntosUsados               int                   `json:"puntos_usados"`
	PuntosTotales              int                   `json:"puntos_totales"`
	PuntosPorVencer            int                   `json:"puntos_por_vencer"`
	IdEstatus                  int                   `json:"id_estatus"`
	Estatus                    string                `json:"estatus"`
	UltimaCalificacion         int                   `json:"ultima_calificacion"`
	IdPremiosDisponibles       string                `json:"id_premios_disponibles"`
	PremiosDisponibles         string                `json:"premios_disponibles"`
	PremiosPorVencer           int                   `json:"premios_por_vencer"`
	PremiosDisponiblesCantidad int                   `json:"premios_disponibles_cantidad"`
	BlackList                  string                `json:"black_list"`
	IdDepartamento             int                   `json:"id_departamento"`
	Ciudad                     string                `json:"ciudad"`
	Usuario                    usuusuarios.DtoSalida `json:"usuario"`
}
