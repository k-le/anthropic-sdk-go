package anthropic_test

import (
	"testing"

	"github.com/3JoB/resty-ilo"
	fc "github.com/3JoB/fasthttp-client"
)

func BenchmarkResty(b *testing.B) {
	b.ResetTimer()
	client := resty.New()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req, err := client.R().Get("https://example.com/")
			if err !=nil {
				panic(err)
			}
			req.RawBody().Close()
		}
	  })
}

func BenchmarkFast(b *testing.B) {
	b.ResetTimer()
	pool := fc.NewClientPool()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			client := pool.Get().(*fc.Client)
			_, err := client.Get("https://example.com/")
			if err !=nil {
				panic(err)
			}
			pool.Put(client)
		}
	  })
}

/*1.1.3

BenchmarkResty-32    	       1	1360634700 ns/op	  355216 B/op	    2421 allocs/op
BenchmarkFast-32     	       1	1214312100 ns/op	  178616 B/op	    1740 allocs/op

*/

/*1.1.2

goos: windows
goarch: amd64
pkg: github.com/3JoB/anthropic-sdk-go
cpu: Intel(R) Xeon(R) CPU E5-2670 0 @ 2.60GHz
BenchmarkResty-32    	       1	1437605100 ns/op	  363360 B/op	    2426 allocs/op
PASS
ok  	github.com/3JoB/anthropic-sdk-go	3.313s
*/

/* 1.1.1

goos: windows
goarch: amd64
pkg: github.com/3JoB/anthropic-sdk-go
cpu: Intel(R) Xeon(R) CPU E5-2670 0 @ 2.60GHz
BenchmarkResty-32    	       1	1377795500 ns/op	  354680 B/op	    2423 allocs/op
PASS
ok  	github.com/3JoB/anthropic-sdk-go	3.117s
*/