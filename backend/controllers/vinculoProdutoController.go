package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/database"
	"backend/models"
)

func ListarVinculos(c *gin.Context) {

	query := `
		SELECT
			vp.id,

			vp.produto_pai_id,
			pai.nome,

			vp.componente_id,
			comp.nome,

			vp.quantidade,

			vp.data_vinculo

		FROM vinculos_produtos vp

		INNER JOIN produtos pai
			ON pai.id = vp.produto_pai_id

		INNER JOIN produtos comp
			ON comp.id = vp.componente_id
	`

	rows, erro := database.DB.Query(query)

	if erro != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": erro.Error(),
		})

		return
	}

	defer rows.Close()

	vinculos := []models.VinculoProduto{}

	for rows.Next() {

		var vinculo models.VinculoProduto

		err := rows.Scan(

			&vinculo.ID,

			&vinculo.ProdutoPaiID,
			&vinculo.ProdutoPaiNome,

			&vinculo.ComponenteID,
			&vinculo.ComponenteNome,

			&vinculo.Quantidade,

			&vinculo.DataVinculo,
		)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": err.Error(),
			})

			return
		}

		vinculos = append(vinculos, vinculo)
	}

	c.JSON(http.StatusOK, vinculos)
}

func CriarVinculo(c *gin.Context) {

	var vinculo models.VinculoProduto

	// RECEBE JSON
	if erro := c.ShouldBindJSON(&vinculo); erro != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": erro.Error(),
		})

		return
	}

	// VERIFICA ESTOQUE
	var estoqueAtual int

	err := database.DB.QueryRow(`
		SELECT quantidade
		FROM produtos
		WHERE id = ?
	`, vinculo.ComponenteID).Scan(&estoqueAtual)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": "Erro ao consultar estoque",
		})

		return
	}

	// VALIDA ESTOQUE
	if vinculo.Quantidade > estoqueAtual {

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Quantidade maior que estoque disponível",
		})

		return
	}

	// TRANSAÇÃO
	tx, err := database.DB.Begin()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	// INSERE VÍNCULO
	insert, err := tx.Exec(`
		INSERT INTO vinculos_produtos (
			produto_pai_id,
			componente_id,
			quantidade
		)

		VALUES (?, ?, ?)
	`,
		vinculo.ProdutoPaiID,
		vinculo.ComponenteID,
		vinculo.Quantidade,
	)

	if err != nil {

		tx.Rollback()

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	// ATUALIZA ESTOQUE
	_, err = tx.Exec(`
		UPDATE produtos
		SET quantidade = quantidade - ?
		WHERE id = ?
	`,
		vinculo.Quantidade,
		vinculo.ComponenteID,
	)

	if err != nil {

		tx.Rollback()

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	// CONFIRMA
	err = tx.Commit()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	id, _ := insert.LastInsertId()

	c.JSON(http.StatusCreated, gin.H{
		"id":       id,
		"mensagem": "Vínculo criado com sucesso!",
	})
}
func ListarVinculosPorProduto(c *gin.Context) {

	id := c.Param("id")

	query := `
		SELECT
			vp.id,

			vp.produto_pai_id,
			p.nome AS produto_pai_nome,

			vp.componente_id,
			c.nome AS componente_nome,

			vp.quantidade,

			vp.data_vinculo

		FROM vinculos_produtos vp

		INNER JOIN produtos p
			ON p.id = vp.produto_pai_id

		INNER JOIN produtos c
			ON c.id = vp.componente_id

		WHERE vp.produto_pai_id = ?
	`

	rows, err := database.DB.Query(query, id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	defer rows.Close()

	var vinculos []models.VinculoProduto

	for rows.Next() {

		var v models.VinculoProduto

		err := rows.Scan(
			&v.ID,

			&v.ProdutoPaiID,
			&v.ProdutoPaiNome,

			&v.ComponenteID,
			&v.ComponenteNome,

			&v.Quantidade,

			&v.DataVinculo,
		)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": err.Error(),
			})

			return
		}

		vinculos = append(vinculos, v)
	}

	c.JSON(http.StatusOK, vinculos)
}

