package marvin

import "fmt"

func GenerateOptional(conf Config, cmd Flags) {

	if conf.ProjectInfo.IncludeSQL { // if include_sql = true
		GenerateSql(conf, cmd)
	}

	if conf.ProjectInfo.IncludeFetch {
		GenerateFetchFunctions(conf, cmd)
	}
}

// if there are modules in the config, generate the python and typescript fetch clients
// in the client dir
func GenerateFetchFunctions(conf Config, cmd Flags) {
	client_dir := fmt.Sprintf("./%s/client/", conf.ProjectInfo.Name)

	// if the public dir does not exist, create it
	if !PathExists(client_dir) {
		CreateDir(client_dir)
	}
	// if there are modules in the config file, create the typescript file
	if len(conf.Modules) != 0 {
		GenerateSingleTemplate(conf, typescript_fetch_template, "/client/")
		GenerateSingleTemplate(conf, python_fetch_template, "/client/")
	}
}
