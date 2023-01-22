package marvin

import (
	"fmt"
)

func GenerateModules(conf Config, cmd Flags) {

	project_name := conf.ProjectInfo.Name
	modules := conf.Modules

	// if there are modules in the config file
	if len(modules) != 0 {
		// for each  module in modules
		for i := 0; i < len(modules); i++ {

			module := modules[i]

			// pass down structs that hold global project data + single module
			data := Project{
				ProjectInfo: conf.ProjectInfo,
				Modules:     module,
			}

			module_dir := ModuleDir(project_name, module.Name)

			// if the dir does not exist or regen is true
			if !PathExists(module_dir) || cmd.DangerousRegen == "true" {
				CreateDir(module_dir) // create the dir
			}

			// if there are endpoints in the module, create files that hold them
			if len(data.Modules.Endpoints) != 0 {
				CreateModuleFile("templates/module/controllers.gen.go.tmpl", module_dir, data)
				CreateModuleFile("templates/module/body.gen.go.tmpl", module_dir, data)
			}
		}
	}
}

func CreateModuleFile(template_path string, module_dir string, data Project) {
	template_name, output_file := GenerateTemplateAndOutputName(template_path)
	full_output_path := fmt.Sprintf("./%s%s", module_dir, output_file)
	ExecuteTemplate(template_name, template_path, full_output_path, data)
	fmt.Println(CREATED_MSG, full_output_path)
}
