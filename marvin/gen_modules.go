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
			if !PathExists(module_dir) || cmd.DangerousRegen {
				CreateDir(module_dir) // create the dir
			}

			endpoints := data.Modules.Endpoints

			// if there are endpoints in the module, create files that hold them
			if len(endpoints) != 0 {
				// Create the file which holds controllers
				CreateModuleFile("templates/module/controllers.gen.go.tmpl", module_dir, data)
				// Create the file which holds functions that convert golang structs to ts interfaces

				if cmd.Gut {
					CreateModuleFile("templates/module/ts_interfaces.gen.go.tmpl", module_dir, data)
				}

				// if there is at least a single endpoint which has a body that does not have
				// a lenght of 0, generate the `body.gen.go` file with all of the body structs
				for q := 0; q < len(endpoints); q++ {
					endpoint_body := endpoints[q].Body
					if len(endpoint_body) != 0 {
						CreateModuleFile("templates/module/body.gen.go.tmpl", module_dir, data)
						break
					}
					// break
				}
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

/*


go run main.go -gut=true  \
    -config="./examples/v0.7.0/gomarvin-fiber_with_modules.json" generate




*/
