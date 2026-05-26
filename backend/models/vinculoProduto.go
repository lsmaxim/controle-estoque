package models

type VinculoProduto struct {
	ID int `json:"id"`

	ProdutoPaiID int `json:"produto_pai_id"`

	ProdutoPaiNome string `json:"produto_pai_nome"`

	ComponenteID int `json:"componente_id"`

	ComponenteNome string `json:"componente_nome"`

	Quantidade int `json:"quantidade"`

	DataVinculo string `json:"data_vinculo"`
}
