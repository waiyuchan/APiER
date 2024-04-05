<div align="center">
    <img src="resources/img/logo_with_name.png" height="100px"/>
    <h1>APiER: API 开放管理平台</h1>
</div>

<div align="center">
    <img src="https://img.shields.io/badge/language-Golang-blue.svg" />
    <img src="https://img.shields.io/badge/backend_frame-Gin-6db33f.svg" />
    <img src="https://img.shields.io/badge/ORM-Gorm-red.svg" />
</div>

<h4 align="center">
    中文 | <a href="README.md">English</a>
</h4>


阿皮尔接口开放管理平台（APiER API Open Management Platform）是一个先进的API管理与开放平台，旨在为企业和开发者提供一个高效、安全、易于使用的API集成、管理和共享环境。

## 环境准备

如果发现Go get失败，可以考虑以下操作：

打开你的终端并执行

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

### macOS 或 Linux

打开你的终端并执行

```shell
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```

或者

```shell
echo "export GO111MODULE=on" >> ~/.profile
echo "export GOPROXY=https://goproxy.cn" >> ~/.profile
source ~/.profile
```

### Windows

打开你的 PowerShell 并执行

```shell
C:\> $env:GO111MODULE = "on"
C:\> $env:GOPROXY = "https://goproxy.cn"
```

或者

```text
1. 打开“开始”并搜索“env”
2. 选择“编辑系统环境变量”
3. 点击“环境变量…”按钮
4. 在“<你的用户名> 的用户变量”章节下（上半部分）
5. 点击“新建…”按钮
6. 选择“变量名”输入框并输入“GO111MODULE”
7. 选择“变量值”输入框并输入“on”
8. 点击“确定”按钮
9. 点击“新建…”按钮
10. 选择“变量名”输入框并输入“GOPROXY”
11. 选择“变量值”输入框并输入“https://goproxy.cn”
12. 点击“确定”按钮
```

## 运行项目

### 方式一：直接运行

直接执行 `go build main.go`

### 方式二：使用 `air` 运行（调试模式下推荐使用）

#### 步骤一：安装 `air`

```shell
go install github.com/cosmtrek/air@latest
```

#### 步骤二：初始化 `.toml` 文件

```shell
air init
```

#### 步骤三：选择要运行的 `.toml` 文件

可以设置多个 `.toml` 文件，比如业务后台服务 [.air.api.toml](.air.api.toml)、 后台管理服务 [.air.web.toml](.air.web.toml)

#### 步骤四：启动项目

以启动后台管理服务为例，后台管理服务的toml文件是 [.air.web.toml](.air.web.toml)，那么，运行指令在项目根目录运行：

```shell
air -c .air.web.toml
```
