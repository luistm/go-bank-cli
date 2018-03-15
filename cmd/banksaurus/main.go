package main

import (
	"errors"
	"fmt"
	"os"

	docopt "github.com/docopt/docopt-go"
	"github.com/luistm/banksaurus/cmd/banksaurus/commands"
	"github.com/luistm/banksaurus/cmd/banksaurus/configurations"
)

var intro = "    \n    Your command line finance manager.\n\n"

var usage = `Usage:
	banksaurus -h | --help
	banksaurus report --input <file> [ --grouped ]
	banksaurus load --input <file>
	banksaurus seller change <id> --pretty <name>
	banksaurus seller new <name>
	banksaurus seller show`

var options = `

Options:
	--grouped     The report is present grouped by seller
	--input       The path to the records list.
	--name        Specifies the name.
	-h --help     Show this screen.`

var version = "banksaurus 1.1.0"

func setup() error {
	if configurations.IsDev() {
		return nil
	}

	// Create home dir if not exists
	_, err := os.Stat(configurations.ApplicationHomePath())
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(configurations.ApplicationHomePath(), 0700)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

var errGeneric = errors.New("Error while performing operation")

func main() {
	err := setup()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to setup application")
		os.Exit(2)
	}

	arguments, err := docopt.Parse(intro+usage+options, nil, true, version, false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	command, err := commands.New(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, errGeneric.Error())
		os.Exit(2)
	}

	err = command.Execute(arguments)
	if err != nil {
		fmt.Fprintf(os.Stderr, errGeneric.Error()+"\n")
		os.Exit(2)
	}
}