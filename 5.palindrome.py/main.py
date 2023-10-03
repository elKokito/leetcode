def main():

    print(longuestPal("abada"))

def longuestPal(text: str) -> str:
    maxPal = ""
    for index in range(len(text)):
        end = len(text)
        while end > index:
            if isPal(text[index:end]) and (end - index) > len(maxPal):
                maxPal = text[index:end]
            end -= 1
    return maxPal

def isPal(text: str) -> bool:
    return text == text[::-1]


if __name__ == '__main__':
    main()
