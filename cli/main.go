package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var language string
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Value:       "english",
				Usage:       "language for the greeting",
				Destination: &language, //此处的变量只能在函数内被取到值
			},

			&cli.StringFlag{
				Name:        "lang1",
				Value:       "English",
				Usage:       "language for the greeting",
				Destination: &language, //此处的变量只能在函数内被取到值
			},
		},
		Action: func(cCtx *cli.Context) error {
			name := "someone"
			if cCtx.NArg() > 0 { //统计参数的个数
				name = cCtx.Args().Get(0)
			}
			if language == "s" { //判定选项值，如果为spanish则走Hola分支
				fmt.Println("Hola", name)
			}
			if language == "English" {
				fmt.Println("Hello1", name)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
func test() {

}
