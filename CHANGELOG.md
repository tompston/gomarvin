# Gomarvin changelog

<!-- ### v0.3.2

- refactored generated fetch function comments
  - moved fetch function opt param explanation to the interface.
  - API object now includes the `config_version` key that shows from 
    which version the current file is generated.
  - if the fetch function has a body, hover over the interface in the JSDoc comments to see the required validation for the fields
-->


### v0.3.1

- added optional cmd arg for generating only the typescript fetch functions file
- renamed generated `main.gen.ts` file to `gomarvin.gen.ts`
- gitignore \*.gen.txt (generated sql text files)
- fixed sqlc queries `:one`, `:many` syntax
- cli arg `dangerous_regen` renamed to `dangerous-regen`
- Improved generated fetch functions
  - JSDOC support
  - endpoint body interfaces documents the validation for the fields
  - new variable that groups the fetch functions by the module
    - this way you can import / find the correct fetch function easier
  - each fetch function has an optional object param that can be used when params need to be changed

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

- **Support for Echo framework**
- refactored typescript functions to fetch endpoints
- automatic formatting after codegen using gofmt

### v0.1.0

- init release
