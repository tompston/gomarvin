### Ideas

- Include faker?
- Add form / xml fields for go structs?

```bash
reflex -r '\.go' -s -- sh -c "go run main.go"
```

### codegen

- [x] endpoints would have automatic validation if Body exists
- [x] Validate fields are written from the frontend
- [x] All we do for the most part is generate those functions that then can be copied and modified. Using the generated file and editing it could be bad because then we limit how the user wants to sturcture the project
- [x] Automatically generate Struct files for Go and Interface files for Typescript
- [x] Routers for modules are imported only once, on the first run. That way we can test if they work correctly
- [x] Need to also add the gomarvin_version in project info.
- [x] import something from the validate package to avoid error when there are no bodies for endpoints (like when generating and not editing basic crud endpoints)
- [x] SQLC sql files are also regenerated on each run. To avoid errors, we use .txt files to store the generated code. That way user does not have to worry about further regens
- [x] On gen init, if include_sql is true, also add sqlc files
  - sqlc.yaml (incliding readme info as comments )
  - /sql/functions.sql so that the update thing works
- [x] rename all regenerated files to have .gen in the filename
- [x] add time.Time field for body field params
- [x] add include_ts to project_info
- [x] add types to generated TS interfaces
- [x] update validator package version
- [x] whitespace in validate struct will break the app, so remove them
- [x] test regex validation for email with validator
- [x] explain init project structure in README.md
- [x] add the new body types to TS interfaces
- [x] add mysql + sqlite support (functions.sql, sqlc.yaml, database.go, go.mod)
- [ ] finish the TS functions that fetch the endpoints
- [ ] refactor functions that generate templates with one module at a time
- [ ] Ts endpoints key is interface
- [ ] Expand postman tests
- [ ] figure out the next framework to add once this is done

### Frontend

- [x] Saving whole object to localstorage on change
- [x] Individual boolean switches for endpoints
- [x] create a thing that creates basic crud endpoints
- [x] syntax ERR when no modules
- [x] Validation with v-model values
- [x] Need a way to delete things also
- [x] Sort endpoints based on http method
- [x] write a func that deals with storing and validating config in localstorage
- [x] test Import
- [x] settings is a tab
- [x] refactor delete functions
- [x] make things look nicer
- [x] Create a new example with user module (registration, etc)
- [ ] adjust styles for mobile
- [ ] add more TS support


<!--

## References

- [sqlite example](https://github.com/bopbi/simple-todo/blob/master/simple-todo.go)
- [Go types](https://golangbyexample.com/all-basic-data-types-golang/)
- [fiber bodyparser thing]( https://docs.gofiber.io/api/ctx#bodyparser )
- [software versioning]( https://stackoverflow.com/questions/2864448/best-practice-software-versioning )

- [ ] gitignore .gen files

 -->
