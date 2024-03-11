package main

func main() {

}

func TowerBuilder(nFloors int) (tower []string) {
	/*	if nFloors == 0 {
		return []string{}
	}*/
	starsInBasement := (nFloors * 2) - 1
	stars := 1
	for level := 1; level <= nFloors; level++ {

		if level != 1 {
			stars += 2
		}
		spaces := (starsInBasement - stars) / 2
		tower = append(tower, stringBuilder(stars, spaces))
	}

	return tower
}

func stringBuilder(stars, spaces int) (str string) {
	for i := 0; i < spaces; i++ {
		str += " "
	}
	for i := 0; i < stars; i++ {
		str += "*"
	}
	for i := 0; i < spaces; i++ {
		str += " "
	}
	return
}
