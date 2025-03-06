package main

import "fmt"

func main() {
	map2 := make(map[string]int)
	map2["ram"] = 100
	map2["shyam"] = 0
	map2["rohan"] = 80
	fmt.Println("ram", map2["ram"])
	fmt.Println("shyam", map2["shyam"])
	fmt.Println("Anurag", map2["Anurag"])

	value, exist := map2["ram"]

	fmt.Println("value:", value, "exist:", exist)
}
