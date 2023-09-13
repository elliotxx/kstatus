<div align = "center">
<!-- <p> -->
<!--     <img width="160" src="https://github.com/elliotxx/kstatus/blob/main/example_logo.png?sanitize=true"> -->
<!-- </p> -->
<h2>Low Dependency/Scalable Kubernetes Resource Status Check Library</h2>
<a title="Go Reference" target="_blank" href="https://pkg.go.dev/github.com/elliotxx/kstatus"><img src="https://pkg.go.dev/badge/github.com/elliotxx/kstatus.svg"></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/elliotxx/kstatus"><img src="https://goreportcard.com/badge/github.com/elliotxx/kstatus?style=flat-square"></a>
<a title="Coverage Status" target="_blank" href="https://coveralls.io/github/elliotxx/kstatus?branch=main"><img src="https://img.shields.io/coveralls/github/elliotxx/kstatus/main"></a>
<a title="Code Size" target="_blank" href="https://github.com/elliotxx/kstatus"><img src="https://img.shields.io/github/languages/code-size/elliotxx/kstatus.svg?style=flat-square"></a>
<br>
<a title="GitHub release" target="_blank" href="https://github.com/elliotxx/kstatus/releases"><img src="https://img.shields.io/github/release/elliotxx/kstatus.svg"></a>
<a title="License" target="_blank" href="https://github.com/elliotxx/kstatus/blob/main/LICENSE"><img src="https://img.shields.io/github/license/elliotxx/kstatus"></a>
</p>
</div>

This is a low dependency/scalable Kubernetes resource status check library that performs magic modifications based on [kubernetes-sigs/cli-utils/kstatus](https://github.com/kubernetes-sigs/cli-utils/blob/master/pkg/kstatus/README.md).

## 📜 Language

[English](https://github.com/elliotxx/kstatus/blob/main/README.md) | [简体中文](https://github.com/elliotxx/kstatus/blob/main/README-zh.md)


## ✨ Core Feature
* ⚡ Low Dependency
* 🌲 Scalable


## ⚙️ Usage
```shell
go get github.com/elliotxx/kstatus
```


## 📖 Example
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

Output:

```shell
{
    "type": "Reconciling",
    "status": "True",
    "reason": "LessUpdated",
    "message": "Updated: 0/1"
}
```

More examples: [./example_test.go](./example_test.go)
