package main

import (
	"fmt"
	"time"
)

type person struct {
	name      string
	birthDate int
}

func (p person) getKoreanAge() int {
	currentYear := time.Now().Year()
	koreanAge := currentYear - p.birthDate + 1

	return koreanAge
}
func (p person) sayHello() {
	kAge := p.getKoreanAge()
	fmt.Printf("Hello! My name is %s and I'm %d", p.name, kAge)
}

func main() {
	milky := person{"milky", 1997}
	milky.sayHello()
}
