# APiER: API Open Management Platform

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

Finish.

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

Finish.

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

Finish.