package marvin

import (
	"fmt"
)

func generateInitDirsAndFiles(conf Config, cmd Flags) {
	projectName := conf.ProjectInfo.Name
	projectPath := fmt.Sprintf("./%s", projectName)

	// if the project dir does not exist or regen is true
	if !PathExists(projectPath) || cmd.DangerousRegen {
		// first, generate init dirs
		for i := 0; i < len(INIT_PROJECT_DIRS); i++ {
			CreateDir(fmt.Sprintf("%s/%s", projectName, INIT_PROJECT_DIRS[i]))
		}
		// then generate init files
		generateTemplates(conf, INIT_PROJECT_TEMPLATES)
		renameToDotFiles(conf)
	}

	// if the api/vx/server/ dir for the specified version does not exist or regen is true,
	// create it in the projec_name/internal/api/ dir
	apiVersion := ApiVersion(conf.ProjectInfo.APIPrefix)
	apiVersionDir := fmt.Sprintf("%s/internal/api/%s/server/", projectPath, apiVersion)
	if !PathExists(apiVersionDir) || cmd.DangerousRegen {
		CreateDir(apiVersionDir)
		generateTemplates(conf, initModuleTemplates(apiVersion))
	}
}

// Don't know why but using dotfiles in Templates struct causes errror when executing
// the template. So just rename files to dotfiles once the thing is generated
func renameToDotFiles(conf Config) {
	projectName := conf.ProjectInfo.Name

	RenameFileIfExists(
		fmt.Sprintf("./%s/conf/env.dev", projectName),
		fmt.Sprintf("./%s/conf/.env.dev", projectName),
	)

	RenameFileIfExists(
		fmt.Sprintf("./%s/gitignore", projectName),
		fmt.Sprintf("./%s/.gitignore", projectName),
	)
}
