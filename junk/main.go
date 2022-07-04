package main

import (
	"bytes"
	"fmt"
	"io"
	"testing/iotest"
)

type myInt int64

func (m *myInt) Jopa() {
	*m = 21
}

var bufferSize = 1000

func main() {
	text := make([]byte, bufferSize)
	read := make([]byte, bufferSize)

	buf := bytes.NewBuffer(text)
	errBuf := iotest.DataErrReader(buf)

	n, err := io.ReadAll(errBuf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(n), " ", len(read))
}
