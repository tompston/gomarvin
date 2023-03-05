package marvin

var init_project_dirs = [...]string{
	"/",
	"/settings/database/",
	"/lib/response/",
	"/lib/validate/",
	"/lib/convert/",
	"/lib/utils/",
	"/app/",
	"/public/",
}

const typescript_fetch_template = "templates/optional/ts/gomarvin.gen.ts.tmpl"

// const typescript_seed_template = "templates/optional/ts/seeder.gen.ts.tmpl"

var init_project_templates = []Template{
	{template_path: "templates/init/main.go.tmpl", output_dir: "/"},
	{template_path: "templates/init/README.md.tmpl", output_dir: "/"},
	{template_path: "templates/init/env.tmpl", output_dir: "/"}, // a dot at the start breaks this
	{template_path: "templates/init/gitignore.tmpl", output_dir: "/"},
	{template_path: "templates/init/go.mod.tmpl", output_dir: "/"},
	{template_path: "templates/init/Dockerfile.tmpl", output_dir: "/"},
	{template_path: "templates/init/settings.go.tmpl", output_dir: "/settings/"},
	{template_path: "templates/init/env.go.tmpl", output_dir: "/settings/"},
	{template_path: "templates/init/database.go.tmpl", output_dir: "/settings/database/"},
	{template_path: "templates/init/response.go.tmpl", output_dir: "/lib/response/"},
	{template_path: "templates/init/validate.go.tmpl", output_dir: "/lib/validate/"},
	{template_path: "templates/init/utils.go.tmpl", output_dir: "/lib/utils/"},
	{template_path: "templates/init/router.go.tmpl", output_dir: "/app/"},
	{template_path: "templates/init/server.go.tmpl", output_dir: "/app/"},
	// convert
	{template_path: "templates/init/convert.go.tmpl", output_dir: "/lib/convert/"},
}

var optional_sql_templates = []Template{
	{template_path: "templates/optional/sql/functions.sql.tmpl", output_dir: "/db/sql/"},
	{template_path: "templates/optional/sql/sqlc.yaml.tmpl", output_dir: "/db/"},
}
