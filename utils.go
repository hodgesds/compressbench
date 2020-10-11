package compressbench

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/hodgesds/perf-utils"
	"github.com/stretchr/testify/require"
	"golang.org/x/sys/unix"
)

const (
	kallsyms = "/proc/kallsyms"
)

func runIPCBench(b *testing.B, f func(b *testing.B)) {
	eventAttrs := []unix.PerfEventAttr{
		perf.CPUInstructionsEventAttr(),
		perf.CPUCyclesEventAttr(),
		perf.CPURefCyclesEventAttr(),
	}
	perf.RunBenchmarks(
		b,
		f,
		perf.BenchStrict,
		eventAttrs...,
	)
}

func runL1Bench(b *testing.B, f func(b *testing.B)) {
	eventAttrs := []unix.PerfEventAttr{
		perf.L1DataEventAttr(unix.PERF_COUNT_HW_CACHE_OP_READ, unix.PERF_COUNT_HW_CACHE_RESULT_ACCESS),
		perf.L1DataEventAttr(unix.PERF_COUNT_HW_CACHE_OP_READ, unix.PERF_COUNT_HW_CACHE_RESULT_MISS),
		perf.L1InstructionsEventAttr(unix.PERF_COUNT_HW_CACHE_OP_READ, unix.PERF_COUNT_HW_CACHE_RESULT_MISS),
	}
	perf.RunBenchmarks(
		b,
		f,
		perf.BenchStrict,
		eventAttrs...,
	)
}

func testData(b *testing.B) []byte {
	f, err := os.Open(kallsyms)
	require.NoError(b, err)
	data, err := ioutil.ReadAll(f)
	require.NoError(b, err)
	return data
}

func runBenchmark(b *testing.B, name string, f func([]byte) error) {
	data := testData(b)
	bench := func(b *testing.B) {
		b.ReportMetric(float64(len(data)), "bytes")
		b.SetBytes(int64(len(data)))
		for i := 0; i < b.N; i++ {
			f(data)
		}
	}
	b.Run(name+"_ipc", func(b *testing.B) {
		runIPCBench(b, bench)
	})
	b.Run(name+"_L1", func(b *testing.B) {
		runL1Bench(b, bench)
	})
}
