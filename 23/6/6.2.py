record = [41667266,244104712281040]
count=0
for test_time in range(record[0]):
    if test_time*(record[0]-test_time) > record[1]:
        print(test_time)
        count=count+1
print(f'count: {count}')