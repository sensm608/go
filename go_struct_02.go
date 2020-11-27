package main

import "fmt"

type part struct {
	desc  string
	count int
}

func showInfo(p part) {
	fmt.Println("desc: ", p.desc)
	fmt.Println("count: ", p.count)
}

func minimumOrder(desc string) part {
	var p part
	p.desc = desc
	p.count = 1000
	return p
}

func main() {
	p := minimumOrder("Hex bolts")
	fmt.Println(p.desc,p.count)
	//var bolts part
	//bolts.desc = "Hex bolts"
	//bolts.count = 100
	//showInfo(bolts)
}
