package clientes

import (
	clientesdto "leal.co/clientes-comercio/src/clientes/domain/dtos"
)

type Cliente struct {
	Uid                        string      `bson:"uid,omitempty"`
	IdComercio                 int         `bson:"id_comercio,omitempty"`
	Cedula                     string      `bson:"cedula,omitempty"`
	IdTipoDocumento            int         `bson:"id_tipo_documento,omitempty"`
	TipoDocumento              string      `bson:"tipo_documento,omitempty"`
	Nombre                     string      `bson:"nombre,omitempty"`
	Apellido                   string      `bson:"apellido,omitempty"`
	Fullname                   string      `bson:"fullname,omitempty"`
	Email                      string      `bson:"email,omitempty"`
	Celular                    string      `bson:"celular,omitempty"`
	Genero                     string      `bson:"genero,omitempty"`
	FechaRegistroLeal          string      `bson:"fecha_registro_leal,omitempty"`
	FechaRegistroComercio      string      `bson:"fecha_registro_comercio,omitempty"`
	Cumpleanos                 string      `bson:"cumpleanos,omitempty"`
	CumpleanosMes              int         `bson:"cumpleanos_mes,omitempty"`
	CumpleanosDia              int         `bson:"cumpleanos_dia,omitempty"`
	GrupoEdad                  string      `bson:"grupo_edad,omitempty"`
	Edad                       int         `bson:"edad,omitempty"`
	EstadoLeal                 string      `bson:"estado_leal,omitempty"`
	EstadoComercio             string      `bson:"estado_comercio,omitempty"`
	PlataformaMovil            interface{} `bson:"plataforma_movil,omitempty"`
	EnvioMailGlobal            int         `bson:"envio_mail_global,omitempty"`
	EnvioMail                  int         `bson:"envio_mail,omitempty"`
	EnvioSMSGlobal             int         `bson:"envio_sms_global,omitempty"`
	EnvioSMS                   int         `bson:"envio_sms,omitempty"`
	EnvioPushGlobal            int         `bson:"envio_push_global,omitempty"`
	EnvioPushComercio          int         `bson:"envio_push_comercio,omitempty"`
	FechaPerfilCompletado      string      `bson:"fecha_perfil_completado,omitempty"`
	IdCiudad                   int         `bson:"id_ciudad,omitempty"`
	Ciudad                     string      `bson:"ciudad,omitempty"`
	CodPais                    string      `bson:"cod_pais,omitempty"`
	IDDepartamento             int         `bson:"id_departamento,omitempty"`
	PuntosActivos              int         `bson:"puntos_activos,omitempty"`
	PuntosVencidos             int         `bson:"puntos_vencidos,omitempty"`
	PuntosUsados               int         `bson:"puntos_usados,omitempty"`
	PuntosTotales              int         `bson:"puntos_totales,omitempty"`
	PuntosPorVencer            int         `bson:"puntos_por_vencer,omitempty"`
	IDEstatus                  int         `bson:"id_estatus,omitempty"`
	Estatus                    string      `bson:"estatus,omitempty"`
	UltimaCalificacion         int         `bson:"ultima_calificacion,omitempty"`
	PremiosDisponiblesCantidad int         `bson:"premios_disponibles_cantidad,omitempty"`
	IDPremiosDisponibles       string      `bson:"id_premios_disponibles,omitempty"`
	PremiosDisponibles         string      `bson:"premios_disponibles,omitempty"`
	BlackList                  string      `bson:"black_list,omitempty"`
	PremiosPorVencer           int         `bson:"premios_por_vencer,omitempty"`
	EstadoUsuario              string      `bson:"estadoUsuario,omitempty"`
	TotalTx                    int         `bson:"totalTx,omitempty"`
	TicketPromedio             float64     `bson:"ticket_promedio,omitempty"`
	Frecuencia                 float64     `bson:"frecuencia,omitempty"`
	GastoTotal                 float64     `bson:"gasto_total,omitempty"`
	IDSucursal                 []int       `bson:"id_sucursal,omitempty"`
	IDPremioReclamados         interface{} `bson:"id_premio_reclamados,omitempty"`
	PremiosReclamados          interface{} `bson:"premios_reclamados,omitempty"`
	PremiosReclamadosCantidad  int         `bson:"premios_reclamados_cantidad,omitempty"`
	Transacciones              int         `bson:"transacciones,omitempty"`
	Banderazos                 int         `bson:"banderazos,omitempty"`
	Cargas                     int         `bson:"cargas,omitempty"`
	Redenciones                int         `bson:"redenciones,omitempty"`
	Visitas                    int         `bson:"visitas,omitempty"`
	UltimaVisita               string      `bson:"ultima_visita,omitempty"`
	UltimaVisitaDias           int         `bson:"ultima_visita_dias,omitempty"`
	UltimaVisitaHoras          int         `bson:"ultima_visita_horas,omitempty"`
	FechaUltimaCalificacion    interface{} `bson:"fecha_ultima_calificacion,omitempty"`
	Productos                  []string    `bson:"productos,omitempty"`
	Etiquetas                  int         `bson:"etiquetas,omitempty"`
	Categorias                 []int       `bson:"categorias,omitempty"`
	IdsComercios               []int       `bson:"ids_comercios, omitempty"`
}

