package marvin

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// name of the folder in the project to which the new modules will be added
var ModuleOutputDir = "modules"

// if you pass a nested path of folders that do not exist, this function
// will also create those folders
func CreateDir(path string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		return nil, err
	}
	return os.Create(path)
}

// "task" 	-> "task_module"
func ModuleDirName(module_name string) string {
	return strings.ToLower(fmt.Sprintf("%s%s", module_name, "_module"))
}

// "my_project" 	-> "my_project/modules"
func ProjectModuleDirPath(project_name string) string {
	return fmt.Sprintf("%s/%s", project_name, ModuleOutputDir)
}

// ("my_Project", "Task") --> my_project/modules/task_module/
func ModuleDir(project_name string, module_name string) string {
	return fmt.Sprintf("%s/%s/", ProjectModuleDirPath(project_name), ModuleDirName(module_name))
}

// pass in the relative path to the dir you want to check. If the dir exists, return true
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return !os.IsNotExist(err)
	}
	return true
}

func RenameFileIfExists(original_path string, new_path string) {
	if PathExists(original_path) {
		e := os.Rename(original_path, new_path)
		if e != nil {
			log.Fatal(e)
		}
	}
}
