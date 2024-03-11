# def can_construct(word, arr) -> bool:
#     table = [False for i in range(0, len(word) + 1)]
#     table[0] = True
#     for pos in range(0, len(word) + 1):
#         if table[pos] is True:
#             for sub in arr:
#                 l_sub = len(sub)
#                 if word[pos:pos + l_sub] == sub:
#                     table[pos + l_sub] = True
#
#     return table[len(word)]
#
#
# print(f'{can_construct("test", ["t", "es"])}')
# print(f'{can_construct("skateboard", ["bo", "rd", "ate", "t", "ska", "sk", "boar"])}')
# print(f'{can_construct("enterapotentpot", ["t", "a", "p", "ent", "enter", "ot", "o"])}')
# print(f'{can_construct("eeeeeeeeeeeeeeeeeeeeeeeef", ["ee", "e", "eeeee", "eee", "eeeeee", "e", "e", "e"])}')


# def count_contruct(word, arr) -> int:
#     table = [0 for i in range(0, len(word) + 1)]
#     table[0] = 1
#     for pos in range(0, len(word) + 1):
#         if table[pos] > 0:
#             for sub in arr:
#                 l_sub = len(sub)
#                 if word[pos:pos + l_sub] == sub:
#                     table[pos + l_sub] += 1
#
#     return table[len(word)]
#
#
# print(f'{count_contruct("test", ["t", "es"])}')
# print(f'{count_contruct("skateboard", ["bo", "rd", "ate", "t", "ska", "sk", "boar"])}')
# print(f'{count_contruct("enterapotentpot", ["t", "a", "p", "ent", "enter", "ot", "o"])}')
# print(f'{count_contruct("eeeeeeeeeeeeeeeeeeeeeeeef", ["ee", "e", "eeeee", "eee", "eeeeee", "e", "e", "e"])}')


def all_construct(word, arr) -> int:
    table = [[] for i in range(0, len(word) + 1)]
    table[0].append([])
    for pos in range(0, len(word) + 1):
        if len(table[pos]) > 0:
            for sub in arr:
                l_sub = len(sub)
                if word[pos:pos + l_sub] == sub:
                    for element in table[pos]:
                        newElement = element.copy()
                        newElement.append(sub)
                        table[pos + l_sub].append(newElement)

    return table[-1]


print(f'{all_construct("test", ["t", "es"])}')
print(f'{all_construct("abcdef", ["ab", "cd", "def", "abcd", "abc", "ef"])}')
print(f'{all_construct("skateboard", ["bo", "rd", "ate", "t", "ska", "sk", "boar"])}')
print(f'{all_construct("enterapotentpot", ["t", "a", "p", "ent", "enter", "ot", "o"])}')
# print(f'{all_construct("eeeeeeeeeeeeeeeeeeeeeeeef", ["ee", "e", "eeeee", "eee", "eeeeee"])}')
