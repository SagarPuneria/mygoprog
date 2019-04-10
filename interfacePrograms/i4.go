package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	Animal
	Llama
	Cat
	JavaProgrammer
}

func (d *Dog) Speak() string {
	fmt.Println(d.Animal.Speak())
	fmt.Println(d.Llama.Speak())
	fmt.Println(d.JavaProgrammer.Speak())
	fmt.Println(d.Cat.Speak())
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l *Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	/*//animals := []Animal{&Dog{Animal: &Cat{}}, &Cat{}, &Llama{}, &JavaProgrammer{}}
	var animals = []Animal{&Dog{Animal: &Cat{}}, &Cat{}, &Llama{}, &JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}*/
	var ani Animal = &Dog{Cat: Cat{}, Llama: Llama{}, JavaProgrammer: JavaProgrammer{}, Animal: &Cat{}}
	fmt.Println(ani.Speak())
}
