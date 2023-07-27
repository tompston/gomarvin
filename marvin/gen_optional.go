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
	pyClientDir := fmt.Sprintf("./%s/client/py/", conf.ProjectInfo.Name)
	tsClientDir := fmt.Sprintf("./%s/client/ts/", conf.ProjectInfo.Name)

	// if there are modules in the config file, create the typescript and python client
	if len(conf.Modules) != 0 {
		if !PathExists(pyClientDir) {
			CreateDir(pyClientDir)
		}
		GenerateSingleTemplate(conf, python_fetch_template, "/client/py/")

		if !PathExists(tsClientDir) {
			CreateDir(tsClientDir)
		}
		GenerateSingleTemplate(conf, typescript_fetch_template, "/client/ts/")
	}
}
