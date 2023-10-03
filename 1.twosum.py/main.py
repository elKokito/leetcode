def main():
    nums = [5, 10, 15, 4]

    print(twoSum(nums, 15))

def twoSum(nums: list, total: int):
    for idx, value in enumerate(nums):
        prevIndex, prevValue = idx, value
        for idxx, valuee in enumerate(nums[idx:]):
            if prevValue + valuee == total:
                return [prevIndex, idxx+prevIndex]
    return []

main()
