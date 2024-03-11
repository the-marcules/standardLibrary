def can_sum(target, arr) -> bool:
    table = [False for i in range(target + 1)]

    table[0] = True

    for pos in range(0, target):
        if table[pos] == True:
            for i in arr:
                if pos + i <= target:
                    table[pos + i] = True

    return table[target]


#
# print(f'{can_sum(7, [5, 3, 4])}')
# print(f'{can_sum(5, [5, 3, 4])}')
# print(f'{can_sum(1, [5, 3, 4])}')
# print(f'{can_sum(6, [5, 3, 4])}')
# print(f'{can_sum(8, [5, 3, 4])}')
# print(f'{can_sum(9, [5, 3, 4])}')
# print(f'{can_sum(10, [5, 3, 4])}')
# print('big one:')
# print(f'{can_sum(300, [7, 14])}')


def how_sum(target, arr) -> list:
    table = [None for i in range(target + 1)]

    table[0] = []

    for pos in range(0, target):
        if table[pos] is not None:
            for i in arr:
                if pos + i <= target:
                    table[pos + i] = table[pos] + [i]
    return table[target]


# print(f'{how_sum(7, [5, 3, 4])}')


def best_sum(target, arr) -> list:
    table = [None for i in range(target + 1)]

    table[0] = []

    for pos in range(0, target):
        if table[pos] is not None:
            for i in arr:
                if pos + i <= target:
                    if table[pos + i] is None or len(table[pos + i]) > len(table[pos] + [i]):
                        table[pos + i] = table[pos] + [i]
    return table[target]


print(f'{best_sum(8, [5, 3, 4])}')
