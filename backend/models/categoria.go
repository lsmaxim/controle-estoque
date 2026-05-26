package models

type Categoria struct {
	ID           int    `json:"id"`
	Nome         string `json:"nome"`
	DataCadastro string `json:"data_cadastro"`
}