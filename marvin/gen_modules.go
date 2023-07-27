package marvin

import (
	"fmt"
	"strings"
)

func generateModules(conf Config, flag Flags) {
	projectName := conf.ProjectInfo.Name
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

			apiPrefix := ApiVersion(conf.ProjectInfo.APIPrefix)
			moduleDir := ModuleDir(projectName, module.Name, apiPrefix)

			// if the dir does not exist or regen is true
			if !pathExists(moduleDir) || flag.DangerousRegen {
				createDir(moduleDir) // create the dir
			}

			endpoints := data.Modules.Endpoints

			// if there are endpoints in the module, create files that hold them
			if len(endpoints) != 0 {
				// Create the file which holds controllers
				createModuleFile("templates/module/controllers.gen.go.tmpl", moduleDir, data)

				// Create the file which holds functions that convert golang structs to ts interfaces
				// if the -gut flag was set as true
				if flag.Gut {
					createModuleFile("templates/module/typescript.gen.go.tmpl", moduleDir, data)
				}

				// if there is at least a single endpoint which has a body that does not have
				// a lenght of 0, generate the `body.gen.go` file with all of the body structs
				for q := 0; q < len(endpoints); q++ {
					endpoint_body := endpoints[q].Body
					if len(endpoint_body) != 0 {
						createModuleFile("templates/module/body.gen.go.tmpl", moduleDir, data)
						break
					}
				}
			}
		}
	}
}

// Create a file in the module directory
func createModuleFile(templatePath, moduleDir string, data Project) {
	template_name, outputFile := createTemplateAndOutputName(templatePath)
	full_output_path := fmt.Sprintf("./%s%s", moduleDir, outputFile)
	ExecuteTemplate(template_name, templatePath, full_output_path, data)
	fmt.Println(CREATED_MSG, full_output_path)
}

// name of the folder in the project to which the new modules will be added
const ModuleOutputDir = "modules"

// "task" 	-> "task_module"
func ModuleDirName(moduleName string) string {
	return strings.ToLower(fmt.Sprintf("%s%s", moduleName, "_module"))
}

func ProjectModuleDirPath(projectName, apiVersion string) string {
	return fmt.Sprintf("%s/internal/api/%s/%s", projectName, apiVersion, ModuleOutputDir)
}

// ("my_Project", "Task") --> my_project/modules/task_module/
func ModuleDir(projectName, moduleName, apiVersion string) string {
	return fmt.Sprintf("%s/%s/",
		ProjectModuleDirPath(projectName, apiVersion),
		ModuleDirName(moduleName))
}
