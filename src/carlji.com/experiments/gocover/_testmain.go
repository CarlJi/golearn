
package main

import (

	"testing"
	"testing/internal/testdeps"


	_test "carlji.com/experiments/flag"



	_cover0 "carlji.com/experiments/flag"



)

var tests = []testing.InternalTest{

	{"TestC", _test.TestC},

}

var benchmarks = []testing.InternalBenchmark{

}

var examples = []testing.InternalExample{

}



// Only updated by init functions, so no need for atomicity.
var (
	coverCounters = make(map[string][]uint32)
	coverBlocks = make(map[string][]testing.CoverBlock)
)

func init() {
	
	
	coverRegisterFile("carlji.com/experiments/flag/funcA.go", _cover0.GoCover_0.Count[:], _cover0.GoCover_0.Pos[:], _cover0.GoCover_0.NumStmt[:])
	
	coverRegisterFile("carlji.com/experiments/flag/main.go", _cover0.GoCover_1.Count[:], _cover0.GoCover_1.Pos[:], _cover0.GoCover_1.NumStmt[:])
	
	
}

func coverRegisterFile(fileName string, counter []uint32, pos []uint32, numStmts []uint16) {
	if 3*len(counter) != len(pos) || len(counter) != len(numStmts) {
		panic("coverage: mismatched sizes")
	}
	if coverCounters[fileName] != nil {
		// Already registered.
		return
	}
	coverCounters[fileName] = counter
	block := make([]testing.CoverBlock, len(counter))
	for i := range counter {
		block[i] = testing.CoverBlock{
			Line0: pos[3*i+0],
			Col0: uint16(pos[3*i+2]),
			Line1: pos[3*i+1],
			Col1: uint16(pos[3*i+2]>>16),
			Stmts: numStmts[i],
		}
	}
	coverBlocks[fileName] = block
}


func main() {



	testing.RegisterCover(testing.Cover{
		Mode: "set",
		Counters: coverCounters,
		Blocks: coverBlocks,
		CoveredPackages: "",
	})

	m := testing.MainStart(testdeps.TestDeps{}, tests, benchmarks, examples)

	_test.TestMain(m)

}

