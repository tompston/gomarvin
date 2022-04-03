package marvin

// predefined name of config file that will be looked up if no other config path is specified
var ConfigName = "gomarvin.json"

// credits -> https://mholt.github.io/json-to-go/
type Config struct {
	ProjectInfo ProjectInfo `json:"project_info"`
	Modules     []Modules   `json:"modules"`
}
type ProjectInfo struct {
	GoVersion       float64 `json:"go_version"`
	Name            string  `json:"name"`
	Framework       string  `json:"framework"`
	Port            int     `json:"port"`
	APIPrefix       string  `json:"api_prefix"`
	ConfigVersion   float64 `json:"config_version"`
	DbType          string  `json:"db_type"`
	IncludeSQL      bool    `json:"include_sql"`
	IncludeTsFetch  bool    `json:"include_ts_fetch"`
	GomarvinVersion string  `json:"gomarvin_version"`
}
type URLParams struct {
	Field string `json:"field"`
	Type  string `json:"type"`
}
type Body struct {
	Field    string `json:"field"`
	Type     string `json:"type"`
	Validate string `json:"validate"`
}
type Endpoints struct {
	URL            string      `json:"url"`
	Method         string      `json:"method"`
	ControllerName string      `json:"controller_name"`
	URLParams      []URLParams `json:"url_params"`
	Body           []Body      `json:"body"`
}
type Modules struct {
	Name      string      `json:"name"`
	Endpoints []Endpoints `json:"endpoints"`
}

// self defined
type Template struct {
	template_path string
	output_dir    string // output dir into which the template will be created
}
type Project struct { // same as Config struct, only difference is that holds one Module
	ProjectInfo ProjectInfo `json:"project_info"`
	Modules     Modules     `json:"modules"`
}
