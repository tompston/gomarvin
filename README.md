<h4 align="center">
<img src="./assets/gomarvin.svg" height="70">

<h2 align="center">Generate boilerplate for Gin / Fiber / Echo / Chi REST servers.</h2>

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
import * as F from "../../../gomarvin.gen"; // import the generated file
import { CommentEndpoints } from "../../../gomarvin.gen"; // or import only the Comment module endpoints
import { GetUserById } from "../../../gomarvin.gen"; // or just import a single fetch function

// fetch a user by id
async function FetchGetUserByIdEndpoint() {
  let res = await GetUserById(1);
  let users = await res.json();
  console.log(users);
}

// create a new user
async function FetchCreateUserEndpoint() {
  let res = await F.CreateUser({
    username: "qweqwe",
    email: "qwe@qwe.com",
    age: 20,
    password: "very-long-and-good-password",
  });

  let user = await res.json();
  console.log(user);
}

// append optional string to the existing endpoint url
async function FetchEndpointWithAppendedUrl() {
  const res = await F.GetUserById(10, { append_url: "?name=jim" });
  console.log(res);
}

// Use both optional values
// - append a string to the fetch url
// - define a new options object used in the fetch request
async function FetchWithAppendedUrlAndCustomOptions() {
  const res = await F.GetUserById(10, {
    options: { method: "DELETE" },
    append_url: "?name=jim",
  });
  console.log(res);
}

// Fetch a single endpoint from the Comment module
async function FetchCommentById() {
  const res = await CommentEndpoints.GetComment(20);
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
go run main.go -dangerous-regen="true" -config="./examples/v0.4.0/gomarvin-chi_with_modules.json" generate
cd chi_with_modules
go mod tidy
go mod download
code .
cd ..

# generate only the fetch client thing
go run main.go              \
    -config="./examples/v0.4.0/gomarvin-chi_with_modules.json" \
    -fetch-only="true" generate

// uncomment lower line to call generated sqlc functions with db connection
// var Queries = sqlc.New(database.DB)

git add .
git commit -m "next"
git push

GOOS=darwin GOARCH=arm64 go build -o gomarvin main.go

# release
git add .
git commit -m "gomarvin: release v0.4.1"
git tag v0.4.1
git push origin v0.4.1
GOPROXY=proxy.golang.org go list -m github.com/tompston/gomarvin@v0.4.1

-->
