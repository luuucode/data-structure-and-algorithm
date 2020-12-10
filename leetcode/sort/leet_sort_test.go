package sort

import "testing"

func TestMergeIntervals(t *testing.T) {

	//[[1,3],[2,6],[8,10],[15,18]]
	arr := [][]int{{5,6},{4,4},{3,3},{2,2},{5,5}}
	intervals := MergeIntervals(arr)
	t.Log(intervals)
}

func TestValidAnagram(t *testing.T) {
	s := "anagram"
	y := "nagaram"
	b := ValidAnagram(s, y)
	t.Log("ValidAnagram result: ",b)
}

func TestSortColors(t *testing.T) {
	arr := []int{2,2,0,0,2,1,1,1,0}
	SortColors(arr)
	t.Log(arr)

}

func TestLargeNumber(t *testing.T) {
	arr := []int{999}
	num := LargeNumber(arr)
	t.Log(num)
}

var nums1 = []int{1,2,2,1}
var nums2 = []int{2}

func TestIntersection(t *testing.T) {
	arr := Intersection(nums1, nums2)
	t.Log(arr)
}

func TestIntersection2(t *testing.T) {
	arr := Intersection2(nums1, nums2)
	t.Log(arr)
}