package compressbench

import (
	"io/ioutil"
	"testing"

	"github.com/golang/snappy"
)

func BenchmarkSnappy(b *testing.B) {
	w := snappy.NewWriter(ioutil.Discard)
	f := func(data []byte) error {
		_, err := w.Write(data)
		if err != nil {
			return err
		}
		return w.Flush()
	}
	runBenchmark(b, "snappy", f)
}
