// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fix implements the ``go fix'' command.
package fix

import (
	"carlji.com/experiments/gocover/go1.11.2/pkg/base"
	"carlji.com/experiments/gocover/go1.11.2/pkg/cfg"
	"carlji.com/experiments/gocover/go1.11.2/pkg/load"
	"carlji.com/experiments/gocover/go1.11.2/pkg/modload"
	"carlji.com/experiments/gocover/go1.11.2/pkg/str"
	"fmt"
	"os"
)

var CmdFix = &base.Command{
	Run:       runFix,
	UsageLine: "go fix [packages]",
	Short:     "update packages to use new APIs",
	Long: `
Fix runs the Go fix command on the packages named by the import paths.

For more about fix, see 'go doc cmd/fix'.
For more about specifying packages, see 'go help packages'.

To run fix with specific options, run 'go tool fix'.

See also: go fmt, go vet.
	`,
}

func runFix(cmd *base.Command, args []string) {
	printed := false
	for _, pkg := range load.Packages(args) {
		if modload.Enabled() && !pkg.Module.Main {
			if !printed {
				fmt.Fprintf(os.Stderr, "go: not fixing packages in dependency modules\n")
				printed = true
			}
			continue
		}
		// Use pkg.gofiles instead of pkg.Dir so that
		// the command only applies to this package,
		// not to packages in subdirectories.
		files := base.RelPaths(pkg.InternalAllGoFiles())
		base.Run(str.StringList(cfg.BuildToolexec, base.Tool("fix"), files))
	}
}
