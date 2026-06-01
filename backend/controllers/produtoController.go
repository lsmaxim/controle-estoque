package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/database"
	"backend/models"
)

// =======================
// LISTAR PRODUTOS
// =======================
func ListarProdutos(c *gin.Context) {

	query := `
		SELECT
			produtos.id,
			produtos.categoria_id,
			categorias.nome,
			produtos.tipo,
			produtos.nome,
			produtos.marca,
			produtos.modelo,
			produtos.quantidade,
			produtos.quantidade_minima,
			produtos.setor,
			produtos.descricao_tecnica,
			produtos.observacao,
			produtos.data_cadastro
		FROM produtos
		INNER JOIN categorias
			ON categorias.id = produtos.categoria_id
	`

	rows, err := database.DB.Query(query)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	defer rows.Close()

	produtos := []models.Produto{}

	for rows.Next() {

		var produto models.Produto

		err := rows.Scan(
			&produto.ID,
			&produto.CategoriaID,
			&produto.CategoriaNome,
			&produto.Tipo,
			&produto.Nome,
			&produto.Marca,
			&produto.Modelo,
			&produto.Quantidade,
			&produto.QuantidadeMinima,
			&produto.Setor,
			&produto.DescricaoTecnica,
			&produto.Observacao,
			&produto.DataCadastro,
		)

		if err != nil {
			fmt.Println(err)
			continue
		}

		produtos = append(produtos, produto)
	}

	c.JSON(http.StatusOK, produtos)
}

// =======================
// CRIAR PRODUTO
// =======================
func CriarProduto(c *gin.Context) {

	var produto models.Produto

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	result, err := database.DB.Exec(`
		INSERT INTO produtos (
			categoria_id,
			tipo,
			nome,
			marca,
			modelo,
			quantidade,
			quantidade_minima,
			setor,
			descricao_tecnica,
			observacao
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		produto.CategoriaID,
		produto.Tipo,
		produto.Nome,
		produto.Marca,
		produto.Modelo,
		produto.Quantidade,
		produto.QuantidadeMinima,
		produto.Setor,
		produto.DescricaoTecnica,
		produto.Observacao,
	)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	id, _ := result.LastInsertId()

	RegistrarHistorico(
		int(id),
		"Administrador",
		"CADASTRO",
		"Equipamento cadastrado",
	)

	c.JSON(http.StatusCreated, gin.H{
		"id":       id,
		"mensagem": "Produto criado com sucesso!",
	})
}

// =======================
// EXCLUIR PRODUTO
// =======================
func ExcluirProduto(c *gin.Context) {

	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	var total int

	err := database.DB.QueryRow(`
		SELECT COUNT(*)
		FROM vinculos_produtos
		WHERE produto_pai_id = ?
	`, id).Scan(&total)

	if err != nil {

		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao verificar vínculos do produto.",
		})

		return
	}

	if total > 0 {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Este produto possui componentes vinculados. Remova os vínculos antes da exclusão.",
		})

		return
	}

	_, err = database.DB.Exec(`
		DELETE FROM produtos
		WHERE id = ?
	`, id)

	if err != nil {

		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao excluir produto.",
		})

		return
	}

	RegistrarHistorico(
		idInt,
		"Administrador",
		"EXCLUSAO",
		"Produto excluído",
	)

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Produto excluído com sucesso!",
	})
}

// =======================
// ATUALIZAR PRODUTO
// =======================
func AtualizarProduto(c *gin.Context) {

	id := c.Param("id")

	var produto models.Produto

	if err := c.ShouldBindJSON(&produto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	result, err := database.DB.Exec(`
		UPDATE produtos SET
			categoria_id = ?,
			tipo = ?,
			nome = ?,
			marca = ?,
			modelo = ?,
			quantidade = ?,
			quantidade_minima = ?,
			setor = ?,
			descricao_tecnica = ?,
			observacao = ?
		WHERE id = ?
	`,
		produto.CategoriaID,
		produto.Tipo,
		produto.Nome,
		produto.Marca,
		produto.Modelo,
		produto.Quantidade,
		produto.QuantidadeMinima,
		produto.Setor,
		produto.DescricaoTecnica,
		produto.Observacao,
		id,
	)

	if err != nil {
		fmt.Println("ERRO UPDATE:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"erro": "Produto não encontrado",
		})
		return
	}

	idInt, _ := strconv.Atoi(id)

	RegistrarHistorico(
		idInt,
		"Administrador",
		"EDICAO",
		"Produto atualizado",
	)

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Produto atualizado com sucesso!",
	})
}
