<!-- # [gomarvin -- links to documentation]() -->

<h4 align="center">
<img src="./assets/gomarvin.svg" width="84">

<!-- <h1 align="center">
gomarvin -->

<h3 align="center">

Generate controllers for Gin / Fiber REST servers.

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

# to format the project after each codegen, also run this
gofmt -s -w .
```

4. run lower commands

```bash
cd GENERATED_SERVER
go mod tidy
go mod download
gofmt -s -w .
go run main.go
```

## CLI

```
gomarvin -h

-config
      Specify path to the gomarvin config file (default "gomarvin.json")
-dangerous_regen
      Regenerate everything. If set to true, init server will be regenerated and  all previous changes will be lost (default "false")
```

<!--
git add .
git commit -m "next"
git push

GOOS=darwin GOARCH=arm64 go build -o gomarvin main.go

git add .
git commit -m "gomarvin: init release v0.1.0"
git tag v0.1.0

 -->
