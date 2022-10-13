package main

import (
	"flag"
	"os"

	marvin "github.com/tompston/gomarvin/marvin"
)

// go run main.go -h
//
//go:generate go run main.go
func main() {

	args := os.Args[1:]

	flags := &marvin.CmdArgs{}
	flag.StringVar(&flags.DangerousRegen, "dangerous-regen", "false", "Regenerate everything. If set to true, init server will be regenerated and  all previous changes will be lost")
	flag.StringVar(&flags.FetchOnly, "fetch-only", "false", "generate only the typescript file that holds fetch function")
	flag.StringVar(&flags.ConfigPath, "config", marvin.ConfigName, "Specify path to the gomarvin config file")
	flag.Parse()

	marvin.Run(flags, args)
}
