mappings = {}
with open('input.txt') as file:
    for line in file:
        if "seeds:" in line:
            seeds = line.split(":")[1].strip().split(" ")
            print(seeds)
        if "seed-to-soil" in line:
            loading = "seed-to-soil"
            mappings.update({loading: []})
        if "soil-to-fertilizer" in line:
            loading = "soil-to-fertilizer"
            mappings.update({loading: []})
        if "fertilizer-to-water" in line:
            loading = "fertilizer-to-water"
            mappings.update({loading: []})
        if "water-to-light" in line:
            loading = "water-to-light"
            mappings.update({loading: []})
        if "light-to-temperature" in line:
            loading = "light-to-temperature"
            mappings.update({loading: []})
        if "temperature-to-humidity" in line:
            loading = "temperature-to-humidity"
            mappings.update({loading: []})
        if "humidity-to-location" in line:
            loading = "humidity-to-location"
            mappings.update({loading: []})
        if line[0].isdigit():
            mappings[loading].append(line.strip().split(' '))
    for seed in seeds:
        for key in mappings:
            for map in mappings[key]:
                if int(seed) in range(int(map[1]),int(map[1])+int(map[2])):
                    seed = int(map[0])+(int(seed)-int(map[1]))
                    if key == 'humidity-to-location':
                        print(seed)
                    break
