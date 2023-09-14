package marvin

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/tompston/gomarvin/marvin/utils/convert"
)

//go:embed templates/init/*
//go:embed templates/module/*
//go:embed templates/optional/sql/*
//go:embed templates/optional/client/*

var EMBEDDED_TEMPLATES embed.FS

func executeTemplate(templateName, templatePath, fullOutputPath string, data interface{}) {

	tmpl, err := template.New(templateName).Funcs(TEMPLATE_FUNCTIONS).ParseFS(EMBEDDED_TEMPLATES, templatePath)
	if err != nil {
		log.Fatalf("Could not parse struct template: %v\n", err)
	}

	out, err := os.Create(fullOutputPath)
	if err != nil {
		log.Fatalf("Could not create file: %v\n", err)
	}
	defer out.Close()

	if err := tmpl.Execute(out, data); err != nil {
		log.Fatalf("Could not execute template: %v\n", err)
	}
}

// generate 2 strings from the template path, that are needed to run the executeTemplate() func
//   - temp_name = return everything that is after the last / to get the name of the template.
func createTemplateAndOutputName(templatePath string) (string, string) {
	templateName := regexp.MustCompile(`[^/]*$`).FindString(templatePath)
	out := strings.TrimSuffix(templateName, filepath.Ext(templateName))
	return templateName, out
}

// pass down config and array of predefined tempates and generate them
func generateTemplates(conf Config, templates []Template) {
	for i := 0; i < len(templates); i++ {
		outputDir := templates[i].outputDir
		templatePath := templates[i].templatePath
		generateSingleTemplate(conf, templatePath, outputDir)
	}
}

// pass down template data, full path to template and output dir
//   - generateSingleTemplate(conf, "templates/optional/ts/gomarvin.gen.ts.tmpl", "/ts/")
func generateSingleTemplate(conf Config, templatePath, outputDir string) {
	projectName := conf.ProjectInfo.Name
	projectOutputDir := fmt.Sprintf("./%s%s", projectName, outputDir)
	templateName, outputFile := createTemplateAndOutputName(templatePath)
	full_output_path := fmt.Sprintf("%s%s", projectOutputDir, outputFile)
	executeTemplate(templateName, templatePath, full_output_path, conf)
	fmt.Println(CREATED_MSG, full_output_path)
}

// includesRequired checks if the input string contains a "required" value
func includesRequired(x string) bool {
	return strings.Contains(x, "required")
}

// if the validate string includes a validate property, return the default field.
// Else return the field as an optional param.
func TypescriptField(fieldName, validateField string) string {
	if includesRequired(validateField) {
		return fieldName
	}
	return fmt.Sprintf("%v?", fieldName)
}

// if the validate string includes a validate property, return the default field.
// Else return the field as an optional param.
func PythonDataclasstField(fieldName, validateField string) string {
	if includesRequired(validateField) {
		return fieldName
	}
	return fieldName
}

func BodyTypeToGoStructType(bodyType string) string {
	if bodyType == "any" {
		return "interface{}"
	}
	return bodyType
}

func EndpointHasFieldWithTime(endpoints []Endpoints) bool {
	for _, endpoint := range endpoints {
		for _, bodyField := range endpoint.Body {
			if bodyField.Type == "time.Time" {
				return true
			}
		}
	}
	return false
}

var TEMPLATE_FUNCTIONS = template.FuncMap{
	"ConvertToTitle":                      convert.ToTitle,
	"ConvertToLowercase":                  convert.ToLowercase,
	"ConvertToUppercase":                  convert.ToUppercase,
	"ConvertToLowercaseTitle":             convert.ToLowercaseTitle,
	"ConvertToCamelCase":                  convert.ToCamelCase,
	"ConvertToPlural":                     convert.ToPlural,
	"ConvertToLowercasePlural":            convert.ToLowercasePlural,
	"WrapInCurlyBraces":                   convert.WrapInCurlyBraces,
	"WrapInCurlyBracesWithAppendedString": convert.WrapInCurlyBracesWithAppendedString,
	"ConvertLastCharTo":                   convert.LastCharTo,
	// other util functions
	"TypescriptField":          TypescriptField,
	"PythonDataclasstField":    PythonDataclasstField,
	"ApiVersion":               ApiVersion,
	"BodyTypeToGoStructType":   BodyTypeToGoStructType,
	"EndpointHasFieldWithTime": EndpointHasFieldWithTime,
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
