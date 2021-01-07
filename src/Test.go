package main

import "fmt"

func main() {
	var Map map[int]bool
	Map = make(map[int]bool)
	Map[2] = true
	ok, _ := Map[1]
	fmt.Println(ok)
}
