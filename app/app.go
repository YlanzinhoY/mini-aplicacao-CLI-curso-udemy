package app

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)


func Gerar() *cli.App {

	flags := []cli.Flag {
		cli.StringFlag{
			Name: "host",
			Value: "localhost",
		},
	}

	app := cli.NewApp()
	app.Name = "Aplicação de linha Comando"
	app.Usage = "Busca IPs e Nome de servidor na internet"

	app.Commands = []cli.Command {
		{
			Name: "ip",
			Usage: "Busca IPS de endereços na internet",
			Flags: flags,
			Action: buscarIps ,
		},
		{
			Name: "servidores",
			Usage: "busca o nome dos servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}



func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidor, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidor {
		fmt.Println(servidor.Host)
	}
}



