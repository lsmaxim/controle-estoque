package models

type Movimentacao struct {
	ID int `json:"id"`

	ProdutoID int `json:"produto_id"`

	Tipo string `json:"tipo"`

	Quantidade int `json:"quantidade"`

	Observacao string `json:"observacao"`

	DataMovimentacao string `json:"data_movimentacao"`

	UsuarioEmail string `json:"usuario_email"`
}
