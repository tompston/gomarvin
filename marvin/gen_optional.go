package marvin

import "fmt"

func GenerateOptional(conf Config, cmd Flags) {

	if conf.ProjectInfo.IncludeSQL { // if include_sql = true
		GenerateSql(conf, cmd)
	}

	if conf.ProjectInfo.IncludeFetch {
		GenerateFetchFunctions(conf, cmd)
	}
}

// if there are modules in the config, generate gomarvin.gen.ts in the /public dir
func GenerateFetchFunctions(conf Config, cmd Flags) {

	public_dir := fmt.Sprintf("./%s/public/", conf.ProjectInfo.Name)

	// if the public dir does not exist, create it
	if !PathExists(public_dir) {
		CreateDir(public_dir)
	}
	// if there are modules in the config file, create the typescript file
	if len(conf.Modules) != 0 {
		GenerateSingleTemplate(conf, typescript_fetch_template, "/public/")
	}
}

// if there are modules in the config, generate gomarvin.gen.ts
func GenerateOnlyFetchFunctions(conf Config, cmd Flags) {
	if len(conf.Modules) != 0 {
		template_name, output_file := GenerateTemplateAndOutputName(typescript_fetch_template)
		ExecuteTemplate(template_name, typescript_fetch_template, "./gomarvin.gen.ts", conf)
		fmt.Println(CREATED_MSG, output_file)
	}
}

// func GenerateSeeder(conf Config, cmd Flags) {
// 	if len(conf.Modules) != 0 {
// 		template_name, output_file := GenerateTemplateAndOutputName(typescript_fetch_template)
// 		ExecuteTemplate(template_name, typescript_seed_template, "./seeder.gen.ts", conf)
// 		fmt.Println(CREATED_MSG, output_file)
// 	}
// }
