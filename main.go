package main

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/fundary/bitauth"
	"github.com/fundary/bitpay/client"
)

func PanicIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Bitpay"
	app.Usage = "fight the loneliness!"
	app.Action = cli.ShowAppHelp
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "env",
			Value: "prod",
			Usage: "Which Bitpay environment that should be used, production or test",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "generate",
			ShortName: "g",
			Usage:     "Generate a new public/private key pair and a client id",
			Action:    GenerateKeysAndSIN,
		},
		{
			Name:      "new-token",
			ShortName: "n",
			Usage:     "Generate new token from Bitpay API using a generated client ID, requires 3 arguments (label, clientID, facade)",
			Action:    NewToken,
			Flags:     flags,
		},
		{
			Name:      "claim-token",
			ShortName: "c",
			Usage:     "Claim a token by pairing code, requires 3 arguments (label, clientID, pairing code)",
			Action:    ClaimToken,
			Flags:     flags,
		},
	}

	err := app.Run(os.Args)
	PanicIf(err)
}

func GenerateKeysAndSIN(c *cli.Context) {
	sin, err := bitauth.GenerateSIN()
	PanicIf(err)
	println("Public key: " + hex.EncodeToString(sin.PublicKey))
	println("Private key: " + hex.EncodeToString(sin.PrivateKey))
	println("Client ID: ", string(sin.SIN))
}

func NewToken(c *cli.Context) {
	if len(c.Args()) < 3 {
		println("Requires 3 arguments, see usage")

		return
	}

	var url string
	if c.String("env") == "prod" {
		url = client.APIBaseProd
	} else {
		url = client.APIBaseTest
	}

	bitpay := client.NewClient(url)

	tokenResp, err := bitpay.NewToken(c.Args()[0], c.Args()[1], client.Facade(c.Args()[2]))
	PanicIf(err)

	json, err := json.MarshalIndent(&tokenResp, "", "	")
	PanicIf(err)

	println(string(json))
}

func ClaimToken(c *cli.Context) {
	if len(c.Args()) < 3 {
		println("Requires 3 arguments, see usage")

		return
	}

	var url string
	if c.String("env") == "prod" {
		url = client.APIBaseProd
	} else {
		url = client.APIBaseTest
	}

	bitpay := client.NewClient(url)

	tokenResp, err := bitpay.ClaimToken(c.Args()[0], c.Args()[1], c.Args()[2])
	PanicIf(err)

	json, err := json.MarshalIndent(&tokenResp, "", "	")
	PanicIf(err)

	println(string(json))
}
