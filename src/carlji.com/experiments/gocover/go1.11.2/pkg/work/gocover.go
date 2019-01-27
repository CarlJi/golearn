package work

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"carlji.com/experiments/gocover/go1.11.2/pkg/load"
	"qiniupkg.com/x/log.v7"
)

// recompile package with cover stub

// TODO: 注意查看go build的pkgdir参数，也许可以利用下
func goCoverInit() {
	if os.Getenv("GO_COVER") != "" {
		pkgsFilter = pkgsCoverMarkedAsNeed
	}
}

// pkgsCoverMarkedAsNeed mark packages with cover flag
// except packages for system and packages in vendor
func pkgsCoverMarkedAsNeed(pkgs []*load.Package) []*load.Package {
	log.Printf("COVER FLAG!!")
	newPkgs := load.PackageList(pkgs)
	for _, p := range newPkgs {
		// TODO: do not execute coverage analysis on Golang internal packages
		if strings.Contains(p.ImportPath, "covs") {
			log.Printf("add cover analysis to file: %s", p.ImportPath)
			p.Internal.CoverMode = "count"
			p.Internal.CoverVars = declareCoverVars(p, p.GoFiles...)
		}
	}

	return newPkgs
}

// addCoverageAnalysis automatically export two api routes in the main package in order to
// support code coverage analysis
// API: /cover/profiles 返回当前覆盖率详情
// API: /cover/rate 返回当前覆盖率百分比
func addCoverageAnalysis(p *load.Package) {

}

// isTestFile reports whether the source file is a set of tests and should therefore
// be excluded from coverage analysis.
func isTestFile(file string) bool {
	// We don't cover tests, only the code they test.
	return strings.HasSuffix(file, "_test.go")
}

// declareCoverVars attaches the required cover variables names
// to the files, to be used when annotating the files.
func declareCoverVars(p *load.Package, files ...string) map[string]*load.CoverVar {
	coverVars := make(map[string]*load.CoverVar)
	coverIndex := 0
	// We create the cover counters as new top-level variables in the package.
	// We need to avoid collisions with user variables (GoCover_0 is unlikely but still)
	// and more importantly with dot imports of other covered packages,
	// so we append 12 hex digits from the SHA-256 of the import path.
	// The point is only to avoid accidents, not to defeat users determined to
	// break things.
	sum := sha256.Sum256([]byte(p.ImportPath))
	h := fmt.Sprintf("%x", sum[:6])
	for _, file := range files {
		if isTestFile(file) {
			continue
		}

		// For a package that is "local" (imported via ./ import or command line, outside GOPATH),
		// we record the full path to the file name.
		// Otherwise we record the import path, then a forward slash, then the file name.
		// This makes profiles within GOPATH file system-independent.
		// These names appear in the cmd/cover HTML interface.
		var longFile string
		if p.Internal.Local {
			longFile = filepath.Join(p.Dir, file)
		} else {
			longFile = path.Join(p.ImportPath, file)
		}

		coverVars[file] = &load.CoverVar{
			File: longFile,
			Var:  fmt.Sprintf("GoCover_%d_%x", coverIndex, h),
		}

		log.Printf("declareCoverVars, file: %s, var: %s \n", longFile, fmt.Sprintf("GoCover_%d_%x", coverIndex, h))
		coverIndex++
	}
	return coverVars
}
