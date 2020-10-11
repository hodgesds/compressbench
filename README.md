# Go Compression Benchmarks
This is an example repo that demonstrates the types of data that can be
generated for Go benchmarks. Specifically, with the use of the
[`perf`](https://github.com/hodgesds/perf-utils) library.

Almost any sort of system event (ie `perf list` or
`/sys/kernel/debug/tracing/available_events`) can be used to annotate Go
benchmarks. The difficulty is setting up the correct profilers using
[`perf_event_open`](https://www.man7.org/linux/man-pages/man2/perf_event_open.2.html).


# Example
One thing that this example does is limit the number of profiled hardware
events per benchmark run. This is due to the fact that CPUs have limited
profiling support and profilers are scheduled on hardware. For more info on
this check out this [blog
post](https://hadibrais.wordpress.com/2019/09/06/the-linux-perf-event-scheduling-algorithm/).
Here is an example of the output from a test run:

```
goos: linux
goarch: amd64
pkg: github.com/hodgesds/compressbench
BenchmarkGzip/gzip_no_compression_ipc-8         	    1140	   1020410 ns/op	9728.88 MB/s	   9927445 bytes	   3324072 hw_cycles/op	   5509615 hw_instr/op	   2834654 hw_ref_cycles/op	      39 B/op	       0 allocs/op
BenchmarkGzip/gzip_no_compression_L1-8          	    1171	   1016819 ns/op	9763.24 MB/s	   9927445 bytes	    469611 cache_l1d_miss/op	    961275 cache_l1d_read/op	        73.5 cache_l1i_miss/op	      38 B/op	       0 allocs/op
BenchmarkGzip/gzip_best_speed_ipc-8             	      14	  74959988 ns/op	 132.44 MB/s	   9927445 bytes	 243873970 hw_cycles/op	 491701177 hw_instr/op	 207813804 hw_ref_cycles/op	    3227 B/op	      15 allocs/op
BenchmarkGzip/gzip_best_speed_L1-8              	      15	  74811975 ns/op	 132.70 MB/s	   9927445 bytes	   7169538 cache_l1d_miss/op	 268460927 cache_l1d_read/op	     74628 cache_l1i_miss/op	    3013 B/op	      14 allocs/op
BenchmarkGzip/gzip_best_compression_ipc-8       	       1	1113706886 ns/op	   8.91 MB/s	   9927445 bytes	3616870757 hw_cycles/op	6309155481 hw_instr/op	3071880513 hw_ref_cycles/op	  859096 B/op	     231 allocs/op
BenchmarkGzip/gzip_best_compression_L1-8        	       1	1110565720 ns/op	   8.94 MB/s	   9927445 bytes	 247652098 cache_l1d_miss/op	1664696449 cache_l1d_read/op	     82087 cache_l1i_miss/op	   45760 B/op	     221 allocs/op
BenchmarkGzip/gzip_default_ipc-8                	       4	 292626046 ns/op	  33.93 MB/s	   9927445 bytes	 957650053 hw_cycles/op	1753393177 hw_instr/op	 817361269 hw_ref_cycles/op	   11350 B/op	      53 allocs/op
BenchmarkGzip/gzip_default_L1-8                 	       4	 292341482 ns/op	  33.96 MB/s	   9927445 bytes	  34350759 cache_l1d_miss/op	 474046297 cache_l1d_read/op	     42857 cache_l1i_miss/op	   11410 B/op	      54 allocs/op
BenchmarkGzip/gzip_huffman_ipc-8                	      30	  38023038 ns/op	 261.09 MB/s	   9927445 bytes	 123477217 hw_cycles/op	 331963520 hw_instr/op	 105231286 hw_ref_cycles/op	    1513 B/op	       7 allocs/op
BenchmarkGzip/gzip_huffman_L1-8                 	      30	  37688437 ns/op	 263.41 MB/s	   9927445 bytes	    796028 cache_l1d_miss/op	  85418306 cache_l1d_read/op	     11986 cache_l1i_miss/op	    1521 B/op	       7 allocs/op
BenchmarkLzw/lzw_ipc-8                          	      10	 110821433 ns/op	  89.58 MB/s	   9927445 bytes	 364244702 hw_cycles/op	 556310722 hw_instr/op	 307044492 hw_ref_cycles/op	    4540 B/op	      21 allocs/op
BenchmarkLzw/lzw_L1-8                           	       9	 111201452 ns/op	  89.27 MB/s	   9927445 bytes	   3535195 cache_l1d_miss/op	 167227827 cache_l1d_read/op	     13319 cache_l1i_miss/op	    5047 B/op	      23 allocs/op
BenchmarkSnappy/snappy_ipc-8                    	      69	  16919915 ns/op	 586.73 MB/s	   9927445 bytes	  55261654 hw_cycles/op	  75146302 hw_instr/op	  47183424 hw_ref_cycles/op	     658 B/op	       3 allocs/op
BenchmarkSnappy/snappy_L1-8                     	      67	  16878001 ns/op	 588.19 MB/s	   9927445 bytes	   1499924 cache_l1d_miss/op	  10139639 cache_l1d_read/op	      2688 cache_l1i_miss/op	     674 B/op	       3 allocs/op
BenchmarkZlib/zlib_no_compression_ipc-8         	     253	   4689095 ns/op	2117.13 MB/s	   9927445 bytes	  15086950 hw_cycles/op	  55673337 hw_instr/op	  12869976 hw_ref_cycles/op	     179 B/op	       0 allocs/op
BenchmarkZlib/zlib_no_compression_L1-8          	     259	   4617137 ns/op	2150.13 MB/s	   9927445 bytes	    470461 cache_l1d_miss/op	  10274816 cache_l1d_read/op	       187 cache_l1i_miss/op	     176 B/op	       0 allocs/op
BenchmarkZlib/zlib_best_speed_ipc-8             	      14	  78588003 ns/op	 126.32 MB/s	   9927445 bytes	 509301722 hw_cycles/op	1084197596 hw_instr/op	 217243293 hw_ref_cycles/op	    3272 B/op	      15 allocs/op
BenchmarkZlib/zlib_best_speed_L1-8              	      14	  77862437 ns/op	 127.50 MB/s	   9927445 bytes	   7141116 cache_l1d_miss/op	 287370530 cache_l1d_read/op	     66022 cache_l1i_miss/op	    3244 B/op	      15 allocs/op
BenchmarkZlib/zlib_best_compression_ipc-8       	       1	1123071860 ns/op	   8.84 MB/s	   9927445 bytes	3614172616 hw_cycles/op	6359507608 hw_instr/op	3080394603 hw_ref_cycles/op	  859296 B/op	     236 allocs/op
BenchmarkZlib/zlib_best_compression_L1-8        	       1	1106879684 ns/op	   8.97 MB/s	   9927445 bytes	 247319012 cache_l1d_miss/op	1674258126 cache_l1d_read/op	     77879 cache_l1i_miss/op	   45424 B/op	     215 allocs/op
BenchmarkZlib/zlib_default_ipc-8                	       4	 297186962 ns/op	  33.40 MB/s	   9927445 bytes	 972441977 hw_cycles/op	1803405965 hw_instr/op	 826169292 hw_ref_cycles/op	   11406 B/op	      54 allocs/op
BenchmarkZlib/zlib_default_L1-8                 	       4	 295998276 ns/op	  33.54 MB/s	   9927445 bytes	  34356358 cache_l1d_miss/op	 483427495 cache_l1d_read/op	     44238 cache_l1i_miss/op	   11456 B/op	      55 allocs/op
BenchmarkZlib/zlib_huffman_ipc-8                	      27	  41583820 ns/op	 238.73 MB/s	   9927445 bytes	 270824467 hw_cycles/op	 765490873 hw_instr/op	 115365865 hw_ref_cycles/op	    1673 B/op	       7 allocs/op
BenchmarkZlib/zlib_huffman_L1-8                 	      27	  41337869 ns/op	 240.15 MB/s	   9927445 bytes	   1593954 cache_l1d_miss/op	 190109153 cache_l1d_read/op	     23297 cache_l1i_miss/op	    1688 B/op	       8 allocs/op
PASS
ok  	github.com/hodgesds/compressbench	143.882s
```

# Future Direction
In theory it's possible to add eBPF programs to benchmarks as well, I haven't
taken the time to fully investigate that yet.
