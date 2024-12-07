def get_type(hands_bids):
    hand = hands_bids[0]
    sorted_hand=sorted(hand)
    if all(char == hand[0] for char in hand):
        print(f'five of a kind:{hand}')
        return(19000000)
    #four of a kind
    elif all(char == sorted_hand[0] for char in sorted_hand[0:4]) \
            or all(char == sorted_hand[1] for char in sorted_hand[1:5]):
        print(f'four of a kind:{hand}')
        return(180000)
    #full house
    elif (all(char == sorted_hand[0] for char in sorted_hand[0:3]) \
            or all(char == sorted_hand[1] for char in sorted_hand[1:4]) \
            or all(char == sorted_hand[2] for char in sorted_hand[2:5])) \
            and len(set(hand)) == 2:
        print(f'full house:{hand}')
        return(17000)
    #three of a kind
    elif all(char == sorted_hand[0] for char in sorted_hand[0:3]) \
            or all(char == sorted_hand[1] for char in sorted_hand[1:4]) \
            or all(char == sorted_hand[2] for char in sorted_hand[2:5]):
        print(f'three of a kind:{hand}')
        return(1600)
    #two pair
    elif len(set(hand)) == 3:
        print(f'two pair:{hand}')
        return(150)
    #one pair
    elif len(set(hand)) == 4:
        print(f'one pair:{hand}')
        return(14)
    #high card
    else:
       print(f'high card:{hand}')
       return(1)
    
def print_type(hand):
    sorted_hand=sorted(hand)
    if all(char == hand[0] for char in hand):
        print(f'five of a kind:{hand}')
    #four of a kind
    elif all(char == sorted_hand[0] for char in sorted_hand[0:4]) \
            or all(char == sorted_hand[1] for char in sorted_hand[1:5]):
        print(f'four of a kind:{hand}')
    #full house
    elif (all(char == sorted_hand[0] for char in sorted_hand[0:3]) \
            or all(char == sorted_hand[1] for char in sorted_hand[1:4]) \
            or all(char == sorted_hand[2] for char in sorted_hand[2:5])) \
            and len(set(hand)) == 2:
        print(f'full house:{hand}')
    #three of a kind
    elif all(char == sorted_hand[0] for char in sorted_hand[0:3]) \
            or all(char == sorted_hand[1] for char in sorted_hand[1:4]) \
            or all(char == sorted_hand[2] for char in sorted_hand[2:5]):
        print(f'three of a kind:{hand}')
    #two pair
    elif len(set(hand)) == 3:
        print(f'two pair:{hand}')
    #one pair
    elif len(set(hand)) == 4:
        print(f'one pair:{hand}')
    #high card
    else:
       print(f'high card:{hand}')

def custom_sort(hand):
    total = 0
    sort_req = ['2','3','4','5','6','7','8','9','T','J','Q','K','A']
    multipliers = [70000,4000,300,20,1]
    for index, char in enumerate(hand):
        total=total+(sort_req.index(char)+1)*multipliers[index]
    return(total)

hands_bids = []
out_list = []
with open('input.txt') as file:
    for line in file:
        hands_bids.append(line.strip().split(" "))
    for hand_bid in hands_bids:
        hand_bid.append((get_type(hand_bid)*130000)+(custom_sort(hand_bid[0])))
        out_list.append(hand_bid)
sorted_list = sorted(out_list, key=lambda x: x[2])
print(sorted_list)

total = 0
for index, hand in enumerate(sorted_list):
    total = total+(int(hand[1])*(index+1))
    print(hand)
    print(print_type(hand[0]))
    print(index+1)
    print(int(hand[1])*(index+1))
print(f'total:{total}')