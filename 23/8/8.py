network = {}
with open('input.txt') as file:
    for index, line in enumerate(file):
        if index==0:
            instructions = [*line.strip()]
        if index>=2:
            key = line.split(" ")[0]
            left = line.split(" ")[2][1:4]
            right = line.split(" ")[3][0:3]
            network[key] = [left,right]

iteration = 0

node = 'AAA'
while(True):
    print(instructions[ iteration % len(instructions)])
    print(['L','R'].index(instructions[ iteration % len(instructions)]))
    node = network[node][['L','R'].index(instructions[ iteration % len(instructions)])]
    iteration = iteration+1
    if node == 'ZZZ':
        print(iteration)
        exit()