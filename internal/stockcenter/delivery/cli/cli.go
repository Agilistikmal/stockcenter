package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type CliHandler struct {
	app *cli.App
}

func New() *CliHandler {
	app := &cli.App{
		Name: "stockcenter",
	}

	return &CliHandler{
		app: app,
	}
}

func (h *CliHandler) Run() {
	err := h.app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *CliHandler) Client() {
	h.app.Commands = append(h.app.Commands, &cli.Command{
		Name:    "client",
		Aliases: []string{"c", "cli"},
		Usage:   "Start command line client",
		Action: func(ctx *cli.Context) error {
			menu := Scan("Select Menu")
			fmt.Println("Selected", menu)
			return nil
		},
	})
}
