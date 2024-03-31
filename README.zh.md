# APiER: API 管理开放平台

**Chinese Version**: [中文](README.zh.md)

## 项目介绍

Web API 统一管理开放平台——阿皮尔（APiER）平台

## 环境准备

如果发现Go get失败，可以考虑以下操作：

打开你的终端并执行

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

完成。

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

完成。

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

完成。