package marvin

var init_project_dirs = [...]string{
	"/",
	"/cmd/api/",
	"/pkg/response/",
	"/pkg/validate/",
	"/pkg/convert/",
	"/pkg/settings/",
	"/pkg/settings/database/",
	"/pkg/environment/",
	"/pkg/server/",
	"/public/",
}

const typescript_fetch_template = "templates/optional/ts/gomarvin.gen.ts.tmpl"

// const typescript_seed_template = "templates/optional/ts/seeder.gen.ts.tmpl"

// var init_project_templates = []Template{
// 	{template_path: "templates/init/main.go.tmpl", output_dir: "/"},
// 	{template_path: "templates/init/README.md.tmpl", output_dir: "/"},
// 	{template_path: "templates/init/env.tmpl", output_dir: "/"}, // a dot at the start breaks this
// 	{template_path: "templates/init/gitignore.tmpl", output_dir: "/"},
// 	{template_path: "templates/init/go.mod.tmpl", output_dir: "/"},
// 	{template_path: "templates/init/Dockerfile.tmpl", output_dir: "/"},
// 	{template_path: "templates/init/settings.go.tmpl", output_dir: "/settings/"},
// 	{template_path: "templates/init/env.go.tmpl", output_dir: "/settings/"},
// 	{template_path: "templates/init/database.go.tmpl", output_dir: "/settings/database/"},
// 	{template_path: "templates/init/response.go.tmpl", output_dir: "/lib/response/"},
// 	{template_path: "templates/init/validate.go.tmpl", output_dir: "/lib/validate/"},
// 	{template_path: "templates/init/router.go.tmpl", output_dir: "/app/"},
// 	{template_path: "templates/init/server.go.tmpl", output_dir: "/app/"},
// 	{template_path: "templates/init/convert.go.tmpl", output_dir: "/lib/convert/"},
// }

var init_project_templates = []Template{
	{template_path: "templates/init/main.go.tmpl", output_dir: "/cmd/api/"}, // main.go
	{template_path: "templates/init/args.go.tmpl", output_dir: "/cmd/"},
	{template_path: "templates/init/README.md.tmpl", output_dir: "/"},
	{template_path: "templates/init/env.tmpl", output_dir: "/"}, // a dot at the start breaks this
	{template_path: "templates/init/gitignore.tmpl", output_dir: "/"},
	{template_path: "templates/init/go.mod.tmpl", output_dir: "/"},
	{template_path: "templates/init/settings.go.tmpl", output_dir: "/pkg/settings/"},
	{template_path: "templates/init/env.go.tmpl", output_dir: "/pkg/settings/"},
	{template_path: "templates/init/database.go.tmpl", output_dir: "/pkg/settings/database/"},
	{template_path: "templates/init/response.go.tmpl", output_dir: "/pkg/response/"},
	{template_path: "templates/init/validate.go.tmpl", output_dir: "/pkg/validate/"},
	{template_path: "templates/init/router.go.tmpl", output_dir: "/pkg/server/"},
	{template_path: "templates/init/server.go.tmpl", output_dir: "/pkg/server/"},
	{template_path: "templates/init/convert.go.tmpl", output_dir: "/pkg/convert/"},
	{template_path: "templates/init/setup.go.tmpl", output_dir: "/pkg/environment/"},
}

var optional_sql_templates = []Template{
	{template_path: "templates/optional/sql/functions.sql.tmpl", output_dir: "/db/sql/"},
	{template_path: "templates/optional/sql/sqlc.yaml.tmpl", output_dir: "/db/"},
}
