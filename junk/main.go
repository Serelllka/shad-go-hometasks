package main

import "fmt"

type node struct {
	name  string
	inner *node
}

func main() {
	var item = &node{}
	item.name = "lol"
	fmt.Println(item.name)
}
