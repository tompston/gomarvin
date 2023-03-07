package main

import (
	"flag"
	"os"

	marvin "github.com/tompston/gomarvin/marvin"
)

func main() {
	// go run main.go -h

	args := os.Args[1:]
	flags := &marvin.Flags{}
	flag.BoolVar(&flags.DangerousRegen, "dangerous-regen", false, "Regenerate everything. If set to true, init server will be regenerated and  all previous changes will be lost")
	flag.BoolVar(&flags.FetchOnly, "fetch-only", false, "generate only the typescript file that holds fetch function")
	flag.StringVar(&flags.ConfigPath, "config", marvin.ConfigName, "Specify path to the gomarvin config file")
	flag.BoolVar(&flags.Gut, "gut", false, "generate additional file in the modules dir which concats all of the functions that convert possible response structs to typescript interfaces")
	flag.Parse()

	marvin.Run(flags, args)
}
