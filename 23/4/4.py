import os, sys

total = 0
with open('input.txt') as file:
    for line in file:
        winning_numbers = line.split("|")[0].split(":")[1].strip().split()
        my_numbers = line.split("|")[1].strip().split()
        c = sum(el in my_numbers for el in winning_numbers)
        print(winning_numbers)
        print(my_numbers)
        if c == 0:
            print(f'{c},{0}')
            continue
        elif c == 1:
            total = total+1
            print(f'{c},{1}')
        elif c == 2:
            total = total+2
            print(f'{c},{2}')
        else:
            total = total + 2**(c-1)
            print(f'{c},{2**(c-1)}')
print(total)