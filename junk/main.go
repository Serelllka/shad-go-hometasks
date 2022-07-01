package main

type myInt int64

func (m myInt) Jopa() {

}

func main() {
	var x myInt
	x = 7
	x.Jopa()
}
