def fibo_tablulation(n):
    table = [0 for i in range(0, n + 1)]
    table[1] = 1
    for i in range(1, n):
        table[i + 1] += table[i]
        if i + 2 <= n:
            table[i + 2] += table[i]
    # print(table)

    return table[n]


print(f'result: {fibo_tablulation(5)}')
print(f'result: {fibo_tablulation(6)}')
print(f'result: {fibo_tablulation(7)}')
