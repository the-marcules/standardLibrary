def permutations(string, index=0):
    output = list()

    # early exit conditions
    if len(string) == 0:
        return output

    if len(string) == 1:
        output.append(string)
        return output

    tmp = permutations(string, index + 1)

    output.append()


def run_tests():

    test_cenarios = [
        {
            "test_str": "a",
            "test_result": ["a"],
        },
        {
            "test_str": "ab",
            "test_result": ["ab", "ba"],
        },
        {
            "test_str": "",
            "test_result": [],
        },
        {
            "test_str": "abc",
            "test_result": ['abc', 'bac', 'bca', 'acb', 'cab', 'cba'],
        },
    ]

    for test_cenario in test_cenarios:
        # assert (permutations(
        #     test_cenario["test_str"]).sort() == test_cenario["test_result"].sort())
        print("testing input of '", test_cenario["test_str"], "':", "Pass" if permutations(
            test_cenario["test_str"]) == test_cenario["test_result"] else "Fail")


run_tests()
