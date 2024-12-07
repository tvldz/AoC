import math
import itertools

map = []
def expand_universe(map):
    empty_rows = []
    for index, row in enumerate(map):
        if row[0] == '.' and len(set(row)) == 1:
            empty_rows.append(index)
    empty_columns = [x for x in range(len(map[0]))]
    for row in map:
        for index, column in enumerate(row):
            if column=='#':
                try:
                    empty_columns.remove(index)
                except:
                    continue
    print(f'empty_columns: {empty_columns}')
    print(f'empty_rows: {empty_rows}')
    row.insert(col_index,'.')
    
    row_offset=0
    for row_index,row in enumerate(map):
        if row_index in empty_rows:
            map.insert(row_index+row_offset,['.' for i in range(len(map[0]))])
            row_offset+=1
    for row_index,row in enumerate(map):
        col_offset=0
        for col_index, column in enumerate(row):
            if col_index in empty_columns:
                col_offset+=1
                print(f'col_index: {col_index}')
                print(f'col_offset: {col_offset}')
                row.insert(col_index+col_offset,'.')

with open('input.txt') as file:
    for line in file:
        line.strip()
        map.append([x for x in line.strip()])

for line in map:
    print(line)
print('--')
expand_universe(map)
for line in map:
    print(line)

all_galaxies = []
for row_index, row in enumerate(map):
    for col_index, column in enumerate(row):
        if column == '#':
            all_galaxies.append([row_index,col_index])
res = [(a, b) for idx, a in enumerate(all_galaxies) for b in all_galaxies[idx + 1:]]
print(res)
print(len(res))
total = 0
for pair in (res):
    print(f'{pair[0]},{pair[1]}')
    total += abs(pair[0][1] - pair[1][1])+abs(pair[0][0] - pair[1][0])

print(total)
