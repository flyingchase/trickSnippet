if __name__ == "__main__":
    # function role means first class citizens in python
    one = ["one", "two", "three", "four", "five"]
    print(list(sorted(one, key=len)))

    # lambda function and map
    def two(x): return 1 if x == 0 else x * two(x - 1)
    print(list(map(two, range(6))))
