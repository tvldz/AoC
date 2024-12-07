import os, sys


total = 0
max_RGB = [12,13,14]
with open('input.txt') as file:
    for game in file:
        max = [0,0,0]
        game_number = game.split(":")[0].split(" ")[1]
        print(game_number)
        pulls = game.split(":")[1].split(";")
        for pull in pulls:
            counts = pull.split(",")
            for count in counts:
                count = count.strip()
                #print(count)
                if count.strip().split(" ")[1] == 'red' and int(count.strip().split(" ")[0]) > max[0]:
                    max[0] = int(count.strip().split(" ")[0])
                if count.strip().split(" ")[1] == 'green' and int(count.strip().split(" ")[0]) > max[1]:
                    max[1] = int(count.strip().split(" ")[0])
                if count.strip().split(" ")[1] == 'blue' and int(count.strip().split(" ")[0]) > max[2]:
                    max[2] = int(count.strip().split(" ")[0])
        print(max)
        print("---")
        if max[0] <= max_RGB[0] and max[1] <= max_RGB[1] and max[2] <= max_RGB[2]:
            total = total + int(game_number)
    print(total)