package marvin

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
	if len(conf.Modules) != 0 {
		GenerateSingleTemplate(conf, "templates/optional/ts/main.gen.ts.tmpl", "/")
	}
}
