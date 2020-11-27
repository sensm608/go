package main

import "fmt"

type subscriber struct {
	name   string
	rate   int
	active bool
}

func printInfo(s subscriber) {
	fmt.Println("Name: ", s.name)
	fmt.Println("MOnthly rate: ", s.rate)
	fmt.Println("Active? ", s.active)
}

func defaultSub(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5000
	s.active = true
	return s
}

func applyDiscount(s *subscriber){
	s.rate = 4000
}

func main() {
	s1 := defaultSub("Kim")
	s1.rate = 4500
	printInfo(s1)

	s2 := defaultSub("Kim")
	printInfo(s2)

	var s3 subscriber
	applyDiscount(&s3)
	fmt.Println(s3.rate)
}
