package main

func main() {

}

func wave(words string) (wave []string) {
	wave = []string{} // leaving the array uninitialised would be better practice
	for i, c := range words {
		if c == ' ' {
			continue
		}
		upperC := string(c - 'a' + 'A')
		wave = append(wave, words[:i]+upperC+words[i+1:])
	}
	return
}

/*func wave(words string) []string {
	// preparations
	words = strings.ToLower(words)
	waveCount := len(strings.ReplaceAll(words, " ", ""))
	result := make([]string, waveCount)

	// build wave blocks
WAVEBLOCKS:
	for i := range result {
		letterCount := 0
		for n := 0; n < len(words); n++ {
			if words[n:n+1] == " " {
				continue
			}
			if letterCount == i {
				result[i] = integrateWaveToString(words, n)

				continue WAVEBLOCKS
			}
			letterCount++
		}

	}

	return result
}

func integrateWaveToString(str string, pos int) string {
	if pos == 0 {
		return strings.ToUpper(str[pos:pos+1]) + str[pos+1:]
	} else if pos == len(str)-1 {
		return str[:pos] + strings.ToUpper(str[pos:pos+1])
	}
	return str[:pos] + strings.ToUpper(str[pos:pos+1]) + str[pos+1:]

}
*/
