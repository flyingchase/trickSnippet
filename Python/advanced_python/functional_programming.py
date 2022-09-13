if __name__ == "__main__":
    nums = []
    for i in range(5):
        nums.append(i)
    print(nums)
    r = map(lambda x: x*x, nums)
    print(r)
