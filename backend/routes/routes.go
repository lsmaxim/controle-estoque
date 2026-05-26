package routes

import (
	"github.com/gin-gonic/gin"

	"backend/controllers"
)

func ConfigurarRotas(r *gin.Engine) {

	// CATEGORIAS
	r.GET("/categorias", controllers.ListarCategorias)
	r.POST("/categorias", controllers.CriarCategoria)

	// PRODUTOS
	r.GET("/produtos", controllers.ListarProdutos)
	r.POST("/produtos", controllers.CriarProduto)
	r.PUT("/produtos/:id", controllers.AtualizarProduto)
	r.DELETE("/produtos/:id", controllers.ExcluirProduto)

	// VÍNCULOS
	r.GET("/vinculos", controllers.ListarVinculos)
	r.POST("/vinculos", controllers.CriarVinculo)
	r.DELETE("/vinculos/:id", controllers.ExcluirVinculo)

	r.GET("/produtos/:id/vinculos",
		controllers.ListarVinculosPorProduto)

	// AGENDAMENTOS
	r.GET("/agendamentos", controllers.ListarAgendamentos)
	r.POST("/agendamentos", controllers.CriarAgendamento)
	r.PUT("/agendamentos/:id/finalizar",
		controllers.FinalizarAgendamento)

	// MOVIMENTAÇÕES
	r.GET("/movimentacoes",
		controllers.ListarMovimentacoes)

	r.POST("/movimentacoes",
		controllers.CriarMovimentacao)
}
