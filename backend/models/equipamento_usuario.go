package models

import "time"

type EquipamentoUsuario struct {
	ID        int `json:"id"`
	ProdutoID int `json:"produto_id"`
	UsuarioID int `json:"usuario_id"`

	ProdutoNome string `json:"produto_nome,omitempty"`
	UsuarioNome string `json:"usuario_nome,omitempty"`

	DataVinculo time.Time `json:"data_vinculo"`
}
