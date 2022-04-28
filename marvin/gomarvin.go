package marvin

import (
	"fmt"
	"os/exec"
)

func Run(cmd *CmdArgs) {

	if PathExists(cmd.ConfigPath) { // if gomarvin.json or provided config file exists

		conf := ReadConfig(cmd.ConfigPath) // read json config file

		fmt.Println(cmd.FetchOnly)

		// if fetch_only is set to default value ("false"), generate the whole project
		if cmd.FetchOnly == "false" {
			GenerateInit(conf, *cmd)     // generate init dirs and files if project dir does not exist or dangerous_regen="true"
			GenerateModules(conf, *cmd)  // geenerate module dirs and controller files if exist
			GenerateOptional(conf, *cmd) // generate things that are optional
			FormatAfterGen()             // run gofmt to format the project in the dir
		} else if cmd.FetchOnly == "true" {
			// go run main.go -config="examples/v0.3.0/gomarvin-fiber_with_modules.json" -fetch_only="true"
			GenerateOnlyFetchFunctions(conf, *cmd)
		}

	} else {
		fmt.Println("* ERROR :: Could not find the config file!")
	}
}

// run gofmt after codegen to format the generated code correctly / remove whitespace stuff
func FormatAfterGen() {
	cmd := exec.Command("gofmt", "-s", "-w", ".")
	cmd.Run()
}

var CREATED_MSG = "* CREATED"
