# Gomarvin changelog

<!-- 
v0.7.0
- DbConnErrorMessage() 
  - renamed to DbErrorMessage()
  - input type changed from string to error, to reduce the amount of chars written
- Removed the unused app value that was passed in the generated fiber Router__ functions
- Refactor the ValidateStruct function to return an error instead of a struct, so that validating the body could be done as a single function call in the controllers (less loc)
- If the validate field does not include an `required` value, the generated interface for the body type will also be an optional param
- the `body.gen.go` file will be generated only when the module has at least a single body for an endpoint
- Removed `ResponseWithPagination` function. Now the `Response` function takes in the 4th optional parameter. This makes stuff shorter.
- The generated Typescript client now has an interface that defines how the returned response should look like (see ApiResponse)
- * CREATED cli color is now green
- Removing panics in database package
- Refactored the way env variables are loaded 
  - Instead of storing calling them all as strings from the settings, save them to the settings.Environment variable (+ convert to the expected type) when the main() func is executed
- New convert package added, so that the `settings.SetEnvironmentConfig()` could convert the loaded env variables correctly.


TODOS ::

- New postam tests that check what happens when 
  - "{" sent in the payload
  - Invalid body type sent in the payolad
  - Empty object is sent for the TestOptionalParam endpoint
- Test what happens when there are 4 and 5 params in res.Response function (to check if pagination is not returned).
- Test the returend response if the body includes the incorrect type.

- Write more comments for the generated functions to explain what they do ( + examples that show how they work )
-  Figure out how to do tests for the controllers. Two possible options ->
   -  generate deno tests, using the generated fetch functions
   -  generate tests in go, in a file called `controllers_test.gen.go`
- Check out how echo implements the json validation


### Typescript

- Write a typescript func which checks if the data variable holds an interface that is returned is the error fields struct, when the validation fails.
- Maybe refactor the editor / config files to hold a comment value that explains what 
  the generated controllers do, and use that string as a comment before controllers
  and fetch functions.
- Figure out if you can also suppor url query params in some way. This could be slightly tricky tho


### Editor
  - Create a new page in the frontend app that links to the cloudflare builds for different versions
  - The import tab should ideally validate if the passed in string is a valid config file
    and only then allow the import
  - settings config_version throw error if the input is not a valid float

-->

### v0.6.0

- Improved SQL / SQLC queries
  - Better naming
  - Query for cursor pagination with `module_id` and `created_at`
    - Credits / References
      - https://morningcoffee.io/stable-pagination.html
      - https://medium.easyread.co/how-to-do-pagination-in-postgres-with-golang-in-4-common-ways-12365b9fb528
- Renamed `DbConnErrorMessage` -> `DbErrorMessage`
- Refactored payload validation response to be more UI friendly
  - Credits -> https://github.com/go-playground/validator/issues/559
- Refactored pagination links struct
  - Removed Self field
  - omitempty if no value is provided
- Removed the placeholder block in modules that is used for removing possible import errors.
- Added a placeholder function that can be replaced to the function which runs the database query and returns the data + checks for error messages.


### v0.5.0

- Updated go.mod dependencies
- Refactored fetch client
  - Fetch functions now have an required argument which has the settings for the fetch request ( url and headers )

    This allows an easier switch between the dev and the prod environment. As `gomarvin.gen.ts` file is regenerated and ideally should not be edited, switching the api client settings between the dev and prod env can be done only manually, if the api settings are called from the pre-existing settings object. 
    
    If the settings of the client are passed as a argument to the function, that problem dissapears.

    - Note that this could be implemented as a class, but then the unused methods don't get purged, thus increasing the bundle size (at least in local tests)

  - JSDoc comments now include type annotations for the url params


### v0.4.1

- Body of the controller removed from the controller and pushed to a new file
  - this way, when updating the body fields from the editor, you don't need to update the controllers from scratch
- changed cli args to be more intuitive
  - generation works only if you include `generate`

### v0.4.0

- renamed `utils` to `lib`
- `settings` package holds variables from .env file (name of the var the same as the one used in .env file)
- Endpoint Body struct pushed to be inside controller (taken from Mat Ryer https://youtu.be/rWBSMsLG8po?t=1671)
- Block that sets ups the server function is pushed to be a seperate function
- Decoupled SQL query generation from controllers.
- refactored generated fetch function comments
  - moved fetch function opt param explanation to the interface.
  - API object now includes the `config_version` key that shows from
    which version the current file is generated.
  - if the fetch function has a body, hover over the interface in the JSDoc comments to see the required validation for the fields

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
