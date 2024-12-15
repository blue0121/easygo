package id

import "fmt"

type Id interface {
	Bytes() []byte
	fmt.Stringer
}
