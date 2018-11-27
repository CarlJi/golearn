
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
* 



------