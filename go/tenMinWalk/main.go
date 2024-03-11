package main

import "fmt"

type Directions struct {
	Cn     int
	Cs     int
	Ce     int
	Cw     int
	Dir    []rune
	IsEven bool
}

func NewDirection(walk []rune) Directions {
	return Directions{
		Dir:    walk,
		Cn:     0,
		Cs:     0,
		Ce:     0,
		Cw:     0,
		IsEven: false,
	}
}

func (d *Directions) countDirs() {
	for _, direction := range d.Dir {
		switch direction {
		case 'n':
			d.Cn++
			break
		case 's':
			d.Cs++
			break
		case 'e':
			d.Ce++
			break
		case 'w':
			d.Cw++
			break
		}
	}
	if d.Cw == d.Ce && d.Cs == d.Cn {
		d.IsEven = true
	}
}

func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		// too short or too long
		return false
	}

	evalWalk := NewDirection(walk)
	evalWalk.countDirs()

	return evalWalk.IsEven
}

func main() {
	walk := []rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}
	fmt.Printf("Walk is 10 min and returns to start: %v\n", IsValidWalk(walk))
}
