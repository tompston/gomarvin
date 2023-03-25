package marvin

// credits -> https://mholt.github.io/json-to-go/

// predefined name of config file that will be looked up if no other config path is specified
var ConfigName = "gomarvin.json"

type Config struct {
	// Config File
	ProjectInfo ProjectInfo `json:"project_info"`
	Modules     []Modules   `json:"modules"`
}

type Project struct {
	// same as Config struct, only difference is that holds one Module
	ProjectInfo ProjectInfo `json:"project_info"`
	Modules     Modules     `json:"modules"`
}

type ProjectInfo struct {
	// as the go version can be 1.20, storing it as interface{}
	// is the best way to handle it, because javascript toFixed()
	// function returns a string.
	GoVersion       interface{} `json:"go_version"`
	ConfigVersion   interface{} `json:"config_version"`
	Name            string      `json:"name"`
	Framework       string      `json:"framework"`
	Port            int         `json:"port"`
	APIPrefix       string      `json:"api_prefix"`
	DbType          string      `json:"db_type"`
	IncludeSQL      bool        `json:"include_sql"`
	IncludeFetch    bool        `json:"include_fetch"`
	GomarvinVersion string      `json:"gomarvin_version"`
}

type Endpoints struct {
	// Endpoint info
	URL            string      `json:"url"`
	Method         string      `json:"method"`
	ControllerName string      `json:"controller_name"`
	ResponseType   string      `json:"response_type"`
	URLParams      []URLParams `json:"url_params"`
	Body           []Body      `json:"body"`
}

type Modules struct {
	Name      string      `json:"name"`
	Endpoints []Endpoints `json:"endpoints"`
}

type URLParams struct {
	Field string `json:"field"`
	Type  string `json:"type"`
}

type Body struct {
	// Endpoint Body field properties
	Field    string `json:"field"`
	Type     string `json:"type"`
	Validate string `json:"validate"`
}

// self defined
type Template struct {
	template_path string
	output_dir    string // output dir into which the template will be created
}
