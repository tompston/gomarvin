package marvin

import (
	"fmt"
	"strings"

	conv "github.com/tompston/gomarvin/marvin/utils/convert"
)

func generateFetchClients(conf Config, flags Flags) {
	if conf.ProjectInfo.IncludeSQL {
		generateSQL(conf, flags)
	}

	// generate the python and typescript clients if the user wants them
	if conf.ProjectInfo.IncludeFetch {
		pyClientDir := fmt.Sprintf("./%s/client/py/", conf.ProjectInfo.Name)
		tsClientDir := fmt.Sprintf("./%s/client/ts/", conf.ProjectInfo.Name)

		// if there are modules in the config file, create the typescript and python client
		if len(conf.Modules) != 0 {
			if !PathExists(pyClientDir) {
				CreateDir(pyClientDir)
			}
			generateSingleTemplate(conf, PYTHON_FETCH_TEMPLATE, "/client/py/")

			if !PathExists(tsClientDir) {
				CreateDir(tsClientDir)
			}
			generateSingleTemplate(conf, TYPESCRIPT_FETCH_TEMPLATE, "/client/ts/")
		}
	}
}

func generateSQL(conf Config, flags Flags) {

	projectName := conf.ProjectInfo.Name
	modules := conf.Modules
	sqlDir := fmt.Sprintf("./%s/db/sql/", projectName)

	if len(modules) != 0 {
		// if path to /db/sql/ dir does not exist or dangerous-regen is true, create it and create init sql files
		if !PathExists(sqlDir) || flags.DangerousRegen {
			CreateDir(sqlDir)
			generateTemplates(conf, OPTIONAL_SQL_TEMPLATES) // generate sqlc.yml and functions.sql
		}

		// if the /db/sql/ dir exists, loop over modules and pass data to the templates
		for i := 0; i < len(modules); i++ { // for each  module in modules
			module := modules[i]
			data := Project{ProjectInfo: conf.ProjectInfo, Modules: module}

			// regenerate the files that hold the sql tables and queries on each run
			templatePath := "templates/optional/sql/__module__.sql.gen.txt.tmpl"
			template_name, outputFile := createTemplateAndOutputName(templatePath)
			fullOutputPath := fmt.Sprintf("%s%s", sqlDir, outputFile)
			fullOutputPath = strings.Replace(fullOutputPath, "__module__", conv.ConvertToLowercase(module.Name), -1)
			ExecuteTemplate(template_name, templatePath, fullOutputPath, data)

			fmt.Println(CREATED_MSG, fullOutputPath)
		}
	}
}
