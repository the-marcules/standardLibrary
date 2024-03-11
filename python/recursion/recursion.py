import copy


def rev(input, index=0):
    if input == []:
        return [[]]
    if len(input) == 1:
        return [[input]]

    result = []
    # base case
    if index == len(input) - 1:  # index = 2
        result = [input[index]]
        return result

    elif index == 0:
        # else:
        liste = rev(input, index+1)

        for i in range(len(input)):
            for item in liste:
                new_item = copy.deepcopy(item)
                new_item.insert(i, input[index])
                result.append(new_item)

        return result
    else:  # index = 1
        num_to_add = input[index]
        prev_run = rev(input, index+1)
        current_run = copy.deepcopy(prev_run)
        prev_run.append(num_to_add)
        current_run.insert(0, num_to_add)
        result = [prev_run] + [current_run]

        return result


print("final result: ", rev([0, 1, 2]))


def check_output(output, expected_output):
    """
    Return True if output and expected_output
    contains the same lists, False otherwise.

    Note that the ordering of the list is not important.

    Examples:
        check_output([ [0, 1], [1, 0] ] ], [ [1, 0], [0, 1] ]) returns True

    Args:
        output(list): list of list
        expected_output(list): list of list

    Returns:
        bool
    """
    o = copy.deepcopy(output)  # so that we don't mutate input
    e = copy.deepcopy(expected_output)  # so that we don't mutate input

    o.sort()
    e.sort()
    return o == e


print("Pass" if (check_output(rev([0, 1, 2]), [[0, 1, 2], [0, 2, 1], [
      1, 0, 2], [1, 2, 0], [2, 0, 1], [2, 1, 0]])) else "Fail")
print("Pass" if (check_output(rev([]), [[]])) else "Fail")
print("Pass" if (check_output(rev([0]), [[0]])) else "Fail")
# print("Pass" if (check_output(rev([0, 1]), [[0, 1], [1, 0]])) else "Fail")
