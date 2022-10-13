package marvin

import (
	"fmt"
	"os/exec"
)

func Run(cmd *CmdArgs, args []string) {

	fmt.Println(args)

	if len(args) == 0 {
		fmt.Println(gomarvin_info)
	}

	// if "generate" arg is provided in the cli args, generate the project
	if stringExistsInSlice("generate", args) {

		// if gomarvin.json or provided config file path exists
		if PathExists(cmd.ConfigPath) {
			conf := ReadConfig(cmd.ConfigPath) // read json config file

			// if fetch_only is set to default value ("false"), generate the whole project
			if cmd.FetchOnly == "false" {
				GenerateInit(conf, *cmd)     // generate init dirs and files if project dir does not exist or dangerous-regen="true"
				GenerateModules(conf, *cmd)  // geenerate module dirs and controller files if exist
				GenerateOptional(conf, *cmd) // generate things that are optional
				FormatAfterGen()             // run gofmt to format the project in the dir
			} else if cmd.FetchOnly == "true" {
				// go run main.go -config="examples/v0.3.0/gomarvin-fiber_with_modules.json" -fetch-only="true"
				GenerateOnlyFetchFunctions(conf, *cmd)
			}

		} else {
			fmt.Println("* ERROR :: Could not find the config file!")
		}
	}

}

// run gofmt after codegen to format the generated code correctly / remove whitespace stuff
func FormatAfterGen() {
	cmd := exec.Command("gofmt", "-s", "-w", ".")
	cmd.Run()
}

func stringExistsInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

var CREATED_MSG = "* CREATED"

var gomarvin_info = `
Usage:
  gomarvin generate

Flags:
  -x, --experimental   enable experimental features (default: false)
  -f, --file string    specify an alternate config file (default: sqlc.yaml)
  -h, --help           help for sqlc`
