def nth_row_pascal(n):
    prevRow = []

    for rowCount in range(0, n+1):
        fieldsInRow = rowCount + 1
        currentRow = [None] * fieldsInRow
        for currentField in range(0, fieldsInRow):
            lastFieldsIndex = fieldsInRow - 1
            if currentField == lastFieldsIndex or currentField == 0:
                # if first or last item in row it has always the value of 1
                currentRow[currentField] = 1
            elif currentField > 0:
                currentRow[currentField] = prevRow[currentField - 1] + prevRow[currentField]
        prevRow = currentRow

    return currentRow


if __name__ == '__main__':
    print(nth_row_pascal(6))
