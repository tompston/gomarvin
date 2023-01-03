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

1.  Either create a custom config file [with the frontend editor](https://gomarvin.pages.dev/) or copy one of the config file from [examples dir](https://github.com/tompston/gomarvin/tree/main/examples/v0.3.0)

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
gomarvin -fetch-only="true" generate
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
```

### Generated Typescript fetch functions usage example

```js
// import the generated file
import * as F from "../../../chi_with_modules/public/gomarvin.gen" 
// or just import a single fetch function
import { GetUserById } from "../../../chi_with_modules/public/gomarvin.gen"

// either use the default client created from
// the settings of the config file, or create a new one
// (useful when switching environments)
const defaultClient = F.defaultClient

// api client when deployed
const productionClient: F.Client = {
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
  const res = await F.GetUserById(client, 10);
  console.log(res);
}

// append optional string to the existing endpoint url
async function FetchEndpointWithAppendedUrl() {
  const res = await F.GetUserById(client, 10, { append_url: "?name=jim" });
  console.log(res);
}

// define custom options for the fetch request
async function FetchEndpointWithCustomOptions() {
  const res = await F.GetUserById(client, 10, { options: { method: "POST" } });
  console.log(res);
}

// Use both optional values
// - append a string to the fetch url
// - define a new options object used in the fetch request
async function FetchWithAppendedUrlAndCustomOptions() {
  const res = await F.GetUserById(client, 10, {
    options: { method: "DELETE" },
    append_url: "?name=jim",
  });
  console.log(res);
}
```

### Notes

- If formatting does not work, run this

```
gofmt -s -w .
```

- If there are errors after creating a new config file, submit an issue. There might be some edge cases that have not been tested yet. (except for bugs caused by url params and routing, those should be fixed manually)</h3>

### Credits to used packages

_Not installed locally to avoid any dependencies._

- [go-pluralize](https://github.com/gertd/go-pluralize)
- [strcase](https://github.com/iancoleman/strcase)

#### Note on versioning

Versions (tags) have a pattern of `x.y.z`

- `x` is incremented on big releases.
- `y` is incremented when there is a change that breaks the previously generated servers.
- `z` is incremented when there is a change that doesn't break the previously generated servers.

For example, updating from `0.1.0` to `0.1.1` does not break anything. Going from `0.2.0` to `0.3.0` will break stuff.

<!--

# run a local example



// uncomment lower line to call generated sqlc functions with db connection
// var Queries = sqlc.New(database.DB)

git add .
git commit -m "next"
git push

GOOS=darwin GOARCH=arm64 go build -o gomarvin main.go

# release
git add .
git commit -m "gomarvin: release v0.5.0"
git tag v0.5.0
git push origin v0.5.0
GOPROXY=proxy.golang.org go list -m github.com/tompston/gomarvin@v0.5.0

# switch to pre-release dev branch
git checkout v0.6.x

-->
