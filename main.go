package main

import "fmt"
import "io"

type thingy struct {
	data1 int
	data2 string
}

func (t *thingy) modify(in io.Reader) error {

	var err error
	t.data1 = int(in)
	return err

}

func main() {
	fmt.Println("vim-go")
}
