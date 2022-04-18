package marvin

import "fmt"

func GenerateOptional(conf Config, cmd CmdArgs) {

	if conf.ProjectInfo.IncludeSQL { // if include_sql = true
		GenerateSql(conf, cmd)
	}

	if conf.ProjectInfo.IncludeTsFetch {
		GenerateTypescriptFetchFuncs(conf, cmd)
	}
}

// if there are modules in the config, generate /main.gen.ts in the root of the project
func GenerateTypescriptFetchFuncs(conf Config, cmd CmdArgs) {

	public_dir := fmt.Sprintf("./%s/public/", conf.ProjectInfo.Name)

	// if the public dir does not exist, create it
	if PathExists(public_dir) {
		CreateDir(public_dir)
	}

	// if there are module in the config, create Typescript fetch functions
	if len(conf.Modules) != 0 {
		GenerateSingleTemplate(conf, "templates/optional/ts/main.gen.ts.tmpl", "/public/")
	}
}
