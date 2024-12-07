total=1
records = [[41,244],[66,1047],[72,1228],[66,1040]]
for record in records:
    count=0
    for test_time in range(record[0]):
        if test_time*(record[0]-test_time) > record[1]:
            print(test_time)
            count=count+1
    print(f'count: {count}')
    total=total*count
print(f'total: {total}')