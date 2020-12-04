package sort

import "testing"

func TestMergeIntervals(t *testing.T) {

	//[[1,3],[2,6],[8,10],[15,18]]
	arr := [][]int{{5,6},{4,4},{3,3},{2,2},{5,5}}
	intervals := MergeIntervals(arr)
	t.Log(intervals)
}