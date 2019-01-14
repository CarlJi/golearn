package work

import (
	"os"

	"carlji.com/experiments/gocover/go1.11.2/pkg/load"
)

// recompile package with cover stub

// TODO: 注意查看go build的pkgdir参数，也许可以利用下

func gocoverInit() {
	if os.Getenv("GO_COVER") != "" {
		pkgsFilter = pkgsCoverMarkedAsNeed
	}
}

// pkgsCoverMarkedAsNeed mark packages with cover flag
// except packages for system and packages in vendor
func pkgsCoverMarkedAsNeed(pkgs []*load.Package) []*load.Package {

	return pkgs
}
