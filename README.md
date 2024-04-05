<div align="center"><img src="resources/img/logo_with_name.png" height="100px"/></div>

<h2 align="center> APiER: API Open Management Platform </h2>

**Chinese Version**: [中文](README.zh.md)

## Project Introduction

Web API unified management open platform - APiER platform

## Environment preparation

If you find that Go get fails, you can consider the following actions:

Open your terminal and execute

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

### macOS or Linux

Open your terminal and execute

```shell
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```

or

```shell
echo "export GO111MODULE=on" >> ~/.profile
echo "export GOPROXY=https://goproxy.cn" >> ~/.profile
source ~/.profile
```

### Windows

Open your PowerShell and execute

```shell
C:\> $env:GO111MODULE = "on"
C:\> $env:GOPROXY = "https://goproxy.cn"
```

or

```text
1. Open Start and search for "env"
2. Select "Edit System Environment Variables"
3. Click the "Environment Variables..." button
4. Under the "User Variables for <yourusername>" section (top half)
5. Click the “New…” button
6. Select the "Variable Name" input box and enter "GO111MODULE"
7. Select the "Variable value" input box and enter "on"
8. Click the "OK" button
9. Click the “New…” button
10. Select the "Variable Name" input box and enter "GOPROXY"
11. Select the "Variable Value" input box and enter "https://goproxy.cn"
12. Click the "OK" button
```

## Run the project

### Method 1: Run directly

Directly execute `go build main.go`

### Method 2: Use `air` to run (recommended in debug mode)

#### Step 1: Install `air`

```shell
go install github.com/cosmtrek/air@latest
```

#### Step 2: Initialize the `.toml` file

```shell
air init
```

#### Step 3: Select the `.toml` file to run

Multiple `.toml` files can be set, such as business background service [.air.api.toml](.air.api.toml), background
management service [.air.web.toml](.air.web.toml)

#### Step 4: Start the project

Take starting the background management service as an example. The toml file of the background management service
is [.air.web.toml](.air.web.toml). Then, the running command is run in the project root directory:

```shell
air -c .air.web.toml
```

