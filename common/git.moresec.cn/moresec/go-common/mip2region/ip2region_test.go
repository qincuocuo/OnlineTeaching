package mip2region

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkMemorySearch(B *testing.B) {
	fmt.Println(os.Getwd())
	region, err := New("ip2region.db")
	if err != nil {
		B.Error(err)
	}
	for i := 0; i < B.N; i++ {
		region.MemorySearch("104.27.148.30")
	}

}
