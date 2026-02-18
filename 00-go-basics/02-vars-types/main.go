package main

import "fmt"

type Profile struct {
	Name string
	Age  int
}

func Sum(n []int) int {
	t := 0
	for _, v := range n {
		t += v
	}
	return t
}
func Lookup(m map[string]int, k string) (int, bool) { v, ok := m[k]; return v, ok }
func main()                                         { fmt.Println(Sum([]int{1, 2, 3}), Profile{"Ada", 30}) }
