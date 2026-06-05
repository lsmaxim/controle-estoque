package routes

import (
	"github.com/gin-gonic/gin"

	"backend/controllers"
	"backend/middlewares"

)

func ConfigurarRotas(r *gin.Engine) {

	// LOGIN
	r.POST("/login", controllers.Login)

	// ROTAS PROTEGIDAS
	api := r.Group("/")
	api.Use(middlewares.AuthMiddleware())
	{

		api.GET(
			"/relatorios/estoque",
			controllers.GerarRelatorioEstoque,
		)

		api.GET(
			"/produtos/:id/historico",
			controllers.ListarHistorico,
		)
		// CATEGORIAS
		api.GET("/categorias", controllers.ListarCategorias)
		api.POST("/categorias", controllers.CriarCategoria)

		// PRODUTOS
		api.GET("/produtos", controllers.ListarProdutos)
		api.POST("/produtos", controllers.CriarProduto)
		api.PUT("/produtos/:id", controllers.AtualizarProduto)
		api.DELETE("/produtos/:id", controllers.ExcluirProduto)
		
		// VINCULOS
		api.GET("/vinculos", controllers.ListarVinculos)
		api.POST("/vinculos", controllers.CriarVinculo)
		api.DELETE("/vinculos/:id", controllers.ExcluirVinculo)

		api.GET(
			"/produtos/:id/vinculos",
			controllers.ListarVinculosPorProduto,
		)

		// AGENDAMENTOS
		api.GET("/agendamentos", controllers.ListarAgendamentos)
		api.POST("/agendamentos", controllers.CriarAgendamento)

		api.PUT(
			"/agendamentos/:id/finalizar",
			controllers.FinalizarAgendamento,
		)

		// MOVIMENTAÇÕES
		api.GET("/movimentacoes", controllers.ListarMovimentacoes)
		api.POST("/movimentacoes", controllers.CriarMovimentacao)
	}

	r.GET("/equipamentos/:id/qrcode", controllers.GerarQRCode)

	r.GET(
		"/equipamentos/:id/etiqueta",
		controllers.GerarEtiqueta,
	)

	r.GET(
		"/qrcode/equipamento/:id",
		controllers.BuscarProdutoQRCode,
	)
}
