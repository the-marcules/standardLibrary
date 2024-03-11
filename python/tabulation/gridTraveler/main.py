def gridTraveler(m, n):
    table = [[0 for x in range(0, n + 1)] for i in range(m + 1)]
    table[1][1] = 1

    for i in range(0, m + 1):
        for j in range(0, n + 1):
            adder = table[i][j]
            if i + 1 <= m:
                table[i + 1][j] += adder
            if j + 1 <= n:
                table[i][j + 1] += adder

    return table[m][n]


print(f'Possible ways to walk through a 3x3 grid: {gridTraveler(3, 3)}')
print(f'Possible ways to walk through a 4x4 grid: {gridTraveler(4, 4)}')
print(f'Possible ways to walk through a 3x2 grid: {gridTraveler(3,2)}')
print(f'Possible ways to walk through a 18x18 grid: {gridTraveler(18,18)}')
