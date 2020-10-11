package compressbench

import (
	"compress/lzw"
	"io/ioutil"
	"testing"
)

func BenchmarkLzw(b *testing.B) {
	w := lzw.NewWriter(ioutil.Discard, lzw.LSB, 8)
	f := func(data []byte) error {
		_, err := w.Write(data)
		return err
	}
	runBenchmark(b, "lzw", f)
}
