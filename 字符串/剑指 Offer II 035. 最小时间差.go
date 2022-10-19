package 字符串

import (
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	if len(timePoints) > 1440 {
		return 0
	}
	time := []int{}
	for _, timePoint := range timePoints {
		hour, _ := strconv.Atoi(timePoint[:2])
		minutes, _ := strconv.Atoi(timePoint[3:])
		time = append(time, hour*60+minutes)
	}

	sort.Ints(time)
	ret := 24 * 60
	for i := 0; i < len(time)-1; i++ {
		if time[i+1]-time[i] < ret {
			ret = time[i+1] - time[i]
		}
	}

	if 24*60-time[len(time)-1]+time[0] < ret {
		ret = 24*60 - time[len(time)-1] + time[0]
	}

	return ret
}
