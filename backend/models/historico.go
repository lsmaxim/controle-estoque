package models

type Historico struct {
	ID           int    `json:"id"`
	ProdutoID    int    `json:"produto_id"`
	Usuario      string `json:"usuario"`
	Acao         string `json:"acao"`
	Descricao    string `json:"descricao"`
	DataRegistro string `json:"data_registro"`
}
