package main

import (
	"errors"
	"fmt"
)

func main() {
	var fn = func() error { return errors.New("fail") }
	b, _ := do(fn, "")
	fmt.Println(b.String())
}

func do(fn func() error, s string) (fmt.Stringer, error) {
	if err := fn(); err != nil {
		return nil, err
	}

	return stringer(s), nil
}

type stringer string

func (s stringer) String() string {
	return string(s)
}
