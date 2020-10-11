package compressbench

import (
	"compress/gzip"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkGzip(b *testing.B) {
	nw, err := gzip.NewWriterLevel(ioutil.Discard, gzip.NoCompression)
	require.NoError(b, err)

	sw, err := gzip.NewWriterLevel(ioutil.Discard, gzip.BestSpeed)
	require.NoError(b, err)

	cw, err := gzip.NewWriterLevel(ioutil.Discard, gzip.BestCompression)
	require.NoError(b, err)

	dw, err := gzip.NewWriterLevel(ioutil.Discard, gzip.DefaultCompression)
	require.NoError(b, err)

	hw, err := gzip.NewWriterLevel(ioutil.Discard, gzip.HuffmanOnly)
	require.NoError(b, err)

	tests := []struct {
		name string
		w    *gzip.Writer
	}{
		{
			name: "no_compression",
			w:    nw,
		},
		{
			name: "best_speed",
			w:    sw,
		},
		{
			name: "best_compression",
			w:    cw,
		},
		{
			name: "default",
			w:    dw,
		},
		{
			name: "huffman",
			w:    hw,
		},
	}

	for _, test := range tests {
		f := func(data []byte) error {
			_, err := test.w.Write(data)
			if err != nil {
				return err
			}
			return test.w.Flush()
		}
		runBenchmark(b, "gzip_"+test.name, f)
	}
}
