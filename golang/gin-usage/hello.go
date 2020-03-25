package main

import "fmt"

type usage interface {
	hello()
}
type S struct {
	name string
	age  int
}

func (s *S) Read(p []byte) (n int, err error) {
	panic("not implemented") // TODO: Implement
}

func (s S) hello() {
	s.Read([]byte{'a', 'b'})
}

func use() {
	s := &S{
		age: 0,
	}
	fmt.Println(s.name)
	s.Read([]byte{'a', 'b'})
}
