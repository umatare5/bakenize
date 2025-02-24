// Description: A simple CLI to generate a token for the Authorization Header from a username and password.
package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "bakenize",
		Usage:     "Generate a token for the Authorization header from a username and password",
		UsageText: "bakenize",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "username",
				Aliases:  []string{"u"},
				Usage:    "Set a username",
				EnvVars:  []string{"B_USERNAME"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "password",
				Aliases:  []string{"p"},
				Usage:    "Set a password",
				EnvVars:  []string{"B_PASSWORD"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			username := c.String("username")
			password := c.String("password")
			creds := fmt.Sprintf("%s:%s", username, password)
			fmt.Println(base64.StdEncoding.EncodeToString([]byte(creds)))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
