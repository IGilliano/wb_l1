package main

import "fmt"

/* Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human
(аналог наследования). */

type Human struct {
	Name  string
	Age   int
	Skill string
}

type Action struct {
	Human
}

func (h Human) Introduce() {
	fmt.Printf("Hello, world! My name is %s. Im %d years old! I can %s\n", h.Name, h.Age, h.Skill)
}

func main() {
	artist := Action{Human{Name: "Vlad", Age: 42, Skill: "play Sting on guitar"}}
	dancer := Action{Human{Name: "Lena", Age: 25, Skill: "dance pretty good"}}
	fmt.Println(artist)

	artist.Introduce()
	dancer.Introduce()
}
