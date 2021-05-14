package clientesdto

type ActualizarClienteDto struct {
	Uid                   string `json:"uid"`
	Cedula                string `json:"cedula"`
	IdTipoDocumento       int    `json:"id_tipo_documento"`
	Nombre                string `json:"nombre"`
	Apellido              string `json:"apellido"`
	Fullname              string `json:"fullname"`
	Email                 string `json:"email"`
	Celular               string `json:"celular"`
	Genero                string `json:"genero"`
	FechaRegistroLeal     string `json:"fecha_registro_leal"`
	Cumpleanos            string `json:"cumpleanos"`
	CumpleanosMes         int    `json:"cumpleanos_mes"`
	CumpleanosDia         int    `json:"cumpleanos_dia"`
	GrupoEdad             string `json:"grupo_edad"`
	Edad                  int    `json:"edad"`
	EstadoLeal            string `json:"estado_leal"`
	PlataformaMovil       string `json:"plataforma_movil"`
	EnvioMailGlobal       int    `json:"envio_mail_global"`
	EnvioSmsGlobal        int    `json:"envio_sms_global"`
	EnvioPushGlobal       int    `json:"envio_push_global"`
	FechaPerfilCompletado string `json:"fecha_perfil_completado"`
	IdCiudad              int    `json:"id_ciudad"`
	CodPais               string `json:"cod_pais"`
	EstadoUsuario         string `json:"estadoUsuario"`
	IdsComercios          []int  `json:"ids_comercios"`
}
