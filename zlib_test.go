package compressbench

import (
	"compress/zlib"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkZlib(b *testing.B) {
	nw, err := zlib.NewWriterLevel(ioutil.Discard, zlib.NoCompression)
	require.NoError(b, err)

	sw, err := zlib.NewWriterLevel(ioutil.Discard, zlib.BestSpeed)
	require.NoError(b, err)

	cw, err := zlib.NewWriterLevel(ioutil.Discard, zlib.BestCompression)
	require.NoError(b, err)

	dw, err := zlib.NewWriterLevel(ioutil.Discard, zlib.DefaultCompression)
	require.NoError(b, err)

	hw, err := zlib.NewWriterLevel(ioutil.Discard, zlib.HuffmanOnly)
	require.NoError(b, err)

	tests := []struct {
		name string
		w    *zlib.Writer
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
		runBenchmark(b, "zlib_"+test.name, f)
	}
}
