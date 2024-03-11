package types

type Field struct {
	FieldNum          int
	WayToField        []int
	RemainingDistance int
}

type GameField struct {
	Ladders map[int]int
	Snakes  map[int]int
	Dice    []int
	Start   int
	End     int
}
