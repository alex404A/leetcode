numLen = 7
pals = []
for num in range(1000000, 10000000):
    strNum = str(num)
    for i in range(numLen):
        j = numLen - i - 1
        if j < i:
            pals.append(num)
            break
        if strNum[i] != strNum[j]:
            break

intervalSet = set()
for i in range(1, len(pals)):
    intervalSet.add(pals[i] - pals[i - 1])
print(intervalSet)
