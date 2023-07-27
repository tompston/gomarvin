package marvin

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func Run(flags *Flags, args []string) {

	// if only the program is executed, with no args, print info
	if len(args) == 0 {
		fmt.Println(gomarvin_info)
	}

	// if "generate" arg is provided in the cli args, generate the project
	if stringExistsInSlice("generate", args) {
		// if gomarvin.json or a custom path to an existing config file is provided
		if PathExists(flags.ConfigPath) {

			conf := ReadConfig(flags.ConfigPath) // read json config file

			// If the api_prefix value in the json file is incorrect,
			// throw an error and exit
			api_version := ApiVersionInterger(conf.ProjectInfo.APIPrefix)
			_ = api_version

			GenerateInit(conf, *flags)     // generate init dirs and files if project dir does not exist or dangerous-regen="true"
			GenerateModules(conf, *flags)  // geenerate module dirs and controller files if exist
			GenerateOptional(conf, *flags) // generate things that are optional
			FormatAfterGen()               // run gofmt to format the project in the dir

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

Version: 0.10.0

Online Editor:
  https://gomarvin.pages.dev

Flags:
  -config		
	Specify path to the gomarvin config file 
	(default "gomarvin.json")
  -dangerous-regen	
  	Regenerate everything. If set to true, all previous changes will be lost 
	(default "false")
  -gut
	generate additional file in the modules dir which concats all of the 
	functions that convert possible response structs to typescript interfaces
	(default "false")`

func GetValueAfterLastSlash(str string) string {
	split := strings.Split(str, "/")
	return split[len(split)-1]
}

// ApiVersion returns the last element of the api_version string
//
// Example:
//
//	ApiVersion("/api/v1") -> "v1"
//	ApiVersion("/api/v2") -> "v2"
func ApiVersion(api_version string) string {
	return GetValueAfterLastSlash(api_version)
}

func ApiVersionInterger(api_version string) int {
	api_v := ApiVersion(api_version)
	v := api_v[1:]
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("the value after the last / should have the form of [v + int] in api_prefix value!: %v", err)
	}

	return i
}
