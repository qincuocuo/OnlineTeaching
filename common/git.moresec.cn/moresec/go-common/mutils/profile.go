package mutils

import (
	"flag"
	"os"
	"runtime/pprof"
)

var (
	isDebug                                                                                                   bool
	cpuProfilingFile, memProfilingFile, blockProfilingFile, goroutineProfilingFile, threadcreateProfilingFile *os.File
)

func InitProfile() {

	flag.Parse()
	flag.BoolVar(&isDebug, "debug", false, "debug process")
	//fmt.Printf("%v", isDebug)
	if isDebug {
		cpuProfilingFile, _ = os.Create("cpu.prof")
		memProfilingFile, _ = os.Create("mem.prof")
		blockProfilingFile, _ = os.Create("block.prof")
		goroutineProfilingFile, _ = os.Create("goroutine.prof")
		threadcreateProfilingFile, _ = os.Create("threadcreat.prof")
		pprof.StartCPUProfile(cpuProfilingFile)
	}
}

func SaveProfile() {
	if isDebug {
		goroutine := pprof.Lookup("goroutine")
		goroutine.WriteTo(goroutineProfilingFile, 1)
		heap := pprof.Lookup("heap")
		heap.WriteTo(memProfilingFile, 1)
		block := pprof.Lookup("block")
		block.WriteTo(blockProfilingFile, 1)
		threadcreate := pprof.Lookup("threadcreate")
		threadcreate.WriteTo(threadcreateProfilingFile, 1)
		pprof.StopCPUProfile()
	}

}
