digits = {'1':1,
          '2':2,
          '3':3,
          '4':4,
          '5':5,
          '6':6,
          '7':7,
          '8':8,
          '9':9,
          'one':1,
          'two':2,
          'three':3,
          'four':4,
          'five':5,
          'six':6,
          'seven':7,
          'eight':8,
          'nine':9}

total = 0
with open('input.txt') as file:
    for line in file:
        line = line.strip()
        print(line)
        lowest = len(line)
        highest = -1
        l_location = len(line)+1
        r_location = -1
        for key in digits:
            l_location = line.find(key)
            r_location = line.rfind(key)
            if (r_location > highest) and (not r_location < 0):
                highest = r_location
                second_digit = key
            if (l_location < lowest) and (not l_location < 0):
                lowest = l_location
                first_digit = key
        print(f'highest: {highest}')
        print(f'lowest: {lowest}')
        print(f'first_digit: {first_digit}')
        print(f'second_digit: {second_digit}')
        print(f'{digits[first_digit]}{digits[second_digit]}')
        total = total + int(f'{digits[first_digit]}{digits[second_digit]}')
        print(total)
print(total)