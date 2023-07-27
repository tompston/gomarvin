package marvin

import (
	"fmt"
	"strings"

	"github.com/tompston/gomarvin/marvin/utils/convert"
)

func generateInitDirsAndFiles(conf Config, cmd Flags) {
	projectName := conf.ProjectInfo.Name
	projectPath := fmt.Sprintf("./%s", projectName)

	// if the project dir does not exist or regen is true
	if !pathExists(projectPath) || cmd.DangerousRegen {
		// first, generate init dirs
		for i := 0; i < len(INIT_PROJECT_DIRS); i++ {
			createDir(fmt.Sprintf("%s/%s", projectName, INIT_PROJECT_DIRS[i]))
		}
		// then generate init files
		generateTemplates(conf, INIT_PROJECT_TEMPLATES)
		renameToDotFiles(conf)
	}

	// if the api/vx/server/ dir for the specified version does not exist or regen is true,
	// create it in the projec_name/internal/api/ dir
	apiVersion := ApiVersion(conf.ProjectInfo.APIPrefix)
	apiVersionDir := fmt.Sprintf("%s/internal/api/%s/server/", projectPath, apiVersion)
	if !pathExists(apiVersionDir) || cmd.DangerousRegen {
		createDir(apiVersionDir)
		generateTemplates(conf, initModuleTemplates(apiVersion))
	}
}

// Don't know why but using dotfiles in Templates struct causes errror when executing
// the template. So just rename files to dotfiles once the thing is generated
func renameToDotFiles(conf Config) {
	projectName := conf.ProjectInfo.Name

	renameFileIfExists(
		fmt.Sprintf("./%s/conf/env.dev", projectName),
		fmt.Sprintf("./%s/conf/.env.dev", projectName),
	)

	renameFileIfExists(
		fmt.Sprintf("./%s/gitignore", projectName),
		fmt.Sprintf("./%s/.gitignore", projectName),
	)
}

func generateOptionalFiles(conf Config, flags Flags) {
	if conf.ProjectInfo.IncludeSQL {
		generateSQL(conf, flags)
	}

	// generate the python and typescript clients if the user wants them
	if conf.ProjectInfo.IncludeFetch {
		pyClientDir := fmt.Sprintf("./%s/client/py/", conf.ProjectInfo.Name)
		tsClientDir := fmt.Sprintf("./%s/client/ts/", conf.ProjectInfo.Name)

		// if there are modules in the config file, create the typescript and python client
		if len(conf.Modules) != 0 {
			if !pathExists(pyClientDir) {
				createDir(pyClientDir)
			}
			generateSingleTemplate(conf, PYTHON_FETCH_TEMPLATE, "/client/py/")

			if !pathExists(tsClientDir) {
				createDir(tsClientDir)
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
		if !pathExists(sqlDir) || flags.DangerousRegen {
			createDir(sqlDir)
			generateTemplates(conf, OPTIONAL_SQL_TEMPLATES) // generate sqlc.yml and functions.sql
		}

		// if the /db/sql/ dir exists, loop over modules and pass data to the templates
		for i := 0; i < len(modules); i++ {
			module := modules[i]
			data := Project{ProjectInfo: conf.ProjectInfo, Modules: module}

			// regenerate the files that hold the sql tables and queries on each run
			templatePath := "templates/optional/sql/__module__.sql.gen.txt.tmpl"
			template_name, outputFile := createTemplateAndOutputName(templatePath)
			fullOutputPath := fmt.Sprintf("%s%s", sqlDir, outputFile)
			fullOutputPath = strings.Replace(fullOutputPath, "__module__", convert.ToLowercase(module.Name), -1)
			executeTemplate(template_name, templatePath, fullOutputPath, data)

			fmt.Println(CREATED_MSG, fullOutputPath)
		}
	}
}
