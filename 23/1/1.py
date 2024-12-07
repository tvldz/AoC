total = 0
with open('input.txt') as file:
    for line in file:
        digits = []
        line = line.strip()
        for index, char in enumerate(line):
            if char.isdigit():
                digits.append(index)
        total = total+ int(line[min(digits)]+line[max(digits)])
print(total)