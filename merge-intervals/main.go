package merge_intervals

import "math"

type v [2]int
type Intervals []v

func Merge(interval1, interval2 Intervals) (out Intervals) {
	for i, j := 0, 0; i < len(interval1) && j < len(interval2); {
		if isIntersect(interval1[i], interval2[j]) {
			out = append(out, v{
				int(math.Max(float64(interval1[i][0]), float64(interval2[j][0]))),
				int(math.Min(float64(interval1[i][1]), float64(interval2[j][1]))),
			})
		}
		if interval1[i][1] > interval2[j][1] {
			j++
		} else {
			i++
		}
	}
	return
}

func isIntersect(time1, time2 v) bool {
	return !(time1[1] < time2[0] || time2[1] < time1[0])
}
