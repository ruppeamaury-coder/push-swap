main go
package main

import "fmt"

type Piles struct {
	a          []int
	b          []int
	operations []string
}

func (p *Piles) ra() {
	if len(p.a) < 2 {
		return
	}
	p.a = append(p.a[1:], p.a[0])
	p.operations = append(p.operations, "ra")
}

func (p *Piles) sa() {
	if len(p.a) < 2 {
		return
	}
	p.a[0], p.a[1] = p.a[1], p.a[0]
	p.operations = append(p.operations, "sa")
}

func (p *Piles) pb() {
	if len(p.a) == 0 {
		return
	}
	p.b = append([]int{p.a[0]}, p.b...)
	p.a = p.a[1:]
	p.operations = append(p.operations, "pb")
}

func (p *Piles) pa() {
	if len(p.b) == 0 {
		return
	}
	p.a = append([]int{p.b[0]}, p.a...)
	p.b = p.b[1:]
	p.operations = append(p.operations, "pa")
}

func (p *Piles) afficher() {
	fmt.Println("a :", p.a, " b :", p.b)
}

func main() {
	p := Piles{a: []int{5, 2, 9, 1, 7}}
	p.afficher()
	p.pb()
	p.afficher()
	p.ra()
	p.afficher()
	p.pa()
	p.afficher()
	fmt.Println("opérations :", p.operations)
}
