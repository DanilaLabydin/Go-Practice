package main

import (
	"fmt"
)

func main() {
	var nilMap map[string]int
	totalWins := map[string]int{}
	fmt.Println(nilMap)
	fmt.Println(totalWins)

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarach", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)

	ages := make(map[int][]string, 10)
	fmt.Println(ages)

	totalWins = map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"] += 1
	fmt.Println(totalWins)
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])

	// comma - ok
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["world"]
	fmt.Println(v, ok)
	v, ok = m["goodbuy"]
	fmt.Println(v, ok)
	fmt.Println()
	fmt.Println(m)

}
