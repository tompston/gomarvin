package main

import (
	"flag"

	marvin "github.com/tompston/gomarvin/marvin"
)

// go run main.go -h

func main() {

	cmd := &marvin.CmdArgs{}
	flag.StringVar(&cmd.DangerousRegen, "dangerous_regen", "false",
		"Regenerate everything. If set to true, init server will be regenerated and  all previous changes will be lost")
	flag.StringVar(&cmd.FetchOnly, "fetch-only", "false",
		"generate only the typescript file that holds fetch function")
	flag.StringVar(&cmd.ConfigPath, "config", marvin.ConfigName,
		"Specify path to the gomarvin config file")
	flag.Parse()

	marvin.Run(cmd)
}
