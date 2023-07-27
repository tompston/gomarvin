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
//go:embed templates/optional/client/*

var templates embed.FS

func ExecuteTemplate(templateName, templatePath, fullOutputPath string, data interface{}) {

	tmpl, err := template.New(templateName).Funcs(template_functions).ParseFS(templates, templatePath)
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

// generate 2 strings from the template path, that are needed to run the ExecuteTemplate() func
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
	ExecuteTemplate(templateName, templatePath, full_output_path, conf)
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

var template_functions = template.FuncMap{
	// name the template function the same as the imported convert function for predictability
	"ConvertToTitle":                      conv.ConvertToTitle,
	"ConvertToLowercase":                  conv.ConvertToLowercase,
	"ConvertToUppercase":                  conv.ConvertToUppercase,
	"ConvertToLowercaseTitle":             conv.ConvertToLowercaseTitle,
	"ConvertToCamelCase":                  conv.ConvertToCamelCase,
	"ConvertToPlural":                     conv.ConvertToPlural,
	"ConvertToLowercasePlural":            conv.ConvertToLowercasePlural,
	"WrapInCurlyBraces":                   conv.WrapInCurlyBraces,
	"WrapInCurlyBracesWithAppendedString": conv.WrapInCurlyBracesWithAppendedString,
	"ConvertLastCharTo":                   conv.ConvertLastCharTo,
	// other util functions
	"TypescriptField":        TypescriptField,
	"PythonDataclasstField":  PythonDataclasstField,
	"ApiVersion":             ApiVersion,
	"BodyTypeToGoStructType": BodyTypeToGoStructType,
}
