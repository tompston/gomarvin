package marvin

import (
	"fmt"
)

func generateInitDirsAndFiles(conf Config, cmd Flags) {
	projectName := conf.ProjectInfo.Name
	projectPath := fmt.Sprintf("./%s", projectName)

	// if the project dir does not exist or regen is true
	if !PathExists(projectPath) || cmd.DangerousRegen {
		generateInitProjectDirs(projectName)            // first, generate init dirs
		generateTemplates(conf, init_project_templates) // then generate init files
		renameToDotFiles(conf)
	}

	// if the api/vx/server/ dir for the specified version does not exist or regen is true,
	// create it in the projec_name/internal/api/ dir
	api_v := ApiVersion(conf.ProjectInfo.APIPrefix)
	api_version_dir := fmt.Sprintf("%s/internal/api/%s/server/", projectPath, api_v)
	// fmt.Printf("api_version_dir: %v\n", api_version_dir)
	if !PathExists(api_version_dir) || cmd.DangerousRegen {
		CreateDir(api_version_dir)
		generateTemplates(conf, initModuleTemplates(api_v))
	}
}

// loop over dirs specified in the  "init_project_dirs" to create them
func generateInitProjectDirs(projectName string) {
	for i := 0; i < len(init_project_dirs); i++ {
		CreateDir(fmt.Sprintf("%s/%s", projectName, init_project_dirs[i]))
	}
}

// Don't know why but using dotfiles in Templates struct causes errror when executing
// the template. So just rename files to dotfiles once the thing is generated
func renameToDotFiles(conf Config) {
	project := conf.ProjectInfo.Name

	envOriginalPath := fmt.Sprintf("./%s/conf/env.dev", project)
	envNewPath := fmt.Sprintf("./%s/conf/.env.dev", project)
	RenameFileIfExists(envOriginalPath, envNewPath)

	gitignoreOriginalPath := fmt.Sprintf("./%s/gitignore", project)
	gitignoreNewPath := fmt.Sprintf("./%s/.gitignore", project)
	RenameFileIfExists(gitignoreOriginalPath, gitignoreNewPath)
}
