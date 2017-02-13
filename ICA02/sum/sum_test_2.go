package sum

import "testing"
var sum_tests_uint32 = []struct {
	n132       uint32
	n232       uint32
	expected32 uint32
}{
	{1, 2, 3},
	{12000, 10000, 22000},
	{4294967294, 1, 4294967295},
}
func TestSumUint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := SumUint32(v.n132, v.n232); val != v.expected32 {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n132, v.n232, val, v.expected32)
		}
	}
}
