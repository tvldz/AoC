import os, sys

def check_validity(matrix,row,col):
    modifications = [0,-1,1]
    for row_mod in modifications:
        for col_mod in modifications:
            try:
                if matrix[row+row_mod][col+col_mod] not in ['0','1','2','3','4','5','6','7','8','9','.','\n']:
                    print(matrix[row+row_mod][col+col_mod])
                    return True
            except:
                continue
    return False

total = 0
matrix = []
with open('input.txt') as file:
    for line in file:
        matrix.append([x for x in line])
    for row_index, row in enumerate(matrix):
        digits = []
        valid = False
        for col_index, col in enumerate(row):
            if col.isdigit():
                digits.append(col)
                if valid is not True:
                    valid = check_validity(matrix,row_index,col_index)
            else:
                if len(digits) >= 1:
                    if valid == True:
                        total = total + int(''.join(digits))
                    digits = []
                    valid = False

print(total)
