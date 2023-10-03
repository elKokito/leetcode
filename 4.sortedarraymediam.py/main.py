def main():
    arr1 = [1,2,3,4]
    arr2 = [4,5,6,6,7,8]
    print(mediam_sorted_arrays(arr1, arr2))

def mediam_sorted_arrays(arr1, arr2) -> int:
    idx1 = 0
    idx2 = 0
    merged = []
    while idx1 < len(arr1) and idx2 < len(arr2):
        if arr1[idx1] < arr2[idx2]:
            merged.append(arr1[idx1])
            idx1 += 1
        else:
            merged.append(arr2[idx2])
            idx2 += 1
    if idx1 == len(arr1):
        for i in range(idx2, len(arr2)):
            merged.append(arr2[i])
    else:
        for i in range(idx1, len(arr1)):
            merged.append(arr1[i])

    print(merged)
    parity = len(merged) % 2 == 0
    middleIndex = int(len(merged) / 2)
    median = None
    if parity:
        median = (merged[middleIndex - 1] + merged[middleIndex]) /2.0
    else:
        median = merged[middleIndex]
    return median

if __name__ == '__main__':
    main()
