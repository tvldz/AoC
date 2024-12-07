total=0
grand_total = 0
with open('input.txt') as file:
    for line in file:
        print('---')
        all_readings = []
        readings = [int(i) for i in line.strip().split()]
        all_readings.append(readings)
        print(f'readings: {readings}')
        print(f'intitial list: {all_readings}')
        print(f'len: {len(set(readings))}')
        while len(set(readings)) > 1:
            new_reading = []
            for index in range(len(readings)-1):
                new_reading.append(int(readings[index+1])-int(readings[index]))
            all_readings.append(new_reading)
            readings = new_reading
            print(f'reading :{readings} {len(set(readings))}')
        all_readings.append([0 for x in range(len(all_readings[-1])-1)])
        all_readings.reverse()        
        print(f'all readings: {all_readings}')
        for reading in all_readings:
            reading.reverse()
        print(f'all readings reversed: {all_readings}')
        total = 0
        ends = []
        prev = all_readings[0][-1]
        for index, reading in enumerate(all_readings):
            print(f'{prev}')
            total = reading[-1]-prev
            prev = total
            print(total)
        grand_total += total
    print(grand_total)
