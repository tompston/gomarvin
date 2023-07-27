package marvin

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	// message printed when a file is created
	CREATED_MSG = "\033[32m * CREATED\033[0m"

	// predefined name of config file that will be looked up if no other config path is specified
	ConfigName = "gomarvin.json"

	// output text if no args are provided
	GOMARVIN_INFO = `Usage:
	  gomarvin generate
	
	Version: 0.10.0
	
	Online Editor:
	  https://gomarvin.pages.dev
	
	Flags:
	  -config		
		Specify path to the gomarvin config file  (default "gomarvin.json")
	  -dangerous-regen
		  Regenerate everything. If set to true, all previous changes will be lost  (default "false")
	  -gut
		generate functions which convert possible api response structs to typescript interfaces (default "false")`
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

var INIT_PROJECT_DIRS = []string{
	"/",
	"/cmd/api/",
	"/conf/",
	"/pkg/validate/",
	"/pkg/convert/",
	"/pkg/settings/",
	"/pkg/settings/database/",
	"/pkg/settings/cli/",
	"/pkg/app/",
	"/internal/api/response/",
}

const (
	TYPESCRIPT_FETCH_TEMPLATE = "templates/optional/client/gomarvin.ts.tmpl"
	PYTHON_FETCH_TEMPLATE     = "templates/optional/client/gomarvin.py.tmpl"
)

var INIT_PROJECT_TEMPLATES = []Template{
	{templatePath: "templates/init/main.go.tmpl", outputDir: "/cmd/api/"}, // main.go
	{templatePath: "templates/init/args.go.tmpl", outputDir: "/pkg/settings/cli/"},
	{templatePath: "templates/init/README.md.tmpl", outputDir: "/"},
	{templatePath: "templates/init/env.dev.tmpl", outputDir: "/conf/"}, // a dot at the start breaks this
	{templatePath: "templates/init/gitignore.tmpl", outputDir: "/"},
	{templatePath: "templates/init/go.mod.tmpl", outputDir: "/"},
	{templatePath: "templates/init/env.go.tmpl", outputDir: "/pkg/settings/"},
	{templatePath: "templates/init/database.go.tmpl", outputDir: "/pkg/settings/database/"},
	{templatePath: "templates/init/validate.go.tmpl", outputDir: "/pkg/validate/"},
	{templatePath: "templates/init/convert.go.tmpl", outputDir: "/pkg/convert/"},
	{templatePath: "templates/init/app.go.tmpl", outputDir: "/pkg/app/"},
}

var OPTIONAL_SQL_TEMPLATES = []Template{
	{templatePath: "templates/optional/sql/functions.sql.tmpl", outputDir: "/db/sql/"},
	{templatePath: "templates/optional/sql/sqlc.yaml.tmpl", outputDir: "/db/"},
}

func initModuleTemplates(apiPrefix string) []Template {
	return []Template{
		{templatePath: "templates/init/router.go.tmpl", outputDir: fmt.Sprintf("/internal/api/%v/server/", apiPrefix)},
		{templatePath: "templates/init/server.go.tmpl", outputDir: fmt.Sprintf("/internal/api/%v/server/", apiPrefix)},
		{templatePath: "templates/init/response.go.tmpl", outputDir: "/internal/api/response/"},
	}
}
