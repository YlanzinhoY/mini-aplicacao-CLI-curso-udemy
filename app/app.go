package app

import "github.com/urfave/cli"

func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação de linha Comando"
	app.Usage = "Busca IPs e Nome de servidor na internet"

	app.Commands = []cli.Command {
		Name: "ip",
		Usage: "Buscas IPS de endereços na internet",
		Flags: []cli.Flag {
			cli.StringFlag{
				Name: "host",
				Value:  "github.com/YlanzinhoY" ,
			}
		}
	}

	return app
}
