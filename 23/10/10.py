map = []
steps = 0
def start(map):
    for row_index, row in enumerate(map):
        for col_index, col in enumerate(row):
            if col == 'S':
                if map[row_index][col_index-1] in ['-','F','L']:
                    return {'last':[row_index,col_index],'location':[row_index,col_index-1]}
                elif map[row_index][col_index+1] in ['-','J','7']:
                    return {'last':[row_index,col_index],'location':[row_index-1,col_index+1]}
                elif map[row_index+1][col_index] in ['|','L','J']:
                    return {'last':[row_index,col_index],'location':[row_index+1,col_index]}
                elif map[row_index-1][col_index] in ['|','7','F']:
                    return {'last':[row_index,col_index],'location':[row_index-1,col_index]}

with open('input.txt') as file:
    for line in file:
        line.strip()
        map.append([x for x in line])
    beginning = start(map)
    location = beginning['location']
    last = beginning['last']
    print(location)
    print(last)
    while map[location[0]][location[1]] != 'S':
        if map[location[0]][location[1]] == '|':
            if [location[0]+1,location[1]] != [last[0],last[1]]:
                last[0] = location[0]
                last[1] = location[1]
                location[0] = location[0]+1
                location[1] = location[1]
            elif [location[0]-1,location[1]] != [last[0],last[1]]:
                last[0] = location[0]
                last[1] = location[1]
                location[0] = location[0]-1
                location[1] = location[1]
        elif map[location[0]][location[1]] == '-':
            if [location[0],location[1]+1] != [last[0],last[1]]:
                last[0] = location[0]
                last[1] = location[1]
                location[0] = location[0]
                location[1] = location[1]+1
            elif [location[0],location[1]-1] != [last[0],last[1]]:
                last[0] = location[0]
                last[1] = location[1]
                location[0] = location[0]
                location[1] = location[1]-1
        elif map[location[0]][location[1]] == 'L':
            if [location[0]-1,location[1]] != [last[0],last[1]]:
                last[0] = location[0]
                last[1] = location[1]
                location[0] = location[0]-1
                location[1] = location[1]
            elif [location[0],location[1]+1] != [last[0],last[1]]:
                last[0] = location[0]
                last[1] = location[1]
                location[0] = location[0]
                location[1] = location[1]+1
        elif map[location[0]][location[1]] == 'J':
            if [location[0]-1,location[1]] != [last[0],last[1]]:
                last[0] = location[0]
                last[1] = location[1]
                location[0] = location[0]-1
                location[1] = location[1]
            elif [location[0],location[1]-1] != [last[0],last[1]]:
                last[0] = location[0]
                last[1] = location[1]
                location[0] = location[0]
                location[1] = location[1]-1
        elif map[location[0]][location[1]] == '7':
            if [location[0],location[1]-1] != [last[0],last[1]]:
                last[0] = location[0]
                last[1] = location[1]
                location[0] = location[0]
                location[1] = location[1]-1
            elif [location[0]+1,location[1]+1] != [last[0],last[1]]:
                last[0] = location[0]
                last[1] = location[1]
                location[0] = location[0]+1
                location[1] = location[1]
        elif map[location[0]][location[1]] == 'F':
            if [location[0],location[1]+1] != [last[0],last[1]]:
                last[0] = location[0]
                last[1] = location[1]
                location[0] = location[0]
                location[1] = location[1]+1
            elif [location[0]+1,location[1]] != [last[1],last[1]]:
                last[0] = location[0]
                last[1] = location[1]
                location[0] = location[0]+1
                location[1] = location[1]
        steps += 1
        print(map[location[0]][location[1]])
print(steps)