package marvin

import (
	"fmt"
	"os/exec"
)

func Run(flags *Flags, args []string) {

	// if only the program is executed, with no args, print info
	if len(args) == 0 {
		fmt.Println(gomarvin_info)
	}

	// if "generate" arg is provided in the cli args, generate the project
	if stringExistsInSlice("generate", args) {

		// fmt.Println("flags: ", flags.ConfigPath)

		// if gomarvin.json or a custom path to an existing config file is provided
		if PathExists(flags.ConfigPath) {

			conf := ReadConfig(flags.ConfigPath) // read json config file

			// if fetch_only is set to default value ("false"), generate the whole project
			if !flags.FetchOnly {
				GenerateInit(conf, *flags)     // generate init dirs and files if project dir does not exist or dangerous-regen="true"
				GenerateModules(conf, *flags)  // geenerate module dirs and controller files if exist
				GenerateOptional(conf, *flags) // generate things that are optional
				FormatAfterGen()               // run gofmt to format the project in the dir
				// GenerateSeeder(conf, *flags)

			} else if flags.FetchOnly {
				GenerateOnlyFetchFunctions(conf, *flags)
			}

		} else {
			fmt.Println("* ERROR :: Could not find the config file!")
		}
	}
}

// run gofmt after codegen to format the generated code correctly / remove whitespace stuff
func FormatAfterGen() {
	flags := exec.Command("gofmt", "-s", "-w", ".")
	flags.Run()
}

func stringExistsInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// const CREATED_MSG = "* CREATED"
const CREATED_MSG = "\033[32m * CREATED\033[0m"

var gomarvin_info = `Usage:
  gomarvin generate

Version: 0.6.0

Online Editor:
  https://gomarvin.pages.dev

Flags:
  -config		
	Specify path to the gomarvin config file 
	(default "gomarvin.json")
  -dangerous-regen	
  	Regenerate everything. If set to true, all previous changes will be lost 
	(default "false")
  -fetch-only		
  	generate only the typescript file that holds fetch function 
	(default "false")`
