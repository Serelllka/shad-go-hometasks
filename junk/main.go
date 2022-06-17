package main

import (
	"fmt"
)

const (
	Zero = `........
.111111.
.111111.
.11..11.
.11..11.
.11..11.
.11..11.
.11..11.
.11..11.
.111111.
.111111.
........`
)

type Point struct {
	x, y int
}

type Circle struct {
	Point
	radius int
}

func main() {
	str := []rune("jopa")
	fmt.Printf("%c", str[0])
}
