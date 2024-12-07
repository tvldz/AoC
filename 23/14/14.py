map = []
with open('input.txt') as file:
    for line in file:
        line.strip()
        map.append([x for x in line.strip()])

changes = 1
while changes != 0:
    changes = 0
    for row_index, row in enumerate(map):
        for col_index, col in enumerate(row):
            if row_index-1 > -1 and col == 'O' and map[row_index-1][col_index] == '.':
                map[row_index-1][col_index] = 'O'
                map[row_index][col_index] = '.'
                changes += 1
    print('--')
    for row in map:
        print(''.join(row))
    print(f'changes: {changes}')

weight = 0
for row_index,row in enumerate(map):
    for col_index, col in enumerate(row):
        if col == 'O':
            weight += (len(map))-row_index
print(weight)
map = list(zip(*map[::-1]))
for row in map:
    print(''.join(row))
