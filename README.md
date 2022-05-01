<h4 align="center">
<img src="./assets/gomarvin.svg" height="70">

<h2 align="center">Generate boilerplate for Gin / Fiber / Echo / Chi REST servers.</h2>

Generate:

- Init server + controllers for endpoints
- Typescript fetch functions with type-checked payloads for generated endpoints
- SQL files with tables and queries for controllers ( Postgres )

from one config file

## [Documentation](https://gomarvin.pages.dev/docs)

## Install

1.  Either create a custom config file [with frontend editor](https://gomarvin.pages.dev/) or copy one of the config file from [examples dir](https://github.com/tompston/gomarvin/tree/main/examples)

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
gomarvin

# run this if custom config file name or path
gomarvin -config="PATH_TO_CONFIG"
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
gomarvin -h

-config string
      Specify path to the gomarvin config file (default "gomarvin.json")
-dangerous_regen string
      Regenerate everything. If set to true, init server will be regenerated and  all previous changes will be lost (default "false")
```

<!--
-fetch-only string
      generate only the typescript file that holds fetch function (default "false")
-->

### Generated Typescript fetch functions usage example

```js
// import the generated file
import * as F from "../../../gomarvin.gen";
// or import only the Comment module endpoints
import { CommentEndpoints } from "../../../gomarvin.gen";

// fetch a user by id
async function FetchGetUserByIdEndpoint() {
  let res = await F.GetUserById(1);
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

// define custom options for the fetch request
async function FetchEndpointWithCustomOptions() {
  const res = await F.GetUserById(10, { options: { method: "POST" } });
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

// Fetch a singe endpoint from the Comment module
async function FetchCommentById() {
  const res = await CommentEndpoints.GetComment(20);
  console.log(res);
}
```

### Notes

If formatting does not work, run this

```
gofmt -s -w .
```

#### Credits to used packages

- [go-pluralize](https://github.com/gertd/go-pluralize)
- [strcase](https://github.com/iancoleman/strcase)

<!--

# gen

examples/v0.3.0/gomarvin-fiber_with_modules.json

go run main.go -dangerous_regen="true" -config="./previous/gomarvin_-v0.3.0.json"
cd server_with_gin_next
go mod tidy
go mod download
gofmt -s -w .
code .
cd ..


// uncomment lower line to call generated sqlc functions with db connection
// var Queries = sqlc.New(database.DB)


git add .
git commit -m "next"
git push

GOOS=darwin GOARCH=arm64 go build -o gomarvin main.go

# release
git add .
git commit -m "gomarvin: release v0.3.0"
git tag v0.3.0
git push origin v0.3.0
GOPROXY=proxy.golang.org go list -m github.com/tompston/gomarvin@v0.3.0

-->
