package main

import (
	"fmt"
	"sort"
)

func main() {
	var a []string
	a = []string{"999","2","3"}
	sort.Sort(a)
	fmt.Println(a)

}