package models

type Agendamento struct {
	ID int `json:"id"`

	Titulo string `json:"titulo"`

	DataAgendamento string `json:"data_agendamento"`

	Status string `json:"status"`

	Observacao string `json:"observacao"`
}