func FromCarguePuntos(dto clientesdto.CarguePuntosDto) Cliente {
	return Cliente{
		Uid:               dto.Uid,
		Frecuencia:        dto.Frecuencia,
		TicketPromedio:    dto.TicketPromedio,
		GastoTotal:        dto.GastoTotal,
		Transacciones:     dto.Transacciones,
		Banderazos:        dto.Banderazos,
		Cargas:            dto.Cargas,
		Redenciones:       dto.Redenciones,
		Visitas:           dto.Visitas,
		UltimaVisita:      dto.UltimaVisita,
		Etiquetas:         dto.Etiquetas,
		Productos:         dto.Productos,
		Categorias:        dto.Categorias,
		PremiosReclamados: dto.PremiosReclamados,
	}
}

func FromActualizarClienteComercio(dto clientesdto.NuevaActualizacionDto) Cliente {
	return Cliente{
		Uid:                        dto.Uid,
		IdComercio:                 dto.IdComercio,
		TipoDocumento:              dto.TipoDocumento,
		FechaRegistroComercio:      dto.FechaRegistroComercio,
		EstadoComercio:             dto.EstadoComercio,
		EnvioMail:                  dto.EnvioMail,
		EnvioSMS:                   dto.EnvioSMS,
		EnvioPushComercio:          dto.EnvioPushComercio,
		PuntosActivos:              dto.PuntosActivos,
		PuntosVencidos:             dto.PuntosVencidos,
		PuntosUsados:               dto.PuntosUsados,
		PuntosTotales:              dto.PuntosTotales,
		PuntosPorVencer:            dto.PuntosPorVencer,
		IDEstatus:                  dto.IDEstatus,
		Estatus:                    dto.Estatus,
		UltimaCalificacion:         dto.UltimaCalificacion,
		IDPremiosDisponibles:       dto.IDPremiosDisponibles,
		PremiosDisponibles:         dto.PremiosDisponibles,
		PremiosPorVencer:           dto.PremiosPorVencer,
		PremiosDisponiblesCantidad: dto.PremiosDisponiblesCantidad,
		BlackList:                  dto.BlackList,
		IDDepartamento:             dto.IDDepartamento,
		Ciudad:                     dto.Ciudad,
		// usuario??
	}

}

func FromActualizarCliente(dto clientesdto.ActualizarClienteDto) Cliente {
	return Cliente{
		Uid:                   dto.Uid,
		Cedula:                dto.Cedula,
		IdTipoDocumento:       dto.IdTipoDocumento,
		Nombre:                dto.Nombre,
		Apellido:              dto.Apellido,
		Fullname:              dto.Fullname,
		Email:                 dto.Email,
		Celular:               dto.Celular,
		Genero:                dto.Genero,
		FechaRegistroLeal:     dto.FechaRegistroLeal,
		Cumpleanos:            dto.Cumpleanos,
		CumpleanosMes:         dto.CumpleanosMes,
		CumpleanosDia:         dto.CumpleanosDia,
		GrupoEdad:             dto.GrupoEdad,
		Edad:                  dto.Edad,
		EstadoLeal:            dto.EstadoLeal,
		PlataformaMovil:       dto.PlataformaMovil,
		EnvioMailGlobal:       dto.EnvioMailGlobal,
		EnvioSMSGlobal:        dto.EnvioSmsGlobal,
		EnvioPushGlobal:       dto.EnvioPushGlobal,
		FechaPerfilCompletado: dto.FechaPerfilCompletado,
		IdCiudad:              dto.IdCiudad,
		CodPais:               dto.CodPais,
		EstadoUsuario:         dto.EstadoUsuario,
		IdsComercios:          dto.IdsComercios,
	}
}
