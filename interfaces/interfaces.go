package interfaces

import (
	"fmt"
	"math"
)

// Здесь я просто загуглил как расчитывать площадь и радиус фигур (потому что гуманитарий)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func Describe(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())

	switch s.(type) {
	case Rectangle:
		fmt.Println("Rectangle")
	case Circle:
		fmt.Println("Circle")
	}

}

type Bank interface {
	Deposit()
	Withdraw()
	Balance()
}

type DebitAcc struct {
	CurrentBalance uint
}
type CreditAcc struct {
	CurrentBalance int
}

func (d *DebitAcc) Deposit(sum uint) uint {
	d.CurrentBalance += sum
	return d.CurrentBalance
}

func (d *DebitAcc) Withdraw(sum uint) uint {
	d.CurrentBalance -= sum
	return d.CurrentBalance
}

func (d *DebitAcc) Balance() uint {
	return d.CurrentBalance
}

func (c *CreditAcc) Deposit(sum int) int {
	c.CurrentBalance += sum
	return c.CurrentBalance
}

func (c *CreditAcc) Withdraw(sum int) int {
	c.CurrentBalance -= sum
	return c.CurrentBalance
}

func (c *CreditAcc) Balance(sum int) int {
	return c.CurrentBalance
}
