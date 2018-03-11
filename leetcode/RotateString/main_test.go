package rotatestring

import "testing"

// Data is abstraction for test input
type Data struct {
	A      string
	B      string
	Result bool
}

func TestRotateString(t *testing.T) {
	var result bool
	data := make([]Data, 0, 3)
	data = append(data, Data{"abcde", "cdeab", true})
	data = append(data, Data{"abcde", "abced", false})

	// i + 1 == case number
	for i, v := range data {
		result = rotateString(v.A, v.B)
		if result != v.Result {
			t.Errorf("Wrong result in case %d. A: %s, B: %s; got: %v, want: %v\n",
				i+1,
				v.A,
				v.B,
				result,
				v.Result)
		}
	}
}
