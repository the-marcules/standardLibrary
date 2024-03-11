package main

func main() {
	println(BouncingBall(3, 0.66, 1.5))
}

func BouncingBall(h, bounce, window float64) (count int) {
	if h <= window || bounce >= 1 || bounce < 0 {
		return -1
	}

	for h > window {
		//starts falling
		count++

		// bounces to new height
		h = h * bounce
		if h > window {
			count++
		} else {
			break
		}
	}
	return
}
