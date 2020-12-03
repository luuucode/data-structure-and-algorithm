package sort

import "testing"

func TestMergeIntervals(t *testing.T) {

	//[[1,3],[2,6],[8,10],[15,18]]
	//arr := [][]int{ {9,13}, {2,6},{8,10}, {15,18}}
	arr := [][]int{ {1,4}, {4,5}}
	intervals := MergeIntervals(arr)
	t.Log(intervals)
}