package main

import (
	. "ladderGame/pkg/prio-queue"
	. "ladderGame/pkg/types"
)

func aStarLadderGame(game GameField) (result []int) {
	prio := PrioQ{}
	prio.Insert(Field{
		FieldNum:          game.Start,
		WayToField:        []int{},
		RemainingDistance: game.End,
	})
	var visited []int

	for prio.HasItems() {
		current := prio.Pop()
		largest := Field{}

		for _, dice := range game.Dice {
			nextNum := current.FieldNum + dice
			way := append(current.WayToField, dice)

			if nextNum > game.End {
				continue
			}

			ladderDestination, inLadders := game.Ladders[nextNum]
			if nextNum == game.End || inLadders == true && ladderDestination == game.End {
				if len(way) < len(result) || len(result) == 0 {
					result = append([]int{}, way...)
				}
			} else {
				snakeDestination, inSnakes := game.Snakes[nextNum]
				if inSnakes == false && inLadders == false && nextNum > largest.FieldNum {
					largest.FieldNum = nextNum
					largest.WayToField = append([]int{}, way...)
					largest.RemainingDistance = game.End - nextNum
				}

				if inLadders == true || inSnakes == true {
					var destinationField int
					if inLadders == true {
						destinationField = ladderDestination
					} else if inSnakes == true {
						destinationField = snakeDestination
					}
					if notVisited(destinationField, visited) {

						newField := Field{
							FieldNum:          destinationField,
							WayToField:        append([]int{}, way...),
							RemainingDistance: game.End - ladderDestination,
						}

						prio.Insert(newField)
						visited = append(visited, destinationField)

					}
				}

			}

		}

		if largest.FieldNum != 0 && notVisited(largest.FieldNum, visited) {
			prio.Insert(largest)
			visited = append(visited, largest.FieldNum)
		}
	}

	return
}

func notVisited(field int, visited []int) bool {
	for value := range visited {
		if field == value {
			return false
		}
	}
	return true
}
