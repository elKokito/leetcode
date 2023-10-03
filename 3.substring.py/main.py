def substring(text):
    subMax = 1
    currentMax = 0
    currentSub = ""
    for c in text:
        if c not in currentSub:
            currentSub += c
            currentMax += 1
            if currentMax > subMax:
                subMax = currentMax
    return subMax

def main():
    print(substring('abcabcefg'))

if __name__ == '__main__':
    main()
