# Gomarvin changelog

<!--
- added optional cmd arg for generating only the typescript fetch functions file
- renamed generated main.gen.ts file to be gomarvin.gen.ts
- gitignore *.gen.txt (generated sql text files)
- Improved generated fetch functions
  - JSDOC support
  - new variable that groups the fetch functions by the module 

-->

### v0.3.0

- **Support for Chi framework**
- `controllers.gen.go` added to gitignore, as it is meant to be a placeholder file that should be renamed
- controllers don't call the `database.GetDbConn()` function anymore. Instead, use settings.DB variable.
  - refactored to shorten the lenght of controllers
- renamed `.env` db variables
- added a bash command that can rename generated sql files in the `/db/sql/` dir
- responses return null data by default
- creating `/public/` dir on initial run so that it could
  - hold generated typescript fetch file so that it could be shared by frontend easier.
  - hold the gomarvin config
  - optionally host config files ( implemented manually )

### v0.2.0

- **Support for echo framework**
- refactored typescript functions to fetch endpoints
- automatic formatting after codegen using gofmt

### v0.1.0

- init release
