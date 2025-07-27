package main

import "fmt"

type CN struct {
	cn int
}

type m interface {
	add()
	sub()
	mult()
}

func (p *CN) display() {
	fmt.Println("cn is ", p.cn)
}

func (t *CN) add() {
	fmt.Println("adding the number")
}

func (t *CN) sub() {
	fmt.Println("subtracting number")
}

func (t *CN) mult() {
	fmt.Println("multiply number")
}

func imp(p m) {
	p.add()
	p.sub()
	p.mult()
}

func newComplex(cn int) *CN {
	p := CN{
		cn: cn,
	}
	return &p
}

func test() {
	k := newComplex(10)

	k.display()

	imp(k)

}
