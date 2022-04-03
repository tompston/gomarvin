package marvin

import (
	"fmt"
	"strings"
)

func GenerateModules(conf Config, cmd CmdArgs) {

	project_name := conf.ProjectInfo.Name
	modules := conf.Modules

	if len(modules) != 0 { // if there are modules in the config file
		for i := 0; i < len(modules); i++ { // for each  module in modules

			module := modules[i]

			// pass down structs that hold global project data + single module
			data := Project{ProjectInfo: conf.ProjectInfo, Modules: module}

			module_dir := ModuleDir(project_name, module.Name)

			// if the dir does not exist or regen is true
			if !PathExists(module_dir) || cmd.DangerousRegen == "true" {
				CreateDir(module_dir) // create the dir
			}

			// if there are endpoints in the module, generate them
			if len(data.Modules.Endpoints) != 0 {
				template_path := "templates/module/controllers.gen.go.tmpl"
				template_name, output_file := GenerateTemplateAndOutputName(template_path)
				full_output_path := fmt.Sprintf("./%s%s", module_dir, output_file)
				full_output_path = strings.Replace(full_output_path, "__module__", "controllers.gen", -1)
				ExecuteTemplate(template_name, template_path, full_output_path, data)
				fmt.Println(CREATED_MSG, full_output_path)
			}
		}
	}
}
