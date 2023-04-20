package marvin

import (
	"fmt"
)

func GenerateInit(conf Config, cmd Flags) {
	project_name := conf.ProjectInfo.Name
	project_path := fmt.Sprintf("./%s", project_name)

	// if the project dir does not exist or regen is true
	if !PathExists(project_path) || cmd.DangerousRegen {
		GenerateInitProjectDirs(project_name)           // first, generate init dirs
		GenerateTemplates(conf, init_project_templates) // then generate init files

		// go run main.go -dangerous-regen=true generate
		// if the api/vx/server/ dir for the specified version does not exist or regen is true,
		// create it in the projec_name/internal/api/ dir
		api_v := ApiVersion(conf.ProjectInfo.APIPrefix)
		api_version_dir := fmt.Sprintf("%s/internal/api/%s/server/", project_path, api_v)
		// fmt.Printf("api_version_dir: %v\n", api_version_dir)
		if !PathExists(api_version_dir) || cmd.DangerousRegen {
			CreateDir(api_version_dir)
			GenerateTemplates(conf, initModuleTemplates(api_v))
		}

		RenameToDotFiles(conf)
	}
}

// loop over dirs specified in the  "init_project_dirs" to create them
func GenerateInitProjectDirs(project_name string) {
	for i := 0; i < len(init_project_dirs); i++ {
		CreateDir(fmt.Sprintf("%s/%s", project_name, init_project_dirs[i]))
	}
}

// Don't know why but using dotfiles in Templates struct causes errror when executing
// the template. So just rename files to dotfiles once the thing is generated
func RenameToDotFiles(conf Config) {

	project := conf.ProjectInfo.Name

	env_original_path := fmt.Sprintf("./%s/env", project)
	env_new_path := fmt.Sprintf("./%s/.env", project)
	RenameFileIfExists(env_original_path, env_new_path)

	gitignore_original_path := fmt.Sprintf("./%s/gitignore", project)
	gitignore_new_path := fmt.Sprintf("./%s/.gitignore", project)
	RenameFileIfExists(gitignore_original_path, gitignore_new_path)
}
