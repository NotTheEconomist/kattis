package main

import "fmt"

type Die interface {
	Average() float64
	Max() int
	Min() int
	NumSides() int
	String() string
}

type die struct {
	min, max int
}

func (d die) Average() float64 {
	sum := 0
	for n := d.min; n < d.max+1; n++ {
		sum += n
	}
	return float64(sum) / float64(d.NumSides())
}
func (d die) Max() int       { return d.max }
func (d die) Min() int       { return d.min }
func (d die) NumSides() int  { return d.max - d.min + 1 }
func (d die) String() string { return fmt.Sprintf("d%d:%d-%d", d.NumSides(), d.min, d.max) }

func main() {
	gunnar := make([]Die, 2)
	emma := make([]Die, 2)
	for i := 0; i < 2; i++ {
		var dMin, dMax int
		fmt.Scan(&dMin, &dMax)
		curDie := die{min: dMin, max: dMax}
		gunnar[i] = curDie
	}
	for i := 0; i < 2; i++ {
		var dMin, dMax int
		fmt.Scan(&dMin, &dMax)
		curDie := die{min: dMin, max: dMax}
		emma[i] = curDie
	}
	gAvg := gunnar[0].Average() + gunnar[1].Average()
	eAvg := emma[0].Average() + emma[1].Average()

	if eAvg > gAvg {
		fmt.Println("Emma")
	} else if gAvg > eAvg {
		fmt.Println("Gunnar")
	} else {
		fmt.Println("Tie")
	}
}
