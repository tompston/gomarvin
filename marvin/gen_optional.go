package marvin

import (
	"fmt"
	"strings"

	conv "github.com/tompston/gomarvin/marvin/utils/convert"
)

func generateFetchClients(conf Config, cmd Flags) {
	if conf.ProjectInfo.IncludeSQL {
		generateSQL(conf, cmd)
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
			GenerateSingleTemplate(conf, python_fetch_template, "/client/py/")

			if !PathExists(tsClientDir) {
				CreateDir(tsClientDir)
			}
			GenerateSingleTemplate(conf, typescript_fetch_template, "/client/ts/")
		}
	}
}

func generateSQL(conf Config, cmd Flags) {

	projectName := conf.ProjectInfo.Name
	modules := conf.Modules
	project_sql_dir := fmt.Sprintf("./%s/db/sql/", projectName)

	if len(modules) != 0 {
		// if path to /db/sql/ dir does not exist or  dangerous-regen is true, create it and create init sql files
		if !PathExists(project_sql_dir) || cmd.DangerousRegen {
			CreateDir(project_sql_dir)
			generateTemplates(conf, optional_sql_templates) // generate sqlc.yml and functions.sql
		}

		// if the /db/sql/ dir exists, loop over modules and pass data to the templates
		for i := 0; i < len(modules); i++ { // for each  module in modules
			module := modules[i]
			data := Project{ProjectInfo: conf.ProjectInfo, Modules: module}

			// regenerate the files that hold the sql tables and queries on each run
			template_path := "templates/optional/sql/__module__.sql.gen.txt.tmpl"
			template_name, output_file := GenerateTemplateAndOutputName(template_path)
			full_output_path := fmt.Sprintf("%s%s", project_sql_dir, output_file)
			full_output_path = strings.Replace(full_output_path, "__module__", conv.ConvertToLowercase(module.Name), -1)
			ExecuteTemplate(template_name, template_path, full_output_path, data)

			fmt.Println(CREATED_MSG, full_output_path)
		}
	}
}
