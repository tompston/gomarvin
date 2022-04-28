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

	conv "github.com/tompston/gomarvin/marvin/utils/convert"
)

//go:embed templates/init/*
//go:embed templates/module/*
//go:embed templates/optional/sql/*
//go:embed templates/optional/ts/*

var templates embed.FS

func ExecuteTemplate(t_name string, t_path string, full_output_path string, data interface{}) {

	tmpl, err := template.New(t_name).Funcs(template_functions).ParseFS(templates, t_path)

	if err != nil {
		log.Fatalf("Could not parse struct template: %v\n", err)
	}

	out, _ := os.Create(full_output_path)
	defer out.Close()

	if err != nil {
		log.Fatalf("Could not parse struct template: %v\n", err)
	}

	tmpl.Execute(out, data)
}

// generate 2 strings from the template path, that are needed to run the ExecuteTemplate() func
// temp_name == return everything that is after the last / to get the name of the template.
func GenerateTemplateAndOutputName(template_path string) (string, string) {
	temp_name := regexp.MustCompile(`[^/]*$`).FindString(template_path)
	out := strings.TrimSuffix(temp_name, filepath.Ext(temp_name))
	return temp_name, out
}

// pass down config and array of predefined tempates and generate them
func GenerateTemplates(conf Config, templates []Template) {
	for i := 0; i < len(templates); i++ {

		output_dir := templates[i].output_dir
		templ_path := templates[i].template_path
		GenerateSingleTemplate(conf, templ_path, output_dir)
	}
}

// pass down template data, full path to template and output dir
// 	* GenerateSingleTemplate(conf, "templates/optional/ts/gomarvin.gen.ts.tmpl", "/ts/")
func GenerateSingleTemplate(conf Config, template_path string, output_dir string) {
	project_name := conf.ProjectInfo.Name
	project_output_dir := fmt.Sprintf("./%s%s", project_name, output_dir)
	template_name, output_file := GenerateTemplateAndOutputName(template_path)
	full_output_path := fmt.Sprintf("%s%s", project_output_dir, output_file)
	ExecuteTemplate(template_name, template_path, full_output_path, conf)
	fmt.Println(CREATED_MSG, full_output_path)
}

func GenerateTemplatesFromModules(conf Config, template_path string, output_dir string) {}

func WrapInCurlyBraces(x string) string {
	return fmt.Sprintf("{" + x + "}")
}

var template_functions = template.FuncMap{
	// name the template function the same as the imported convert function for predictability
	"ConvertToTitle":           conv.ConvertToTitle,
	"ConvertToLowercase":       conv.ConvertToLowercase,
	"ConvertToUppercase":       conv.ConvertToUppercase,
	"ConvertToLowercaseTitle":  conv.ConvertToLowercaseTitle,
	"ConvertToCamelCase":       conv.ConvertToCamelCase,
	"ConvertToPlural":          conv.ConvertToPlural,
	"ConvertToLowercasePlural": conv.ConvertToLowercasePlural,
	"WrapInCurlyBraces":        WrapInCurlyBraces,
}

const REPLACABLE_TEMPLATE_NAME = "__module__"

const typescript_fetch_template = "templates/optional/ts/gomarvin.gen.ts.tmpl"
