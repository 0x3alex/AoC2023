import math
import numpy as np

with open('input.txt') as doc:
    lines = [line.rstrip('\n') for line in doc]

def count_matches(index):
    match_count = 0
    split_numbers = lines[index].split("|")
    winning = split_numbers[0][split_numbers[0].find(":") + 1:].split()
    on_card = split_numbers[1].split()
    for y in range(0, len(on_card)):
        if on_card[y] in winning:
            match_count = match_count + 1
    return match_count

def day4_part1(input_arr):
    point_sum = 0
    for x in range(0, len(input_arr)):
        if count_matches(x) == 1:
            point_sum = point_sum + 1
        if count_matches(x) > 1:
            point_sum = point_sum + int(math.pow(2, (count_matches(x)-1)))
    return point_sum


def day4_part2(input_arr):
    card_copies = []
    for y in range(0, len(input_arr)):
        card_copies.append(1)
    for x in range(0, len(input_arr)):
        for copies in range(1, card_copies[x]+1):
            if count_matches(x) >= 1:
                for z in range(1, count_matches(x)+1):
                    card_copies[x+z] = card_copies[x+z] + 1
    sum = 0
    print(card_copies)
    for i in range(0, len(card_copies)):
        sum = sum + card_copies[i]
    return sum

print(day4_part1(lines))
print(day4_part2(lines))