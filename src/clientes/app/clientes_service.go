package clientesapp

import (
	clientesadapters "leal.co/clientes-comercio/src/clientes/adapters"
	clientes "leal.co/clientes-comercio/src/clientes/domain"
	clientesdto "leal.co/clientes-comercio/src/clientes/domain/dtos"
	"log"
)

type ClientesService interface {
	ActualizarCliente(clientesdto.ActualizarClienteDto)
	CargarPuntos(clientesdto.CarguePuntosDto)
	ActualizarUsuarioComercio(clientesdto.NuevaActualizacionDto)
}

type clientesService struct {
	repository clientesadapters.ClientesRepository
}

func NewClientesService() ClientesService {
	return &clientesService{
		repository: clientesadapters.NewClientesRepository(),
	}
}

func (c *clientesService) ActualizarCliente(dto clientesdto.ActualizarClienteDto) {
	log.Println(dto)
	cliente := clientes.FromActualizarCliente(dto)
	log.Println(cliente)
	if err := c.repository.SaveCliente(cliente); err != nil {
		log.Printf("error inserting cliente i: %v", err)
	}
}

func (c *clientesService) CargarPuntos(dto clientesdto.CarguePuntosDto) {
	cliente := clientes.FromCarguePuntos(dto)
	if err := c.repository.SaveCliente(cliente); err != nil {
		log.Printf("error inserting cliente i: %v", err)
	}
}

func (c *clientesService) ActualizarUsuarioComercio(dto clientesdto.NuevaActualizacionDto) {
	cliente := clientes.FromActualizarClienteComercio(dto)
	if err := c.repository.SaveCliente(cliente); err != nil {
		log.Printf("error inserting cliente i: %v", err)
	}
}
