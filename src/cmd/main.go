package main

type Human interface {
	Eat(food string)
	Walk(step string)
	Sleep(hour int)
}

type Dog struct {
	Name string
}
type Monkey struct {
	Color string
}

type Man struct {
	Hangout bool
}

func main() {
	// var x Dog
	// var y Monkey
	// var z Man
	// Evolutions(x)
	// Evolutions(y)
	// Evolutions(z)

}

func Evolutions(life Human) {

}

func (m Monkey) Eat(fruit string)
func (e Man) Eat(meet string)
func (d Dog) Eat(vet string)
func (m Monkey) Walk(s string)
func (e Man) Walk(s string)
func (j Monkey) Sleep(m int)
func (e Man) Sleep(h int)
func (e Man) Play(s string)
