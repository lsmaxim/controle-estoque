package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Conectar() {

	usuario := "root"
	senha := "#@control@#"
	banco := "controle_estoque"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		usuario,
		senha,
		banco,
	)

	conexao, erro := sql.Open(
		"mysql",
		dsn,
	)

	if erro != nil {
		panic(erro)
	}

	erro = conexao.Ping()

	if erro != nil {
		panic(erro)
	}

	DB = conexao

	fmt.Println("Banco conectado com sucesso!")
}
