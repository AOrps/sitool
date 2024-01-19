package main

import (
	"fmt"
	"os"
	"log"

	
	"github.com/urfave/cli/v2"
)
  



func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Name: "greet",
		Usage: "fight the loneliness! tut",
		Action: func(*cli.Context) error {
			fmt.Println("sup")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
