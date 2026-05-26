package models

type Produto struct {
	ID int `json:"id"`

	CategoriaID   int    `json:"categoria_id"`
	CategoriaNome string `json:"categoria_nome,omitempty"`

	Tipo string `json:"tipo"`

	Nome   string `json:"nome"`
	Marca  string `json:"marca"`
	Modelo string `json:"modelo"`

	Quantidade       int `json:"quantidade"`
	QuantidadeMinima int `json:"quantidade_minima"`

	Setor string `json:"setor"`

	DescricaoTecnica string `json:"descricao_tecnica"`
	Observacao       string `json:"observacao"`

	DataCadastro string `json:"data_cadastro"`
}
