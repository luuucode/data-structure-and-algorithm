package sort


//TODO 继续优化
func ValidAnagram(s string,t string) bool{
	//判断边界
	ls := len(s)
	lt := len(t)
	if ls == 0 || lt == 0{
		return true
	}
	if s == t{
		return true
	}
	if ls != lt {
		return false
	}
	lsCountMap := WordCount(s)
	ltCountMap := WordCount(t)
	for key,value := range lsCountMap {
		v,ok := ltCountMap[key]
		if v != value || ok == false{
			return false
		}
	}
	return true
}
func WordCount(s string) map[uint8]int {
	if len(s) <= 0{
		return nil
	}
	m := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		key := s[i]
		v,ok := m[key]
		if v == 0 || ok == false{
			m[key] = 1
		} else {
			m[key] += 1
		}
	}
	return m
}