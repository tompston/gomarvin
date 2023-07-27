package marvin

import "fmt"

var INIT_PROJECT_DIRS = [...]string{
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

func initModuleTemplates(apiPrefix string) []Template {
	return []Template{
		{templatePath: "templates/init/router.go.tmpl", outputDir: fmt.Sprintf("/internal/api/%v/server/", apiPrefix)},
		{templatePath: "templates/init/server.go.tmpl", outputDir: fmt.Sprintf("/internal/api/%v/server/", apiPrefix)},
		{templatePath: "templates/init/response.go.tmpl", outputDir: "/internal/api/response/"},
	}
}

var OPTIONAL_SQL_TEMPLATES = []Template{
	{templatePath: "templates/optional/sql/functions.sql.tmpl", outputDir: "/db/sql/"},
	{templatePath: "templates/optional/sql/sqlc.yaml.tmpl", outputDir: "/db/"},
}
