package marvin

import (
	"fmt"
	"strings"

	conv "github.com/tompston/gomarvin/marvin/utils/convert"
)

func GenerateSql(conf Config, cmd CmdArgs) {

	project_name := conf.ProjectInfo.Name
	modules := conf.Modules
	project_sql_dir := fmt.Sprintf("./%s/db/sql/", project_name)

	if len(modules) != 0 { // if there are modules in the config

		// if path to /db/sql/ dir does not exist or  dangerous_regen is true, create it and create init sql files
		if !PathExists(project_sql_dir) || cmd.DangerousRegen == "true" {
			CreateDir(project_sql_dir)
			GenerateTemplates(conf, optional_sql_templates) // generate sqlc.yml and functions.sql
		}

		// if the /db/sql/ dir exists, loop over modules and pass data to the templates
		for i := 0; i < len(modules); i++ { // for each  module in modules

			// project_info := conf.ProjectInfo
			module := modules[i]
			data := Project{ProjectInfo: conf.ProjectInfo, Modules: module}
			lowercase_module_name := conv.ConvertToLowercase(module.Name)

			// regenerate the files that hold the sql tables and queries on each run
			template_path := "templates/optional/sql/__module__.sql.gen.txt.tmpl"
			template_name, output_file := GenerateTemplateAndOutputName(template_path)
			full_output_path := fmt.Sprintf("%s%s", project_sql_dir, output_file)
			full_output_path = strings.Replace(full_output_path, "__module__", lowercase_module_name, -1)
			ExecuteTemplate(template_name, template_path, full_output_path, data)

			fmt.Println(CREATED_MSG, full_output_path)
		}
	}
}
