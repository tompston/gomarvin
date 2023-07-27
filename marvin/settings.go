package marvin

const (
	// message printed when a file is created
	CREATED_MSG = "\033[32m * CREATED\033[0m"

	// predefined name of config file that will be looked up if no other config path is specified
	ConfigName = "gomarvin.json"

	// output text if no args are provided
	gomarvin_info = `Usage:
	  gomarvin generate
	
	Version: 0.10.0
	
	Online Editor:
	  https://gomarvin.pages.dev
	
	Flags:
	  -config		
		Specify path to the gomarvin config file  (default "gomarvin.json")
	  -dangerous-regen
		  Regenerate everything. If set to true, all previous changes will be lost  (default "false")
	  -gut
		generate functions which convert possible api response structs to typescript interfaces (default "false")`
)
