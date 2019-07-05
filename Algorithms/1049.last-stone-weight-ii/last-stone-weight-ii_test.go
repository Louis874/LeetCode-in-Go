package problem1049

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	stones []int
	ans    int
}{

	{
		[]int{2, 7, 4, 1, 8, 1},
		1,
	},

	// 可以有多个 testcase
}

func Test_lastStoneWeightII(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, lastStoneWeightII(tc.stones), "输入:%v", tc)
	}
}

func Benchmark_lastStoneWeightII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lastStoneWeightII(tc.stones)
		}
	}
}
