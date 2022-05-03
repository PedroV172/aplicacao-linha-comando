package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e Nomes de Servidor na internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca I`PS de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "chama",
					Value: "devbook.com.br",
				},
			},
			Action: BuscarIps,
		},
	}
	return app
}

func BuscarIps(c *cli.Context) {
	host := c.String("chama")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}

}
