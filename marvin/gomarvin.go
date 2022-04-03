package marvin

import "fmt"

func Run(cmd *CmdArgs) {

	// fmt.Println(*&cmd.Config)

	if PathExists(cmd.ConfigPath) { //
		// fmt.Println("Cmd arg Test value ::", cmd.ConfigPath) // inspec args
		// conf := ReadConfig(*&cmd.ConfigPath)                 // read json config file
		conf := ReadConfig(cmd.ConfigPath) // read json config file
		// fmt.Println(conf)

		GenerateInit(conf, *cmd)     // generate init dirs and files if project dir does not exist or dangerous_regen="true"
		GenerateModules(conf, *cmd)  // geenerate module dirs and controller files if exist
		GenerateOptional(conf, *cmd) // generate things that are optional

	} else {
		fmt.Println("ERROR :: Could not find the config file!")
	}

	// go run main.go -config="./previous/gomarvin.json"
}

var CREATED_COL = "\033[32m"

// var CREATED_MSG = fmt.Sprintf(string(CREATED_COL), "CREATED")
var CREATED_MSG = "* CREATED"
