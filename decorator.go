package main

type Product interface {
	price() int
}

type Coffe struct{}

type addCereal struct {
	myProduct Product
}

type addChocolate struct {
	myProduct Product
}

func (c *Coffe) price() int {
	return 500
}

func (c *addChocolate) price() int {
	return c.myProduct.price() + 100
}

func (c *addCereal) price() int {
	return c.myProduct.price() + 50
}

/*
func main() {
	coffe := &Coffe{}
	coffeWithChocolate := &addChocolate{myProduct: coffe}
	coffeWithChocolateAndCereal := &addCereal{myProduct: coffeWithChocolate}
	fmt.Printf("The price of the coffe is %d tg", coffeWithChocolateAndCereal.price())
}*/
