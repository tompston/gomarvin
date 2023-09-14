package marvin

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"path/filepath"
)

const (
	// message printed when a file is created
	CREATED_MSG = "\033[32m * CREATED\033[0m"

	// predefined name of config file that will be looked up if no other config path is specified
	ConfigName = "gomarvin.json"
)

// https://golangdocs.com/command-line-arguments-in-golang
type Flags struct {
	DangerousRegen bool   // ( default = false )
	Gut            bool   // ( default = false )
	ConfigPath     string // if no params are  specified, get the config from current dir ( default = "./gomarvin.json" )
}

// https://paulgorman.org/technical/blog/20171113164018.html
func readConfig(path string) Config {
	c := flag.String("c", path, "Specify the configuration file.")
	flag.Parse()
	f, err := os.Open(*c)
	if err != nil {
		log.Fatal("can't open config file: ", err)
	}
	defer f.Close()

	conf := Config{}
	if err := json.NewDecoder(f).Decode(&conf); err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}

	return conf
}

// if you pass a nested path of folders that do not exist, this function
// will also create those folders
func createDir(path string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		return nil, err
	}
	return os.Create(path)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return !os.IsNotExist(err)
	}
	return true
}

func renameFileIfExists(originalPath, newPath string) {
	if pathExists(originalPath) {
		if err := os.Rename(originalPath, newPath); err != nil {
			log.Fatal(err)
		}
	}
}

// output text if no args are provided
const GOMARVIN_INFO = `Usage:
	gomarvin generate

Version: 0.10.1

Online Editor:
	https://gomarvin.pages.dev

Flags:
  -config		
	Specify path to the gomarvin config file  (default "gomarvin.json")

  -dangerous-regen
	Regenerate everything. If set to true, all previous changes will 
	be lost  (default "false")

  -gut
	generate functions which convert possible api response structs to 
	typescript interfaces (default "false")`