func BuscarVinculos(c *gin.Context) {

	filtro := c.Query("q")

	query := `
		SELECT
			vp.id,

			vp.produto_pai_id,
			pai.nome,

			vp.componente_id,
			comp.nome,

			vp.quantidade,

			vp.data_vinculo

		FROM vinculos_produtos vp

		INNER JOIN produtos pai
			ON pai.id = vp.produto_pai_id

		INNER JOIN produtos comp
			ON comp.id = vp.componente_id

		WHERE
			pai.tipo = 'EQUIPAMENTO'
			AND (
				pai.nome LIKE ?
				OR vp.produto_pai_id LIKE ?
				OR comp.nome LIKE ?
			)
	`

	rows, err := database.DB.Query(
		query,
		"%"+filtro+"%",
		"%"+filtro+"%",
		"%"+filtro+"%",
	)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	defer rows.Close()

	var vinculos []models.VinculoProduto

	for rows.Next() {

		var v models.VinculoProduto

		err := rows.Scan(
			&v.ID,

			&v.ProdutoPaiID,
			&v.ProdutoPaiNome,

			&v.ComponenteID,
			&v.ComponenteNome,

			&v.Quantidade,

			&v.DataVinculo,
		)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": err.Error(),
			})

			return
		}

		vinculos = append(vinculos, v)
	}

	c.JSON(http.StatusOK, vinculos)
}

func ExcluirVinculo(c *gin.Context) {

	id := c.Param("id")

	tx, err := database.DB.Begin()

	if err != nil {

		c.JSON(500, gin.H{
			"erro": err.Error(),
		})

		return
	}

	// DADOS DO VÍNCULO
	var componenteID int
	var quantidade int

	// BUSCA VÍNCULO
	err = tx.QueryRow(`
		SELECT componente_id, quantidade
		FROM vinculos_produtos
		WHERE id = ?
	`, id).Scan(&componenteID, &quantidade)

	if err != nil {

		tx.Rollback()

		c.JSON(404, gin.H{
			"erro": "Vínculo não encontrado",
		})

		return
	}

	// DEVOLVE AO ESTOQUE
	_, err = tx.Exec(`
		UPDATE produtos
		SET quantidade = quantidade + ?
		WHERE id = ?
	`,
		quantidade,
		componenteID,
	)

	if err != nil {

		tx.Rollback()

		c.JSON(500, gin.H{
			"erro": err.Error(),
		})

		return
	}

	// EXCLUI VÍNCULO
	_, err = tx.Exec(`
		DELETE FROM vinculos_produtos
		WHERE id = ?
	`, id)

	if err != nil {

		tx.Rollback()

		c.JSON(500, gin.H{
			"erro": err.Error(),
		})

		return
	}

	// FINALIZA
	err = tx.Commit()

	if err != nil {

		c.JSON(500, gin.H{
			"erro": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"mensagem": "Vínculo excluído com sucesso",
	})
}

func ListarComponentesQRCode(c *gin.Context) {

	id := c.Param("id")

	rows, err := database.DB.Query(`
		SELECT
			c.id,
			c.nome,
			c.marca,
			c.modelo,
			vp.quantidade
		FROM vinculos_produtos vp

		INNER JOIN produtos c
			ON c.id = vp.componente_id

		WHERE vp.produto_pai_id = ?
		ORDER BY c.nome
	`, id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})

		return
	}

	defer rows.Close()

	var componentes []gin.H

	for rows.Next() {

		var id int
		var nome string
		var marca string
		var modelo string
		var quantidade int

		err := rows.Scan(
			&id,
			&nome,
			&marca,
			&modelo,
			&quantidade,
		)

		if err != nil {
			continue
		}

		componentes = append(componentes, gin.H{
			"id":         id,
			"nome":       nome,
			"marca":      marca,
			"modelo":     modelo,
			"quantidade": quantidade,
		})
	}

	c.JSON(http.StatusOK, componentes)
}
