package marvin

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

// https://golangdocs.com/command-line-arguments-in-golang
type Flags struct {
	DangerousRegen bool   // true | false	( default = false )
	FetchOnly      bool   // true | false	( default = false )
	Gut            bool   // true | false	( default = false )
	ConfigPath     string // if no params are  specified, get the config from current dir
}

// https://paulgorman.org/technical/blog/20171113164018.html
func ReadConfig(config_path string) Config {
	c := flag.String("c", config_path, "Specify the configuration file.")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		log.Fatal("can't open config file: ", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	Config := Config{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}

	return Config
}
