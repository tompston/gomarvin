<h4 align="center">
<img src="./assets/gomarvin.svg" height="70">

<h2 align="center">Generate boilerplate for Gin / Fiber / Echo / Chi REST servers.</h2>

<div align="center"> <em>AKA Why spend 5 minutes doing it when I can spend 3 months automating it?</em> </div>

<br/>

https://user-images.githubusercontent.com/82293948/197339991-ca049896-7a0e-4fa9-ae2b-787f0bb021b2.mp4

Generate:

- Init server + controllers for endpoints
- Typescript fetch functions with type-checked payloads for generated endpoints
- SQL files with tables and queries for controllers ( Postgres )

from one config file.

## [Documentation](https://gomarvin.pages.dev/docs)

## Install

1.  Either create a custom config file [with the frontend editor](https://gomarvin.pages.dev/) or copy one of the config file from [examples dir](https://github.com/tompston/gomarvin/tree/main/examples)

2.  Install gomarvin

```bash
# go version >= 1.17
go install github.com/tompston/gomarvin@latest

# go version < 1.17
go get github.com/tompston/gomarvin

# or clone the repo and run go run main.go
git clone https://github.com/tompston/gomarvin.git
```

3. run gomarvin

```bash
# run this in the same dir as the config file, if the name of the config is "gomarvin.json"
gomarvin generate

# run this if custom config file name or path
gomarvin -config="PATH_TO_CONFIG" generate

# or generate only the typescript API client file. Useful if you want to generate fetch
# functions for a pre-existing REST API in a fast way.
gomarvin -fetch-only=true generate

# optionally generate the golang struct -> typescript interface converters
gomarvin -gut=true generate
```

4. run lower commands

```bash
cd GENERATED_SERVER
go mod tidy
go mod download
go run main.go
```

## CLI

```
Flags:
  -config		Specify path to the gomarvin config file (default "gomarvin.json")
  -dangerous-regen	Regenerate everything. If set to true, init server will be regenerated and all previous changes will be lost (default "false")
  -fetch-only		generate only the typescript file that holds fetch function (default "false")`
  -gut          generate additional file in the modules dir which concats all of the functions that convert possible response structs to typescript interfaces
```

### Generated Typescript fetch functions usage example

```js
// import the generated file
import * as gomarvin from "../../../chi_with_modules/public/gomarvin.gen" 
// or just import a single fetch function
import { GetUserById } from "../../../chi_with_modules/public/gomarvin.gen"

// either use the default client created from
// the settings of the config file, or create a new one
// (useful when switching environments)
const defaultClient = gomarvin.defaultClient

// api client when deployed
const productionClient: gomarvin.Client = {
  host_url: "http://example.com",
  api_prefix: "/api/v1",
  headers: {
    "Content-type": "application/json;charset=UTF-8",
  },
}

const DEV_MODE = true

// switch to productionClient if DEV_MODE is false
const client = DEV_MODE ? defaultClient : productionClient

// fetch GetUserById endpoint
async function FetchGetUsersById() {
  const res = await gomarvin.GetUserById(client, 10);
  console.log(res);
}

// append optional string to the existing endpoint url
async function FetchEndpointWithAppendedUrl() {
  const res = await gomarvin.GetUserById(client, 10, { append_url: "?name=jim" });
  console.log(res);
}

// define custom options for the fetch request
async function FetchEndpointWithCustomOptions() {
  const res = await gomarvin.GetUserById(client, 10, { options: { method: "POST" } });
  console.log(res);
}

// Use both optional values
// - append a string to the fetch url
// - define a new options object used in the fetch request
async function FetchWithAppendedUrlAndCustomOptions() {
  const res = await gomarvin.GetUserById(client, 10, {
    options: { method: "DELETE" },
    append_url: "?name=jim",
  });
  console.log(res);
}
```

### Disclaimer

I'm still figuring out how to do stuff better, so if there's a new release where the `y` in `x.y.z`  the version ( tag ) is incremented, that means that there are breaking changes. The current versions work, but there's a lot of improvements that can be done to make stuff better.

### Notes

- Gin crashes by default. This is due to the way the framework handles routing.

- bugs caused by url params and routing should be fixed manually.

- Seeding the db
  - As the generated fetch functions are mapped to the endpoints, you can use them to seed the database pretty easily. Using deno with the faker package can make the process trivial.
- If formatting does not work for some reason, run this

  ```
  gofmt -s -w .
  ```
- Using the tool to generate fetch functions for pre-existing REST APIs
  - If you want to generate a client for an already existing REST API in a fast way, you can use the editor to document the needed values to codegen the fetch functions in a fast way
  
  ```bash
  # run this after exporting the config file, assuming that
  # the gomarvin.json file is in the same dir
  gomarvin -fetch-only="true" generate
  ```

- If there are errors after creating a new config file, submit an issue. There might be some edge cases that have not been tested yet. 

### Credits to used packages

_Not installed locally to avoid any dependencies._

- [go-pluralize](https://github.com/gertd/go-pluralize)
- [strcase](https://github.com/iancoleman/strcase)


<!--


# run a local example
GOOS=darwin GOARCH=arm64 go build -o gomarvin main.go

# release
git add .
git commit -m "gomarvin: release v0.7.0"
git tag v0.7.0
git push origin v0.7.0
GOPROXY=proxy.golang.org go list -m github.com/tompston/gomarvin@v0.7.0


# create a new branch
git branch BRANCH_NAME
# switch to branch
git checkout BRANCH_NAME

# merge new branch to main branch
git checkout main
git merge v0.7.x


-->