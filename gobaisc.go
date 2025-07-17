package main

import "fmt"

type Product struct {
	name  string
	price string
	color string
}

func new_Product(name string, price string, color string) *Product {
	p := Product{
		name:  name,
		color: color,
		price: price,
	}

	return &p
}

// method in go -- similar to method in oop language like java

func (p *Product) display() {
	fmt.Println("product details : ")

	fmt.Println("product name : ", p.name)
	fmt.Println("product price : ", p.price)
	fmt.Println("product price : ", p.color)
}

func main() {
	new_P := new_Product("oppo", "5k", "blue")
	fmt.Println(new_P) //

	// calling method here.
	// We cannot call a method on its own — we must call it on an instance. new_P is the instance here
	// new_Product acts like a constructor function (Go doesn't have constructors like Java, but this simulates one).
	// display() is a method that can have any logic, just like a method in a Java class.
	// So, new_Product + display together simulate what you'd have in a Java class with constructor and method.
	new_P.display()
}
