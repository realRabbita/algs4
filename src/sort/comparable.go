package sort

type ComparableSlice interface {
	Swap(i, j int)
	Compare(i, j int) bool
	Length() int
	IsSorted() bool
}

type IntSlice []int

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntSlice) Compare(i, j int) bool {
	return s[i] <= s[j]
}
func (s IntSlice) Length() int {
	return len(s)
}
func (s IntSlice) IsSorted() bool {
	var n = s.Length()
	for i := 0; i < n-1; i++ {
		if !s.Compare(i, i+1) {
			return false
		}
	}
	return true
}

type StringSlice []string

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s StringSlice) Compare(i, j int) bool {
	return s[i] <= s[j]
}
func (s StringSlice) Length() int {
	return len(s)
}
func (s StringSlice) IsSorted() bool {
	var n = s.Length()
	for i := 0; i < n-1; i++ {
		if !s.Compare(i, i+1) {
			return false
		}
	}
	return true
}
