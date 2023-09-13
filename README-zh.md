<div align = "center">
<!-- <p> -->
<!--     <img width="160" src="https://github.com/elliotxx/kstatus/blob/main/example_logo.png?sanitize=true"> -->
<!-- </p> -->
<h2>低依赖/可扩展的 Kubernetes 资源状态检查库</h2>
<a title="Go Reference" target="_blank" href="https://pkg.go.dev/github.com/elliotxx/kstatus"><img src="https://pkg.go.dev/badge/github.com/elliotxx/kstatus.svg"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/elliotxx/kstatus"><img src="https://goreportcard.com/badge/github.com/elliotxx/kstatus?style=flat-square"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/github/elliotxx/kstatus?branch=main"><img src="https://img.shields.io/coveralls/github/elliotxx/kstatus/main"></a>
<a title="Code Size" target="_blank" href="https://github.com/elliotxx/kstatus"><img src="https://img.shields.io/github/languages/code-size/elliotxx/kstatus.svg?style=flat-square"></a>
<br>
<a title="GitHub release" target="_blank" href="https://github.com/elliotxx/kstatus/releases"><img src="https://img.shields.io/github/release/elliotxx/kstatus.svg"></a>
<a title="License" target="_blank" href="https://github.com/elliotxx/kstatus/blob/main/LICENSE"><img src="https://img.shields.io/github/license/elliotxx/kstatus"></a>
</p>
</div>

这是一个低依赖/可扩展的 Kubernetes 资源状态检查库,它基于 [kubernetes-sigs/cli-utils/kstatus](https://github.com/kubernetes-sigs/cli-utils/blob/master/pkg/kstatus/README.md) 进行魔改.

## 📜 多语言

[English](https://github.com/elliotxx/kstatus/blob/main/README.md) | [简体中文](https://github.com/elliotxx/kstatus/blob/main/README-zh.md)


## ✨ 核心特性
* ⚡ 低依赖
* 🌲 可扩展


## ⚙️ 使用方法
```shell
go get github.com/elliotxx/kstatus
```


## 📖 样例
```go
package main

import (
	"github.com/elliotxx/kstatus"
)

func main() {
    // deployment := getDeploymentFromCluster()
    
    res, err := Compute(deployment)
    if err != nil {
        panic(err)
    }
    fmt.Println(toJSON(res))
}
```

输出:

```shell
{
    "type": "Reconciling",
    "status": "True",
    "reason": "LessUpdated",
    "message": "Updated: 0/1"
}
```

更多样例可以参考: [./example_test.go](./example_test.go)
