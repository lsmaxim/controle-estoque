package models

import "time"

type Chamado struct {
	ID              int        `json:"id"`
	Titulo          string     `json:"titulo"`
	Descricao       string     `json:"descricao"`
	Solicitante     string     `json:"solicitante"`
	Setor           string     `json:"setor"`
	EquipamentoID   *int       `json:"equipamento_id"`
	EquipamentoNome string     `json:"equipamento_nome"`
	Status          string     `json:"status"`
	Prioridade      string     `json:"prioridade"`
	Solucao         string     `json:"solucao"`
	DataAbertura    time.Time  `json:"data_abertura"`
	DataFechamento  *time.Time `json:"data_fechamento"`
}
