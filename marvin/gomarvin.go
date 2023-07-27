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
			api_version := getApiVersionInt(conf.ProjectInfo.APIPrefix)
			_ = api_version

			generateInitDirsAndFiles(conf, *flags) // generate init dirs and files if project dir does not exist or dangerous-regen="true"
			GenerateModules(conf, *flags)          // geenerate module dirs and controller files if exist
			generateFetchClients(conf, *flags)     // generate things that are optional
			FormatAfterGen()                       // run gofmt to format the project in the dir

		} else {
			fmt.Println("ERROR: Could not find the config file!")
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
func ApiVersion(apiVersion string) string {
	return GetValueAfterLastSlash(apiVersion)
}

func getApiVersionInt(apiVersion string) int {
	api_v := ApiVersion(apiVersion)
	v := api_v[1:]
	i, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalf("the value after the last / should have the form of [v + int] in api_prefix value!: %v", err)
	}

	return i
}
