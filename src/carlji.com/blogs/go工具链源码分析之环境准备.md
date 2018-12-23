

## 前言

* 复制cmd下go工具链源码

```
➜  go1.11.2 git:(master) ✗ go run main.go

main.go:18:2: use of internal package cmd/go/internal/base not allowed
main.go:19:2: use of internal package cmd/go/internal/bug not allowed
main.go:20:2: use of internal package cmd/go/internal/cfg not allowed
main.go:21:2: use of internal package cmd/go/internal/clean not allowed
```

* 将包名internal改为pkg
* 接下来要准备两个工具:
* 1. 方便打印log， 我选用 	"qiniupkg.com/x/log.v7"， 方便快速定位到代码行数
* 2. 优雅打印go struct工具 go-spew
*    go get -u github.com/davecgh/go-spew/spew



------