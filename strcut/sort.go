package main

import (
	"sort"
	"fmt"
)

type byLength []string

func (s byLength) Len() int {
	return -len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	str := []string{"c123", "a12331", "b11"}
	sort.Strings(str)
	fmt.Print(str)

	sort.Sort(byLength(str))

	fmt.Println(str)

}
