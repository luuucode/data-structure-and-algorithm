package sort

func ValidAnagram(s string,t string) bool{
	//判断边界
	ls := len(s)
	lt := len(t)
	if  ls <= 0 || lt <= 0 || ls != lt {
		return false
	}



	return true
}