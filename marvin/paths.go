package marvin

import "fmt"

var init_project_dirs = [...]string{
	// -- root dirs
	"/",
	"/client/",
	"/cmd/api/",
	"/conf/",
	// -- pkg dirs
	"/pkg/response/",
	"/pkg/validate/",
	"/pkg/convert/",
	"/pkg/settings/",
	"/pkg/settings/database/",
	"/pkg/settings/cli/",
	"/pkg/app/",
	// -- internal dirs
	"/internal/api/response/",
}

const typescript_fetch_template = "templates/optional/client/gomarvin.ts.tmpl"
const python_fetch_template = "templates/optional/client/gomarvin.py.tmpl"

var init_project_templates = []Template{
	{template_path: "templates/init/main.go.tmpl", output_dir: "/cmd/api/"}, // main.go
	{template_path: "templates/init/args.go.tmpl", output_dir: "/pkg/settings/cli/"},
	{template_path: "templates/init/README.md.tmpl", output_dir: "/"},
	{template_path: "templates/init/env.dev.tmpl", output_dir: "/conf/"}, // a dot at the start breaks this
	{template_path: "templates/init/gitignore.tmpl", output_dir: "/"},
	{template_path: "templates/init/go.mod.tmpl", output_dir: "/"},
	{template_path: "templates/init/settings.go.tmpl", output_dir: "/pkg/settings/"},
	{template_path: "templates/init/env.go.tmpl", output_dir: "/pkg/settings/"},
	{template_path: "templates/init/database.go.tmpl", output_dir: "/pkg/settings/database/"},
	{template_path: "templates/init/validate.go.tmpl", output_dir: "/pkg/validate/"},
	{template_path: "templates/init/convert.go.tmpl", output_dir: "/pkg/convert/"},
	{template_path: "templates/init/app.go.tmpl", output_dir: "/pkg/app/"},
}

func initModuleTemplates(api_prefix string) []Template {
	return []Template{
		{template_path: "templates/init/router.go.tmpl", output_dir: fmt.Sprintf("/internal/api/%v/server/", api_prefix)},
		{template_path: "templates/init/server.go.tmpl", output_dir: fmt.Sprintf("/internal/api/%v/server/", api_prefix)},
		{template_path: "templates/init/response.go.tmpl", output_dir: "/internal/api/response/"},
	}
}

var optional_sql_templates = []Template{
	{template_path: "templates/optional/sql/functions.sql.tmpl", output_dir: "/db/sql/"},
	{template_path: "templates/optional/sql/sqlc.yaml.tmpl", output_dir: "/db/"},
}
