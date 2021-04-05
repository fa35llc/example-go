package calc

import (
	"regexp"
	"strconv"
	"testing"
)

var (
	zero int // The initial value of int type is 0
	one  int = 1
	two  int = 2
)

// go test -v -cover
// go test -v ./calc/...
func TestSumSuccess(t *testing.T) {
	r, err := Sum(one, two)

	if err != nil {
		t.Fatalf("Failed func Sum test: %#v", err)
	}

	rs := strconv.Itoa(r)
	pattern := regexp.MustCompile(`^\d$`)
	if !pattern.MatchString(rs) {
		t.Fatalf("Failed func Sum test: Not int")
	}

}

func TestSumFailed(t *testing.T) {

	_, err := Sum(zero, two)

	if err == nil {
		t.Fatalf("Failed func sum test: %#v", err)
	}

}

// go test -bench . -benchmem
// go test -bench . -cpuprofile=cpu.prof
// go test -bench . -memprofile=mem.prof
// go tool pprof cpu.prof go-mod-test
func BenchmarkSum(*testing.B) {
	for i := 0; i < 100000000; i++ {
		Sum(one, two)
	}
}
