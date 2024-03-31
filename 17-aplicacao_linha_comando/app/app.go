package app

import "github.com/urfave/cli"

// retorna a aplicacao de linha de comando pronta pra ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de Servidores na internet"

	return app
}
