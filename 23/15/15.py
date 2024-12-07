with open('input.txt') as f:
    hashes = f.readline().strip()

total = 0

def hash_digest(input):
    chars = [*input]
    current_value = 0
    for char in chars:
        current_value = ((current_value + ord(char))*17) % 256
    return(current_value)


hashes_list = hashes.split(',')
print(hashes_list)
for hash in hashes_list:
    total += hash_digest(hash)
print(total)