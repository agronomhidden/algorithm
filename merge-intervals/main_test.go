package merge_intervals

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	in := []struct {
		arr1   Intervals
		arr2   Intervals
		result Intervals
	}{{
		Intervals{v{2, 4}, v{6, 8}},
		Intervals{v{3, 7}, v{9, 10}},
		Intervals{v{3, 4}, v{6, 7}},
	}, {
		Intervals{v{2, 4}},
		Intervals{v{5, 7}, v{9, 10}},
		nil,
	}, {
		Intervals{v{2, 4}, v{10, 20}, v{22, 25}},
		Intervals{v{1, 30}},
		Intervals{v{2, 4}, v{10, 20}, v{22, 25}},
	}, {
		Intervals{v{2, 6}, v{8, 12}, v{14, 18}, v{20, 24}},
		Intervals{v{5, 9}, v{11, 15}, v{17, 21}, v{23, 25}},
		Intervals{v{5, 6}, v{8, 9}, v{11, 12}, v{14, 15}, v{17, 18}, v{20, 21}, v{23, 24}},
	}}

	for _, test := range in {
		require.Equal(t, test.result, Merge(test.arr1, test.arr2))
	}
}
